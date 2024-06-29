package types

import (
	"fmt"
	"strings"
)

type Certification struct {
	Id          string
	IssuerId    string
	LearnerId   string
	Title       string
	Description string
	IssueDate   int64
	ExpiryDate  int64 // Optional, for certifications that expire
	Skills      []string
	Hash        string
	Creator     string
}

func (c *Certification) Reset() {
	*c = Certification{}
}

func (c *Certification) String() string {
	return fmt.Sprintf(`Certification:
    ID: %s
    Issuer ID: %s
    Learner ID: %s
    Title: %s
    Description: %s
    Issue Date: %d
    Expiry Date: %d
    Skills: %s
    Hash: %s
    Creator: %s`,
		c.Id, c.IssuerId, c.LearnerId, c.Title, c.Description,
		c.IssueDate, c.ExpiryDate, strings.Join(c.Skills, ", "),
		c.Hash, c.Creator)
}

func (*Certification) ProtoMessage() {}
