package servers

import (
	context "context"
	"fmt"
	"time"

	"github.com/emil-petras/db-service/models"
	transProto "github.com/emil-petras/project-proto/transaction"
)

type TransactionServer struct {
	transProto.UnimplementedTransactionServiceServer
}

func (t *TransactionServer) Create(ctx context.Context, in *transProto.CreateTransaction) (*transProto.Transaction, error) {
	transaction := models.Transaction{
		Amount:  int(in.Amount),
		UserID:  uint(in.UserID),
		Created: time.Now().Local(),
	}

	result := models.DB.Create(&transaction)
	if result.RowsAffected == 0 {
		return &transProto.Transaction{}, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", result.Error)
	}

	return transaction.ToProto(), nil
}
