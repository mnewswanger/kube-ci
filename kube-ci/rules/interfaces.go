package rules

type ruleInterface interface {
	Matches(map[string]string) bool
	Validates() bool
}
