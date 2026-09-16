package config_test

import (
	"os"
	"testing"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/config"
)

func TestConfigDefaults(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("expected valid default config, got error: %v", err)
	}

	if cfg.AppAddr != ":8080" {
		t.Errorf("expected default AppAddr :8080, got %s", cfg.AppAddr)
	}
	if cfg.RunnerTimeoutMs != 3000 {
		t.Errorf("expected default RunnerTimeoutMs 3000, got %d", cfg.RunnerTimeoutMs)
	}
	if cfg.AcceptThreshold != 0.95 {
		t.Errorf("expected default AcceptThreshold 0.95, got %f", cfg.AcceptThreshold)
	}
	if cfg.RejustifyThreshold != 0.90 {
		t.Errorf("expected default RejustifyThreshold 0.90, got %f", cfg.RejustifyThreshold)
	}
}

func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		wantErr bool
	}{
		{
			name: "invalid timeout",
			env: map[string]string{
				"RUNNER_TIMEOUT_MS": "-1",
			},
			wantErr: true,
		},
		{
			name: "rejustify greater than accept",
			env: map[string]string{
				"REJUSTIFY_THRESHOLD": "0.96",
				"ACCEPT_THRESHOLD":    "0.95",
			},
			wantErr: true,
		},
		{
			name: "threshold out of range",
			env: map[string]string{
				"ACCEPT_THRESHOLD": "1.5",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				os.Setenv(k, v)
				defer os.Unsetenv(k)
			}
			_, err := config.Load()
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected wantErr=%v, got err=%v", tt.wantErr, err)
			}
		})
	}
}
