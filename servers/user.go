package servers

import (
	context "context"
	"fmt"
	"time"

	"github.com/emil-petras/db-service/models"
	userProto "github.com/emil-petras/project-proto/user"
)

type UserServer struct {
	userProto.UnimplementedUserServiceServer
}

func (s *UserServer) Create(ctx context.Context, in *userProto.CreateUser) (*userProto.User, error) {
	user := models.User{
		Username: in.Username,
		Password: in.Password,
		Balance:  0,
		Updated:  time.Now().Local(),
		Created:  time.Now().Local(),
	}

	result := models.DB.Create(&user)
	if result.RowsAffected == 0 {
		return &userProto.User{}, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("failed to create user: %w", result.Error)
	}

	return user.ToProto(), nil
}

func (s *UserServer) Read(ctx context.Context, in *userProto.ReadUser) (*userProto.User, error) {
	user := models.User{}

	result := models.DB.Where("username = ?", in.Username).Take(&user)
	if result.RowsAffected == 0 {
		return &userProto.User{}, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("failed to read user: %w", result.Error)
	}

	return user.ToProto(), nil
}

func (s *UserServer) Update(ctx context.Context, in *userProto.UpdateUser) (*userProto.User, error) {
	user := models.User{
		Username: in.Username,
		Balance:  uint(in.Balance),
		Updated:  time.Now().Local(),
	}

	tx := models.DB.Begin()
	result := tx.Where("username = ?", in.Username).Select("balance", "updated").Updates(&user)
	if result.RowsAffected == 0 {
		return &userProto.User{}, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("failed to update user: %w", result.Error)
	}

	transaction := models.Transaction{
		Amount:  int(in.Balance),
		UserID:  uint(in.UserID),
		Created: time.Now().Local(),
	}

	result = tx.Create(&transaction)
	if result.RowsAffected == 0 {
		tx.Rollback()
		return &userProto.User{}, nil
	}

	if result.Error != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update transaction: %w", result.Error)
	}

	tx.Commit()

	return user.ToProto(), nil
}
