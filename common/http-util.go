package common

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPClient interface {
	Post(ctx context.Context, url string, requestData interface{}, responseData interface{}, httpHeaders ...map[string]string) error
	Get(ctx context.Context, url string, responseData interface{}, httpHeaders ...map[string]string) error
}

//go:generate mockgen -source=http-util.go -destination=mock-impl/http-client-impl.go -package=mock_impl

type HttpClient struct {
	http.Client
	Callee string
}

func (client HttpClient) getByteData(ctx context.Context, requestData interface{}) ([]byte, error) {
	logger.Info(ctx, requestData)
	request, err := json.Marshal(requestData)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	return request, nil
}

// Post makes the http post request with the given params by calling PerformHttpCall
//
// Parameters :
// 	ctx : http context of the request.
// 		This is used for logging purpose as of today (01-October-2019) and can be generalised
//  	to take advantage of http context as we do in other languages.
//	url : final url of the request.
//		Whatever data is passed here will be used to construct the http request.
//	requestData : Data to be send in the request body
// 		The data will be serialized using json marshal and transferred over the network as a bytes array
//	responseData: Data which will be received in the response.
//		The data will be de-serialized using json unmarshal. Here, the pointer to the struct should be passed
//		to the function, otherwise json.unmarshal will fail
// 	httpHeaders : Headers to be sent in the request.
//		This is an optional argument.
func (client HttpClient) Post(ctx context.Context, url string, requestData interface{}, responseData interface{}, httpHeaders ...map[string]string) error {
	var request []byte
	var err error
	if requestData != nil {
		request, err = client.getByteData(ctx, requestData)
	}
	if err != nil {
		return err
	}
	var req *http.Request
	if request != nil {
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(request))
	} else {
		req, err = http.NewRequest(http.MethodPost, url, nil)
	}
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	return client.PerformHttpCall(ctx, req, client.Callee, responseData, httpHeaders...)
}

// PerformHttpCall makes Http call for the request passed in argument.
func (client HttpClient) PerformHttpCall(ctx context.Context, req *http.Request, callee string, response interface{}, httpHeaders ...map[string]string) error {
	headers := make(map[string]string)
	if len(httpHeaders) != 0 {
		headers = httpHeaders[0]
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	// TODO : Content type header to be passed by caller in the httpHeaders map arg
	req.Header.Add("content-type", "application/json")
	start := time.Now()
	logger.Info(ctx, fmt.Sprintf("Calling %s with the request : %v", callee, req))
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	logger.Info(ctx, fmt.Sprintf("%s latency on API %s is : %v",
		client.Callee, req.URL.RequestURI(), time.Now().Sub(start).Seconds()))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	logger.Info(ctx, fmt.Sprintf("Response from %s is %s", callee, string(body)))
	err = json.Unmarshal(body, response)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	defer resp.Body.Close()
	//fixing an existing bug at this point for returning error if the status code is not 2XX
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		logger.Error(ctx, fmt.Sprintf("Status code returned by %s is not 200", client.Callee))
		return errors.New(fmt.Sprintf("Status code returned by %s is not 200", client.Callee))
	}

	return nil
}

func (client HttpClient) Get(ctx context.Context, url string, responseData interface{}, httpHeaders ...map[string]string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	return client.PerformHttpCall(ctx, req, client.Callee, responseData, httpHeaders...)
}
