package servers

import (
	context "context"
	"fmt"
	"time"

	"github.com/emil-petras/db-service/models"
	tokenProto "github.com/emil-petras/project-proto/token"
)

type TokenServer struct {
	tokenProto.UnimplementedTokenServiceServer
}

func (s *TokenServer) Create(ctx context.Context, in *tokenProto.CreateToken) (*tokenProto.Token, error) {
	token := models.Token{
		Value:      in.Value,
		UserID:     uint(in.UserID),
		ValidUntil: in.ValidUntil.AsTime(),
		Created:    time.Now().Local(),
	}

	result := models.DB.Create(&token)
	if result.RowsAffected == 0 {
		return &tokenProto.Token{}, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("failed to create token: %w", result.Error)
	}

	return token.ToProto(), nil
}

func (s *TokenServer) Read(ctx context.Context, in *tokenProto.ReadToken) (*tokenProto.Token, error) {
	token := models.Token{}

	result := models.DB.Where("value = ?", in.Value).Preload("User").Take(&token)
	if result.RowsAffected == 0 {
		return &tokenProto.Token{}, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("failed to read token: %w", result.Error)
	}

	return token.ToProto(), nil
}
