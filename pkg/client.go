package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/erenkaratas99/COApiCore/pkg/customErrors"
	"github.com/erenkaratas99/COApiCore/pkg/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strconv"
)

type RestClient struct {
	Client *fasthttp.Client
}

var SingleRestClient *RestClient

func NewSingletonClient() *RestClient {
	SingleRestClient = &RestClient{Client: &fasthttp.Client{}}
	return SingleRestClient
}

func (c *RestClient) DoGetRequest(URI, cId string, respModel any) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(URI)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set(middleware.CorrelationID, cId)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := c.ProcessClientResponseData(req, resp, respModel)
	if err != nil {
		return err
	}
	return nil

}

func (c *RestClient) ProcessClientResponseData(req *fasthttp.Request, resp *fasthttp.Response, respModel any) error {
	if err := c.Client.Do(req, resp); err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		log.Info("Expected status code ", fasthttp.StatusOK, " but got ", resp.StatusCode())
		return customErrors.NewHTTPError(400, "ClientErr", "Expected status code 200 but got another")
	}
	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		log.Info("Expected content type application/json but got %s\n", contentType)
		return customErrors.NewHTTPError(400, "ClientErr", "Expected content type application/json but not arrived")
	}
	contentEncoding := resp.Header.Peek("Content-Encoding")
	var body []byte
	if bytes.EqualFold(contentEncoding, []byte("gzip")) {
		fmt.Println("Unzipping...")
		body, _ = resp.BodyGunzip()
	} else {
		body = resp.Body()
	}
	reader := bytes.NewReader(body)
	err := json.NewDecoder(reader).Decode(respModel)
	if err != nil {
		return err
	}
	return nil

}

func BuildLimitOffsetParams(limit, offset int) string {
	offsetStr := strconv.Itoa(offset)
	limitStr := strconv.Itoa(limit)
	//"/?limit=x&offset=y"
	urlParamStr := "/?limit=" + limitStr + "&" + "offset=" + offsetStr
	return urlParamStr
}
