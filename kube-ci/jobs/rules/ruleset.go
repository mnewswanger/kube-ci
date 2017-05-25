package rules

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// Ruleset contains one or more rules or rulesets along with a mode
// Modes:
//   all - match all rules to pass
//   any - match any of the rules to pass
type Ruleset struct {
	Name             string          `json:"name"`
	Mode             string          `json:"mode"`
	Rules            []ruleInterface `json:"rules"`
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

func (rs *Ruleset) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	var err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	return rs.unmarshalRaw(raw)
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

func (rs *Ruleset) unmarshalRaw(raw map[string]interface{}) error {
	if _, e := raw["name"]; e {
		rs.Name = raw["name"].(string)
	}
	if _, e := raw["mode"]; e {
		rs.Mode = raw["mode"].(string)
	}
	if _, e := raw["rules"]; e {
		var rulesRaw = raw["rules"].([]interface{})
		for _, ruleRaw := range rulesRaw {
			// Here, the rule could refer to either a Ruleset or Rule since it's an interface
			var rule = ruleRaw.(map[string]interface{})
			// Rulesets have a rules property; Rules do not
			if _, e = rule["rules"]; e {
				// It's a ruleset
				var rulesetToAdd = Ruleset{}
				if err := rulesetToAdd.unmarshalRaw(rule); err != nil {
					return err
				}
				rs.Rules = append(rs.Rules, &rulesetToAdd)
			} else {
				// It's a rule
				var ruleToAdd = Rule{}
				if err := ruleToAdd.unmarshalRaw(rule); err != nil {
					return err
				}
				rs.Rules = append(rs.Rules, &ruleToAdd)
			}
		}
	}
	return nil
}
