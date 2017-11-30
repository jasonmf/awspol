package awspol

// http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type PolicyDocument struct {
	Version   string           `json:",omitempty"`
	ID        string           `json:"Id,omitempty"`
	Statement []StatementEntry `json:",omitempty"`
}

func (d PolicyDocument) Equals(o PolicyDocument) bool {
	if d.Version != o.Version {
		return false
	}
	if len(d.Statement) != len(o.Statement) {
		return false
	}
	for _, dStat := range d.Statement {
		matched := false
		for _, oStat := range o.Statement {
			if dStat.ExactlyEquals(oStat) {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func ParsePolicyDocument(s string) (PolicyDocument, error) {
	pd := PolicyDocument{}
	err := json.Unmarshal([]byte(s), &pd)
	if err != nil {
		return pd, errors.Wrap(err, "unmarshaling policy JSON")
	}
	return pd, nil
}
