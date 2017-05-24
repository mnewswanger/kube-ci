package rules

import (
	"testing"

	"github.com/fatih/color"
	"github.com/ghodss/yaml"
)

func TestUnmarshalYaml(t *testing.T) {
	var y, e = yaml.Marshal(rules)
	if e != nil {
		panic(e)
	}
	color.White(string(y))

	var rules Ruleset

	var err = yaml.Unmarshal(y, &rules)
	if err != nil {
		t.Error("Unmarshal failed")
	}

	var y2, _ = yaml.Marshal(rules)
	color.Yellow(string(y2))
}

var rules = Ruleset{
	Name: "gitlab_example",
	Mode: "all",
	Rules: []ruleInterface{
		&Rule{
			Name:      "gitlab_request",
			LabelName: "git.repository",
			MatchMode: "exists",
		},
		&Rule{
			Name:       "gitlab_project_foo",
			LabelName:  "git.repository",
			LabelValue: "foo",
			MatchMode:  "exact",
		}, &Ruleset{
			Name: "Branch master or develop",
			Mode: "any",
			Rules: []ruleInterface{
				&Rule{
					Name:       "gitlab_branch_develop",
					LabelName:  "git.branch",
					LabelValue: "develop",
					MatchMode:  "exact",
				},
				&Rule{
					Name:       "gitlab_branch_master",
					LabelName:  "git.branch",
					LabelValue: "master",
					MatchMode:  "exact",
				},
			},
		},
	},
}
