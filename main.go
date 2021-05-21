package main

import (
	"context"

	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/server"
	"gitlab.com/umi7/DezHab_user_backend/utils"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
)

func main() {
	ctx := context.WithValue(context.Background(), constants.RequestID, utils.UUID())
	srv := server.New()
	logger.Info(ctx, "Server running in ", srv.Address)
	srv.ServeHTTP()
}

//go:generate swagger generate spec -o ./resources/swaggerui/swagger.json --scan-models
