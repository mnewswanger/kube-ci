package rules

import "github.com/sirupsen/logrus"

// Ruleset contains one or more rules or rulesets along with a mode
// Modes:
//   all - match all rules to pass
//   any - match any of the rules to pass
type Ruleset struct {
	Name             string `json:"name"`
	Mode             string `json:"mode"`
	Rules            []ruleInterface
	validationErrors []string
}

// Matches returns true if the provided labels match the ruleset
// This applies recursively to provide options for complex matches
func (rs *Ruleset) Matches(labels map[string]string) bool {
	logger.WithFields(logrus.Fields{
		"rule_name": rs.Name,
	}).Debug("Comparing labels against rule")

	if !rs.Validates() {
		return false
	}

	var defaultValue = false
	if rs.Mode == "all" {
		defaultValue = true
	}
	for _, r := range rs.Rules {
		logger.Debug()
		if r.Matches(labels) {
			if rs.Mode == "any" {
				return true
			}
		} else {
			if rs.Mode == "all" {
				return false
			}
		}
	}
	return defaultValue
}

// Validates returns true when the ruleset and all of its children validate
func (rs *Ruleset) Validates() bool {
	var childrenValidated = true
	if rs.Name == "" {
		rs.addValidationError("Missing Name property")
	}
	if rs.Mode != "" {
		switch rs.Mode {
		case "all":
			break
		case "any":
			break
		default:
			rs.addValidationError("Unsupported Mode: " + rs.Mode)
		}
	} else {
		rs.addValidationError("Missing Mode property")
	}
	for _, r := range rs.Rules {
		if !r.Validates() {
			childrenValidated = false
		}
	}
	return childrenValidated
}

func (rs *Ruleset) addValidationError(e string) {
	rs.validationErrors = append(rs.validationErrors, e)
}
