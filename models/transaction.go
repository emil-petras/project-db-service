package models

import (
	"time"

	transProto "github.com/emil-petras/project-proto/transaction"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Transaction struct {
	ID      uint `gorm:"primary_key;autoIncrement;<-:create"`
	Amount  int
	UserID  uint
	User    User      `gorm:"foreignKey:UserID"`
	Created time.Time `gorm:"not null"`
}

func (transaction *Transaction) ToProto() *transProto.Transaction {
	return &transProto.Transaction{
		Id:       uint64(transaction.ID),
		Amount:   int32(transaction.Amount),
		Username: transaction.User.Username,
		Created:  timestamppb.New(transaction.Created),
	}
}
