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
	"gitlab_push_generic": {
		"git.repository": "repo",
		"git.branch":     "feature",
		"git.action":     "push",
	},
	"gitlab_push_foo_develop": {
		"git.repository": "foo",
		"git.branch":     "develop",
		"git.action":     "push",
	},
	"gitlab_push_foo_master": {
		"git.repository": "foo",
		"git.branch":     "master",
		"git.action":     "push",
	},
	"gitlab_merge_foo_master": {
		"git.repository": "foo",
		"git.branch":     "master",
		"git.action":     "merge",
	},
	"generic": {
		"my.label": "test",
		"my.name":  "value",
	},
}

var rulesToTest = map[string]ruleInterface{
	"gitlab_request": &Rule{
		Name:      "gitlab_request",
		LabelName: "git.repository",
		MatchMode: "exists",
	},
	"gitlab_project_foo": &Rule{
		Name:       "gitlab_project_foo",
		LabelName:  "git.repository",
		LabelValue: "foo",
		MatchMode:  "exact",
	},
	"gitlab_branch_develop": &Rule{
		Name:       "gitlab_branch_master",
		LabelName:  "git.branch",
		LabelValue: "develop",
		MatchMode:  "exact",
	},
	"gitlab_branch_master": &Rule{
		Name:       "gitlab_branch_master",
		LabelName:  "git.branch",
		LabelValue: "master",
		MatchMode:  "exact",
	},
	"not_gitlab_request": &Rule{
		Name:        "not_gitlab_request",
		InvertMatch: true,
		LabelName:   "git.repository",
		MatchMode:   "exists",
	},
	"gitlab_master_branch_regex": &Rule{
		Name:       "gitlab_branch_regex",
		LabelName:  "git.branch",
		LabelValue: "m.+r",
		MatchMode:  "regex",
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
	{
		ruleName:     "gitlab_request",
		labelSetName: "gitlab_push_generic",
	}: true,
	{
		ruleName:     "not_gitlab_request",
		labelSetName: "gitlab_push_generic",
	}: false,
	{
		ruleName:     "gitlab_project_foo",
		labelSetName: "gitlab_push_foo_develop",
	}: true,
	{
		ruleName:     "gitlab_branch_master",
		labelSetName: "gitlab_push_foo_develop",
	}: false,
	{
		ruleName:     "gitlab_branch_master",
		labelSetName: "gitlab_push_foo_master",
	}: true,
	{
		ruleName:     "gitlab_branch_master",
		labelSetName: "gitlab_merge_foo_master",
	}: true,
	{
		ruleName:     "gitlab_master_branch_regex",
		labelSetName: "gitlab_merge_foo_master",
	}: true,
}

func TestRuleMatches(t *testing.T) {
	for testData, expectedResult := range ruleRequestMap {
		var r = rulesToTest[testData.ruleName]
		var labels, exists = requestLabels[testData.labelSetName]
		if !exists {
			panic("Invalid label specified for test: " + testData.labelSetName)
		}
		if r.Matches(labels) != expectedResult {
			t.Error(
				"Rule match failed:", testData.labelSetName, "->", testData.ruleName, "Expected:", expectedResult,
			)
		}
	}
}

// TEST RULESETS
var rulesetsToTest = map[string]ruleInterface{
	"gitlab": &Ruleset{
		Name: "gitlab",
		Mode: "all",
		Rules: []ruleInterface{
			rulesToTest["gitlab_request"],
		},
	},
	"gitlab_with_ruleset_child": &Ruleset{
		Name: "gitlab_with_ruleset_child",
		Mode: "all",
		Rules: []ruleInterface{
			&Ruleset{
				Name: "gitlab",
				Mode: "all",
				Rules: []ruleInterface{
					rulesToTest["gitlab_request"],
				},
			},
		},
	},
	"always_fail": &Ruleset{
		Name: "always_fail",
		Mode: "all",
		Rules: []ruleInterface{
			rulesToTest["gitlab_request"],
			rulesToTest["not_gitlab_request"],
		},
	},
	"always_pass": &Ruleset{
		Name: "always_pass",
		Mode: "any",
		Rules: []ruleInterface{
			rulesToTest["gitlab_request"],
			rulesToTest["not_gitlab_request"],
		},
	},
	"gitlab_example": &Ruleset{
		Name: "gitlab_example",
		Mode: "all",
		Rules: []ruleInterface{
			rulesToTest["gitlab_request"],
			rulesToTest["gitlab_project_foo"], &Ruleset{
				Name: "Branch master or develop",
				Mode: "any",
				Rules: []ruleInterface{
					rulesToTest["gitlab_branch_master"],
					rulesToTest["gitlab_branch_develop"],
				},
			},
		},
	},
}

var rulesetRequestMap = map[*struct {
	rulesetName  string
	labelSetName string
}]bool{
	{
		rulesetName:  "gitlab",
		labelSetName: "empty",
	}: false,
	{
		rulesetName:  "gitlab",
		labelSetName: "gitlab_push_generic",
	}: true,
	{
		rulesetName:  "gitlab_with_ruleset_child",
		labelSetName: "gitlab_push_generic",
	}: true,
	{
		rulesetName:  "always_fail",
		labelSetName: "empty",
	}: false,
	{
		rulesetName:  "always_pass",
		labelSetName: "empty",
	}: true,
	{
		rulesetName:  "gitlab_example",
		labelSetName: "gitlab_push_foo_master",
	}: true,
}

func TestRuleSetMatches(t *testing.T) {
	for testData, expectedResult := range rulesetRequestMap {
		var r = rulesetsToTest[testData.rulesetName]
		var labels, exists = requestLabels[testData.labelSetName]
		if !exists {
			panic("Invalid label specified for test: " + testData.labelSetName)
		}
		if r.Matches(labels) != expectedResult {
			t.Error(
				"Ruleset match failed:", testData.labelSetName, "->", testData.rulesetName, "Expected:", expectedResult,
			)
		}
	}
}
