package api

import (
	"gitlab.com/umi7/DezHab_user_backend/api/handler"
	"gitlab.com/umi7/DezHab_user_backend/service/impl"
	"net/http"
)

// HealthCheck is checking health status of dezhab user backend service
//
// It checks health status of all dependencies of dezhab user backend except external services
var HealthCheck = func(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	// impl.HealthCheck.DeepHealthCheck checks health status of postGreSQL
	code, response := impl.HealthCheck.DeepHealthCheck(ctx)
	handler.Make(func(req *http.Request) handler.Response {
		return handler.Response{
			Code:    code,
			Payload: response,
		}
	})(w, req)
}
