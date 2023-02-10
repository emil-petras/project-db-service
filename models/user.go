package models

import (
	"time"

	userProto "github.com/emil-petras/project-proto/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	ID       uint `gorm:"primary_key;autoIncrement;<-:create"`
	Username string
	Password string
	Balance  uint
	Created  time.Time `gorm:"not null"`
	Updated  time.Time `gorm:"not null"`
}

func (user *User) ToProto() *userProto.User {
	return &userProto.User{
		Id:       uint64(user.ID),
		Username: user.Username,
		Password: user.Password,
		Balance:  uint64(user.Balance),
		Updated:  timestamppb.New(user.Updated),
		Created:  timestamppb.New(user.Created),
	}
}
