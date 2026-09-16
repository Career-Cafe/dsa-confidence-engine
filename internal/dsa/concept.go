package dsa

type ComplexitySignal struct {
	ExpectedTime  string `yaml:"expected_time" json:"expected_time"`
	ExpectedSpace string `yaml:"expected_space" json:"expected_space"`
}

type Concept struct {
	ID                  string            `yaml:"id" json:"id"`
	Name                string            `yaml:"name" json:"name"`
	Parent              string            `yaml:"parent" json:"parent"`
	Category            string            `yaml:"category" json:"category"`
	Description         string            `yaml:"description" json:"description"`
	Aliases             []string          `yaml:"aliases" json:"aliases"`
	Keywords            []string          `yaml:"keywords" json:"keywords"`
	RelatedConcepts     []string          `yaml:"related_concepts" json:"related_concepts"`
	ConfusableConcepts  []string          `yaml:"confusable_concepts" json:"confusable_concepts"`
	CodeSignals         []string          `yaml:"code_signals" json:"code_signals"`
	ComplexitySignals   *ComplexitySignal `yaml:"complexity_signals,omitempty" json:"complexity_signals,omitempty"`
	ExplanationExamples []string          `yaml:"explanation_examples" json:"explanation_examples"`
}
