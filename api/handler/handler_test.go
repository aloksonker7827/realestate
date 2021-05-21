package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/umi7/DezHab_user_backend/common"
	"net/http"
	"testing"
)

func TestGetRequest(t *testing.T) {
	testCase := []struct {
		name                string
		mockValidateRequest func(context.Context, *http.Request, common.Validator) (context.Context, error)
		req                 func() *http.Request
		err1                error
		err2                *json.InvalidUnmarshalError
	}{
		{
			name: "x-apiKey-error",
			req: func() (request *http.Request) {
				request, _ = http.NewRequest("POST", "url", bytes.NewBuffer([]byte(`{}`)))
				return
			},
			mockValidateRequest: func(ctx context.Context, req *http.Request, request common.Validator) (cont context.Context, e error) {
				return ctx, errors.New(AuthenticationErr)
			},
			err1: errors.New(AuthenticationErr),
			err2: nil,
		},
		{
			name: "json-unmarshal-error",
			mockValidateRequest: func(ctx context.Context, req *http.Request, request common.Validator) (cont context.Context, e error) {
				return ctx, nil
			},
			req: func() (request *http.Request) {
				request, _ = http.NewRequest("POST", "url", bytes.NewBuffer([]byte(`{}`)))
				return
			},
			err1: nil,
			err2: &json.InvalidUnmarshalError{},
		},
	}

	validateRequestOld := validateRequest
	mockCtrl := gomock.NewController(t)

	defer func() {
		validateRequest = validateRequestOld
		mockCtrl.Finish()
	}()

	ctx := context.Background()
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			validateRequest = tc.mockValidateRequest
			_, err, _ := getRequest(ctx, tc.req(), nil)
			if tc.err1 != nil {
				assert.Equal(t, tc.err1, err)
			} else if tc.err2 != nil {
				assert.Equal(t, tc.err2, err)
			}
		})
	}
}

func TestMakeRestRequest(t *testing.T) {
	testCase := []struct {
		name            string
		mockRestHandler func(ctx context.Context, req *http.Request, request common.Validator) Response
		req             func() *http.Request
		err1            error
		err2            *json.InvalidUnmarshalError
	}{
		{
			name: "make rest call",
			req: func() (request *http.Request) {
				request, _ = http.NewRequest("POST", "url", bytes.NewBuffer([]byte(`{}`)))
				return
			},
			mockRestHandler: func(ctx context.Context, req *http.Request, request common.Validator) Response {
				return Response{}
			},
			err1: errors.New(AuthenticationErr),
			err2: nil,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			resp := MakeRest(nil, tc.mockRestHandler)
			if resp != nil {
				assert.Equal(t, Response{}, Response{})
			}
		})
	}
}
