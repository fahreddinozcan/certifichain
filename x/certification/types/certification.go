package types

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
