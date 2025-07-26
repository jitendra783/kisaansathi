package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kisaanSathi/pkg/config"
	"kisaanSathi/pkg/logger"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	resty "github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type RestCaller interface {
	//invokes a rest call based on method using resty library
	//		method: POST/GET
	//		url: string
	//		body: json, xml, []byte
	//		headers: map[key]value
	//		timeout: api timeoutvalue. default: 500ms
	//		auth: map[key]value, keys expected-> username, password, token
	//		queryparams: map[key]value
	// 		pathparams: map[key]value
	//		formData: map[string]string
	InvokeResty(c *gin.Context, method string, url string, body interface{}, headers map[string]string, timeout int64, auth map[string]string, queryparams map[string]string, pathparams map[string]string, formData map[string]string) ([]byte, int, error)
	//invokes a rest call based on method using native http library
	//		method: POST/GET
	//		body: json, xml, []byte
	//		headers: map[key]value
	//		timeout: api timeoutvalue. default: 500ms
	//		auth: map[key]value, keys expected-> username, password, token
	InvokeHttp(c *gin.Context, method string, url string, body interface{}, headers map[string]string, timeout int64, auth map[string]string, queryparams map[string]string, pathparams map[string]string) ([]byte, int, error)
}

type restCall struct{}

var (
	RetryCount     = 1
	RetryWaitTime  = 100
	DefaultTimeout = 500
)

// rest calling interface initialization function
//
//	fetches the default retry count, retry timeout and retry wait time from config
func GetRestCaller() RestCaller {
	confg := config.GetConfig()
	//populate the default variables
	RetryCount = confg.GetInt("thirdparty.restapi.retrycount")
	RetryWaitTime = confg.GetInt("thirdparty.restapi.retrywaittime")
	DefaultTimeout = confg.GetInt("thirdparty.restapi.timeout")

	return &restCall{}
}

// uses the resty library to make a REST call
//
//	steps within function
//	- creates a request
//		- sets a default timeout of timeout arguments is < 0
//		- sets a default retry fetched from config
//		- sets a default retry wait time fetched from config
//		- sets exponential backoff jitter for retries
//	- adds headers (if any) received in argument
//	- sets the body (if any) parsed in the argument
//	- switches for the HTTP method requested (methods implemented: POST, GET)
//	- sends the request and fetches the response
//	- checks for error and extracts response body and status code
//	- returns response body, status code and error
func (p *restCall) InvokeResty(c *gin.Context, method string, url string, body interface{}, headers map[string]string, timeout int64, auth map[string]string, queryparams map[string]string, pathparams map[string]string, formData map[string]string) ([]byte, int, error) {
	logger.Log(c).Info("http requset info", zap.Any("method", method), zap.Any("url", url), zap.Any("body", body), zap.Any("headers", headers), zap.Any("timeout", timeout), zap.Any("queryParms", queryparams), zap.Any("pathParms", pathparams))
	var (
		request    *resty.Request
		response   *resty.Response
		resp       []byte
		err        error
		statusCode int
	)

	if timeout > 0 {
		request = resty.New().
			SetTimeout(time.Duration(timeout) * time.Millisecond).
			SetRetryCount(RetryCount).
			SetRetryWaitTime(time.Duration(time.Duration(RetryWaitTime)) * time.Millisecond).
			// Default (nil) implies exponential backoff with jitter
			SetRetryAfter(nil).
			R()
	} else {
		request = resty.New().
			SetTimeout(time.Duration(DefaultTimeout) * time.Millisecond).
			SetRetryCount(RetryCount).
			SetRetryWaitTime(time.Duration(time.Duration(RetryWaitTime)) * time.Millisecond).
			// Default (nil) implies exponential backoff with jitter
			SetRetryAfter(nil).
			R()
	}

	request = request.SetHeader("Content-Type", "application/json")
	//set authorization if available
	if len(auth) > 0 {
		if auth["username"] != "" {
			request = request.SetBasicAuth(auth["username"], auth["password"])
		} else if auth["token"] != "" {
			request = request.SetAuthToken(auth["token"])
		}
	}
	//add pathparams if available
	if len(pathparams) > 0 {
		for key, value := range pathparams {
			request = request.SetPathParam(key, value)
		}
	}
	//add queryparams
	if queryparams != nil {
		request = request.SetQueryParams(queryparams)
	}

	//set custom headers
	if len(headers) > 0 {
		for key, val := range headers {
			if len(strings.ReplaceAll(key, " ", "")) > 0 {
				request = request.SetHeader(key, val)
			}
		}
	}

	//Set form data
	if len(formData) > 0 {
		request = request.SetFormData(formData)
	}
	if body != "" {
		request = request.SetBody(body)
	}

	switch method {
	case resty.MethodGet:
		response, err = request.Get(url)
		if err == nil {
			resp = response.Body()
		}
		statusCode = response.StatusCode()
	case resty.MethodPost:
		response, err = request.Post(url)

		if err == nil {
			resp = response.Body()
		}
		statusCode = response.StatusCode()
	default:
		err = fmt.Errorf("invalid request method: %v", method)
	}

	return resp, statusCode, err
}

func (p *restCall) InvokeHttp(c *gin.Context, method string, url string, body interface{}, headers map[string]string, timeout int64, auth map[string]string, queryparams map[string]string, pathparams map[string]string) ([]byte, int, error) {
	logger.Log(c).Info("http requset info", zap.Any("method", method), zap.Any("url", url), zap.Any("body", body), zap.Any("headers", headers), zap.Any("timeout", timeout))
	var (
		req *http.Request
		err error
	)
	if body != nil {
		reqData, _ := json.Marshal(body)
		req, err = http.NewRequest(method, url, bytes.NewBuffer(reqData))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	//set authorization if available
	if len(auth) > 0 {
		if auth["username"] != "" {
			req.SetBasicAuth(auth["username"], auth["password"])
		} else if auth["token"] != "" {
			req.Header.Add("Authorization", "Bearer "+auth["token"])
		}
	}
	//add pathparams if available
	if len(pathparams) != 0 {
		for k, v := range pathparams {
			req.URL.Path = strings.Replace(req.URL.Path, "{"+k+"}", v, -1)
		}
	}
	//add querparams if available
	if queryparams != nil {
		query := req.URL.Query()
		for k, v := range queryparams {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	cli := &http.Client{Timeout: time.Millisecond * time.Duration(timeout)}
	if strings.Contains(url, "trendlyne.com") {
		cli.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		logger.Log(c).Info("trendlyne api will be called without ssl verification")
	}
	resp, err := cli.Do(req)
	/*
		a, b := req.GetBody()
		c, d := ioutil.ReadAll(a)
		fmt.Println("REQUEST-------- ", a, b, string(c), d)
	*/
	if err != nil {
		if strings.Contains(err.Error(), "Client.Timeout") { //handling timeout
			return nil, http.StatusGatewayTimeout, nil
		}
		if resp == nil {
			return nil, http.StatusNotFound, err
		}
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return resBody, resp.StatusCode, nil
}
