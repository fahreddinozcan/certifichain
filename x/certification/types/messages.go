package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

var _ sdk.Msg = &MsgIssueCertification{}

type MsgIssueCertification struct {
	Id          string
	IssuerId    string
	LearnerId   string
	Title       string
	Description string
	IssueDate   int64
	ExpiryDate  int64
	Skills      []string
	Hash        string
	Creator     string
}

func (m *MsgIssueCertification) Reset() {
	*m = MsgIssueCertification{}
}

// String returns a human-readable representation of the MsgIssueCertification.
func (m MsgIssueCertification) String() string {
	return fmt.Sprintf(`Issue Certification:
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
		m.Id, m.IssuerId, m.LearnerId, m.Title, m.Description,
		m.IssueDate, m.ExpiryDate, strings.Join(m.Skills, ", "),
		m.Hash, m.Creator)
}

// ProtoMessage is a no-op method to fulfill the proto.Message interface.
func (*MsgIssueCertification) ProtoMessage() {}

func NewMsgIssueCertification(id, issuerId, learnerId, title, description string, issueDate, expiryDate int64, skills []string, hash, creator string) *MsgIssueCertification {
	return &MsgIssueCertification{
		Id:          id,
		IssuerId:    issuerId,
		LearnerId:   learnerId,
		Title:       title,
		Description: description,
		IssueDate:   issueDate,
		ExpiryDate:  expiryDate,
		Skills:      skills,
		Hash:        hash,
		Creator:     creator,
	}
}

// MsgIssueCertificationResponse defines the response for IssueCertification message
type MsgIssueCertificationResponse struct{}
