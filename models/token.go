package models

import (
	"time"

	tokenProto "github.com/emil-petras/project-proto/token"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Token struct {
	ID         uint `gorm:"primary_key;autoIncrement;<-:create"`
	Value      string
	UserID     uint
	User       User      `gorm:"foreignKey:UserID;references:ID"`
	ValidUntil time.Time `gorm:"not null"`
	Created    time.Time `gorm:"not null"`
}

func (token *Token) ToProto() *tokenProto.Token {
	return &tokenProto.Token{
		Id:         uint64(token.ID),
		Value:      token.Value,
		UserID:     uint64(token.UserID),
		Username:   token.User.Username,
		ValidUntil: timestamppb.New(token.ValidUntil),
		Created:    timestamppb.New(token.Created),
	}
}
