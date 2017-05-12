package rules

import (
	"regexp"
)

// Rule contains match information to apply to the labels provided by the caller or produced by the system
// Each rule has a label to work against and its value.  The comparison is determined by match_mode:
// Match Modes:
//   exact - String comparison
//   exists - Label exists
//   regex - Regex comparison
type Rule struct {
	Name             string `json:"name"`
	LabelName        string `json:"label_name"`
	LabelValue       string `json:"label_value"`
	MatchMode        string `json:"match_mode"`
	InvertMatch      bool   `json:"invert_match"`
	validationErrors []string
	regex            *regexp.Regexp
}

func (r *Rule) Passes(labels map[string]string) bool {
	if !r.Validates() {
		return false
	}

	var labelValue, labelExists = labels[r.LabelName]
	switch r.MatchMode {
	case "exact":
		return (!r.InvertMatch && r.LabelValue == labelValue) || (r.InvertMatch && r.LabelValue != labelValue)
	case "exists":
		return (!r.InvertMatch && labelExists) || (r.InvertMatch && !labelExists)
	case "regex":
		return false
	}
	return false
}

func (r *Rule) Validates() bool {
	if r.Name == "" {
		r.addValidationError("Missing Name property")
	}
	if r.LabelName == "" {
		r.addValidationError("Missing Label Name property")
	}
	if r.MatchMode != "" {
		switch r.MatchMode {
		case "exact":
			break
		case "exists":
			break
		case "regex":
			var e error
			r.regex, e = regexp.Compile(r.LabelValue)
			if e != nil {
				r.addValidationError("Regex failed to compile")
			}
		default:
			r.addValidationError("Unknown Match Mode: " + r.MatchMode)
		}
	} else {
		r.addValidationError("Missing Match Mode property")
	}
	return len(r.validationErrors) == 0
}

func (r *Rule) addValidationError(e string) {
	r.validationErrors = append(r.validationErrors, e)
}
