package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/umi7/DezHab_user_backend/common"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	errorhandler "gitlab.com/umi7/DezHab_user_backend/error-handler"
	"gitlab.com/umi7/DezHab_user_backend/service/impl"
	"gitlab.com/umi7/DezHab_user_backend/utils"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

const (
	apiKeyHeader      = "x-api-key"
	AuthenticationErr = "authentication Error"
	ErrApiKeyMissing  = "x-api-key is missing in the headers"
)

// Response is written to http.ResponseWriter
type Response struct {
	Code    int
	Payload interface{}
}

type RestHandler func(ctx context.Context, req *http.Request, request common.Validator) Response

type RequestHandler func(req *http.Request) Response

//httpContext is used to prepare the context for a given request
func httpContext(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), constants.RequestID, utils.UUID())
	return ctx
}

// validateRequest is used to authenticate the request based on api-key.
// validation will be performed only if the dto implements authenticator interface
var validateRequest = func(ctx context.Context, req *http.Request, request common.Validator) (context.Context, error) {
	if _, ok := request.(common.Authenticator); !ok {
		return ctx, nil
	}
	apiKey := req.Header.Get(apiKeyHeader)
	if apiKey == "" {
		return ctx, errors.New(ErrApiKeyMissing)
	}
	apiClient := impl.ApiClient.GetClient(ctx, apiKey)

	if apiClient.Provider == "" || apiClient.SubTag == "" {
		err := errors.New(AuthenticationErr)
		return ctx, err
	}
	ctx = context.WithValue(ctx, constants.Provider, apiClient.Provider)
	ctx = context.WithValue(ctx, constants.SubTag, apiClient.SubTag)
	ctx = context.WithValue(ctx, constants.ClientID, apiClient.ID)
	logger.Info(ctx, fmt.Sprintf("Provider:%s, SubTag:%s, ClientID:%d", apiClient.Provider, apiClient.SubTag, apiClient.ID))
	return ctx, nil
}

func getParam(values url.Values) string {
	res := ""
	init := 0
	for k, v := range values {
		if init > 0 {
			res += ","
		}
		res += fmt.Sprintf("\"%s\":\"%s\"", k, v[0])
		init++
	}
	return fmt.Sprintf("{%s}", res)
}

// getRequest will return the request data by reading the bytes data from http request
// and uses json unmarshal to deserialize.
func getRequest(ctx context.Context, req *http.Request, request common.Validator) (context.Context, error, errorhandler.ErrorHandler) {
	var err error
	ctx, err = validateRequest(ctx, req, request)
	if err != nil {
		return ctx, err, errorhandler.AuthenticationErr
	}

	reqBody, err := ioutil.ReadAll(req.Body)

	//Hack. TODO : Find a way to deserialize the request into struct in GET request
	if req.Method == http.MethodGet {
		reqBody = []byte(getParam(req.URL.Query()))
	}

	if err != nil {
		logger.Error(ctx, err.Error())
		return ctx, err, errorhandler.RequestErr
	}

	logger.Info(ctx, fmt.Sprintf("Request body : %s", string(reqBody)))
	ctx = context.WithValue(ctx, constants.RequestData, string(reqBody))
	err = json.Unmarshal(reqBody, request)
	if err != nil {
		logger.Error(ctx, err.Error())
		return ctx, err, errorhandler.RequestErr
	}

	err = request.Validate()
	if err != nil {
		logger.Error(ctx, err.Error())
		return ctx, err, errorhandler.ValidationErr
	}

	//Add book keeping entries only for Asynchronous processor
	if _, ok := request.(common.Asynchronous); ok {
		if err = impl.BookKeeping.Create(ctx, "Scheduled"); err != nil {
			return ctx, err, errorhandler.ServerErr
		}
	}

	return ctx, nil, nil
}

func MakeRest(request common.Validator, f RestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := httpContext(req)
		defer func(ctx context.Context) {
			if val, ok := utils.RecoverFromPanic(ctx, recover()); ok {
				writeErr(ctx, req, w, val)
			}
		}(ctx)
		//Creating a new object per request
		reqObjType := reflect.New(reflect.TypeOf(request))
		//This should always pass, so ignoring the ok value.
		reqObj := reqObjType.Interface().(common.Validator)
		var res Response
		ctx, err, errFunc := getRequest(ctx, req, reqObj)
		if err != nil {
			requestErr := errFunc(err)
			res = Response{
				Code:    requestErr.Code,
				Payload: requestErr,
			}
			goto end
		}
		res = f(ctx, req, reqObj)
	end:
		setupResponse(w)
		writeToResponseWrite(ctx, req, w, res)
	}
}

// Make creates a http handler from a request handler func
func Make(f RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func(ctx context.Context) {
			if val, ok := utils.RecoverFromPanic(ctx, recover()); ok {
				writeErr(ctx, req, w, val)
			}
		}(req.Context())
		setupResponse(w)
		if (*req).Method == "OPTIONS" {
			req.Body.Close()
			return
		}
		res := f(req)
		writeToResponseWrite(req.Context(), req, w, res)
	}
}

func writeErr(ctx context.Context, req *http.Request, w http.ResponseWriter, val interface{}) {
	response := Response{
		Code:    http.StatusServiceUnavailable,
		Payload: val,
	}
	setupResponse(w)
	writeToResponseWrite(ctx, req, w, response)
}

func writeToResponseWrite(ctx context.Context, req *http.Request, w http.ResponseWriter, response Response) {
	logger.Info(ctx, fmt.Sprintf("Http response is : %v", response))
	JSON, err := json.Marshal(response.Payload)
	if err != nil {
		logger.Error(ctx, err, "json marshal failed")
	}
	w.WriteHeader(response.Code)
	_, err = w.Write(JSON)
	if err != nil {
		logrus.Error(ctx, err, "json write failed")
	}
	err = req.Body.Close()
	if err != nil {
		logger.Error(ctx, err, "error while closing the body")
	}
}

func setupResponse(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Accept, x-requested-with, Content-Type, "+
		"Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-api-key")
}
