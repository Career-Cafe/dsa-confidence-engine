package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
)

type LocalRunner struct {
	defaultTimeout time.Duration
}

func NewLocalRunner(timeoutMs int) *LocalRunner {
	if timeoutMs <= 0 {
		timeoutMs = 3000
	}
	return &LocalRunner{
		defaultTimeout: time.Duration(timeoutMs) * time.Millisecond,
	}
}

type runnerPayload struct {
	Code       string           `json:"code"`
	Entrypoint string           `json:"entrypoint"`
	Tests      []model.TestCase `json:"tests"`
}

type runnerResponse struct {
	Status      string           `json:"status"`
	PassedCount int              `json:"passed_count"`
	FailedCount int              `json:"failed_count"`
	TotalCount  int              `json:"total_count"`
	Details     []TestCaseDetail `json:"details"`
	Error       string           `json:"error,omitempty"`
}

const pythonHarness = `
import sys
import json
import traceback

def run():
    try:
        raw_input = sys.stdin.read()
        payload = json.loads(raw_input)
    except Exception as e:
        print(json.dumps({
            "status": "FAIL",
            "passed_count": 0,
            "failed_count": 1,
            "total_count": 1,
            "error": f"Failed to parse test payload: {str(e)}",
            "details": []
        }))
        return

    code = payload.get("code", "")
    entrypoint_name = payload.get("entrypoint", "solve")
    tests = payload.get("tests", [])

    env = {}
    try:
        compiled = compile(code, "<submission>", "exec")
        exec(compiled, env)
    except Exception as e:
        tb = traceback.format_exc()
        print(json.dumps({
            "status": "FAIL",
            "passed_count": 0,
            "failed_count": len(tests) if tests else 1,
            "total_count": len(tests) if tests else 1,
            "error": f"Compilation/Execution error: {str(e)}\n{tb}",
            "details": []
        }))
        return

    if entrypoint_name not in env or not callable(env[entrypoint_name]):
        print(json.dumps({
            "status": "FAIL",
            "passed_count": 0,
            "failed_count": len(tests) if tests else 1,
            "total_count": len(tests) if tests else 1,
            "error": f"Entrypoint function '{entrypoint_name}' not found or not callable",
            "details": []
        }))
        return

    entrypoint = env[entrypoint_name]
    passed_count = 0
    failed_count = 0
    details = []

    def normalize(val):
        if val is None:
            return "None"
        if isinstance(val, bool):
            return str(val)
        if isinstance(val, (int, float, str)):
            return str(val)
        try:
            return json.dumps(val, sort_keys=True)
        except Exception:
            return str(val)

    def compare(actual, expected_str):
        # First try raw string match
        actual_str = normalize(actual)
        if actual_str == expected_str:
            return True
        # Try JSON comparison
        try:
            exp_json = json.loads(expected_str)
            act_json = json.loads(actual_str)
            if exp_json == act_json:
                return True
            # For list comparisons where order might not matter or string vs bool
            if isinstance(exp_json, list) and isinstance(act_json, list):
                if exp_json == act_json:
                    return True
        except Exception:
            pass
        # Try python bool representation
        if expected_str.lower() in ("true", "false"):
            if str(actual).lower() == expected_str.lower():
                return True
        return False

    for t in tests:
        test_id = t.get("id", "")
        test_input_str = t.get("input", "")
        expected_str = t.get("expected_output", "")

        try:
            parsed_input = json.loads(test_input_str)
            if isinstance(parsed_input, dict):
                result = entrypoint(**parsed_input)
            elif isinstance(parsed_input, list):
                result = entrypoint(*parsed_input)
            else:
                result = entrypoint(parsed_input)

            actual_str = normalize(result)
            is_pass = compare(result, expected_str)
            if is_pass:
                passed_count += 1
                details.append({
                    "id": test_id,
                    "passed": True,
                    "expected": expected_str,
                    "actual": actual_str
                })
            else:
                failed_count += 1
                details.append({
                    "id": test_id,
                    "passed": False,
                    "expected": expected_str,
                    "actual": actual_str,
                    "error": f"Expected {expected_str}, got {actual_str}"
                })
        except Exception as e:
            failed_count += 1
            details.append({
                "id": test_id,
                "passed": False,
                "expected": expected_str,
                "actual": "",
                "error": f"Runtime error: {str(e)}"
            })

    status = "PASS" if failed_count == 0 and passed_count > 0 else "FAIL"
    print(json.dumps({
        "status": status,
        "passed_count": passed_count,
        "failed_count": failed_count,
        "total_count": len(tests),
        "details": details
    }))

if __name__ == "__main__":
    run()
`

func (r *LocalRunner) Run(ctx context.Context, submission model.Submission, problem model.Problem) (TestResult, error) {
	start := time.Now()

	execCtx, cancel := context.WithTimeout(ctx, r.defaultTimeout)
	defer cancel()

	cmd := exec.CommandContext(execCtx, "python3", "-c", pythonHarness)

	payload := runnerPayload{
		Code:       submission.SourceCode,
		Entrypoint: problem.Entrypoint,
		Tests:      problem.Tests,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return TestResult{
			Status: model.TestStatusFail,
			Error:  fmt.Sprintf("failed to serialize test payload: %v", err),
		}, nil
	}

	cmd.Stdin = bytes.NewReader(payloadBytes)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	durationMs := time.Since(start).Milliseconds()

	if execCtx.Err() == context.DeadlineExceeded {
		return TestResult{
			Status:      model.TestStatusFail,
			PassedCount: 0,
			FailedCount: len(problem.Tests),
			TotalCount:  len(problem.Tests),
			Error:       "Execution timed out",
			DurationMs:  durationMs,
			Output:      stderr.String(),
		}, nil
	}

	stdoutStr := strings.TrimSpace(stdout.String())
	if stdoutStr == "" {
		errMsg := strings.TrimSpace(stderr.String())
		if errMsg == "" && err != nil {
			errMsg = err.Error()
		}
		return TestResult{
			Status:      model.TestStatusFail,
			PassedCount: 0,
			FailedCount: len(problem.Tests),
			TotalCount:  len(problem.Tests),
			Error:       errMsg,
			DurationMs:  durationMs,
		}, nil
	}

	var resp runnerResponse
	if parseErr := json.Unmarshal([]byte(stdoutStr), &resp); parseErr != nil {
		return TestResult{
			Status:      model.TestStatusFail,
			PassedCount: 0,
			FailedCount: len(problem.Tests),
			TotalCount:  len(problem.Tests),
			Error:       fmt.Sprintf("invalid test runner response: %s (stderr: %s)", stdoutStr, stderr.String()),
			DurationMs:  durationMs,
			Output:      stdoutStr,
		}, nil
	}

	return TestResult{
		Status:      model.TestStatus(resp.Status),
		PassedCount: resp.PassedCount,
		FailedCount: resp.FailedCount,
		TotalCount:  resp.TotalCount,
		Details:     resp.Details,
		DurationMs:  durationMs,
		Error:       resp.Error,
		Output:      stdoutStr,
	}, nil
}
