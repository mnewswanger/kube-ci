package rules

import (
	"testing"
)

var validationTestData = map[*Rule]bool{
	&Rule{
		Name:       "",
		LabelName:  "",
		LabelValue: "",
		MatchMode:  "",
	}: false,
	&Rule{
		Name:       "",
		LabelName:  "label",
		LabelValue: "value",
		MatchMode:  "exact",
	}: false,
	&Rule{
		Name:       "Missing label name",
		LabelName:  "",
		LabelValue: "value",
		MatchMode:  "exact",
	}: false,
	&Rule{
		Name:       "Missing label value",
		LabelName:  "label",
		LabelValue: "",
		MatchMode:  "exact",
	}: true,
	&Rule{
		Name:       "Exact Match",
		LabelName:  "label",
		LabelValue: "value",
		MatchMode:  "exact",
	}: true,
	&Rule{
		Name:        "No Exact Match",
		LabelName:   "label",
		LabelValue:  "value",
		MatchMode:   "exact",
		InvertMatch: true,
	}: true,
	&Rule{
		Name:       "Regex Match",
		LabelName:  "label",
		LabelValue: "value",
		MatchMode:  "regex",
	}: true,
	&Rule{
		Name:        "No Regex Match",
		LabelName:   "label",
		LabelValue:  "value",
		MatchMode:   "regex",
		InvertMatch: true,
	}: true,
	&Rule{
		Name:       "Bad Regex",
		LabelName:  "label",
		LabelValue: "(value",
		MatchMode:  "regex",
	}: false,
	&Rule{
		Name:       "Exists",
		LabelName:  "label",
		LabelValue: "value",
		MatchMode:  "exists",
	}: true,
	&Rule{
		Name:       "MatchMode case senitivity",
		LabelName:  "label",
		LabelValue: "value",
		MatchMode:  "Exact",
	}: false,
}

func TestRuleValidationFailures(t *testing.T) {
	for rule, pass := range validationTestData {
		if rule.Validates() != pass {
			t.Error(
				"Rule did not match expected validation state: ", rule.Name,
			)
		}
	}
}

// VALIDATE RULESETS

var requestLabels = map[string]map[string]string{
	"empty": {},
	"gitlab": {
		"git.repository": "repo",
		"git.branch":     "master",
		"git.action":     "push",
	},
	"generic": {
		"my.label": "test",
		"my.name":  "value",
	},
}

var rulesToTest = map[string]Rule{
	"gitlab_request": {
		Name:        "gitlab_request",
		LabelName:   "git.repository",
		LabelValue:  "",
		MatchMode:   "exact",
		InvertMatch: true,
	},
	"not_gitlab_request": {
		Name:       "not_gitlab_request",
		LabelName:  "git.repository",
		LabelValue: "",
		MatchMode:  "exact",
	},
}

var ruleRequestMap = map[*struct {
	ruleName     string
	labelSetName string
}]bool{
	{
		ruleName:     "gitlab_request",
		labelSetName: "empty",
	}: false,
	{
		ruleName:     "not_gitlab_request",
		labelSetName: "empty",
	}: true,
}

func TestRulePasses(t *testing.T) {
	for testData, expectedResult := range ruleRequestMap {
		var r = rulesToTest[testData.ruleName]
		if r.Passes(requestLabels[testData.labelSetName]) != expectedResult {
			t.Error(
				"Rule pass state did not match expected: " + testData.labelSetName + "->" + testData.ruleName,
			)
		}
	}
}

// TEST RULESETS
