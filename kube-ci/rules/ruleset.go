package rules

// Ruleset contains one or more rules or rulesets along with a mode
// Modes:
//   all - match all rules to pass
//   any - match any of the rules to pass
type Ruleset struct {
	Name  string `json:"name"`
	Mode  string `json:"mode"`
	Rules []interface {
		Passes(map[string]string) bool
		Validates() bool
	}
	validationErrors []string
}

func (rs Ruleset) Passes(labels map[string]string) bool {
	if !rs.Validates() {
		return false
	}

	var defaultValue = false
	if rs.Mode == "all" {
		defaultValue = true
	}
	for _, r := range rs.Rules {
		if r.Passes(labels) {
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

func (rs Ruleset) Validates() bool {
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
