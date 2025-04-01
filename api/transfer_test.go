package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mockdb "github.com/puzzaney/simplebank/db/mock"
	db "github.com/puzzaney/simplebank/db/sqlc"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateTransfer(t *testing.T) {
	account1 := randomAccount()
	account1.Currency = "USD"

	account2 := randomAccount()
	account2.Currency = "USD"

	arg := db.TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        10,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := httptest.NewRecorder()

	var transfer db.TransferTxResult

	store := mockdb.NewMockStore(ctrl)

	server := NewServer(store)

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Return(account1, nil)
	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Return(account2, nil)
	store.EXPECT().TransferTx(gomock.Any(), gomock.Eq(arg)).Times(1).Return(transfer, nil)

	transferRequest := createTransferRequest{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        10,
		Currency:      "USD",
	}

	transferRequestJson, err := json.Marshal(transferRequest)
	require.NoError(t, err)

	request := httptest.NewRequest(http.MethodPost, "/transfers", strings.NewReader(string(transferRequestJson)))

	server.router.ServeHTTP(recorder, request)

}
