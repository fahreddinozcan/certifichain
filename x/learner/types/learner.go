package types

import "fmt"

type Learner struct {
	Id        string
	Name      string
	Email     string
	PublicKey string
	Creator   string
}

func (l *Learner) Reset() {
	*l = Learner{}
}

func (l Learner) String() string {
	return fmt.Sprintf(`Learner:
    ID: %s
    Name: %s
    Email: %s
    Public Key: %s
    Creator: %s`,
		l.Id, l.Name, l.Email, l.PublicKey, l.Creator)
}

func (*Learner) ProtoMessage() {}
