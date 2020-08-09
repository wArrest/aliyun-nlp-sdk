package aliyun_nlp_sdk

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const (
	BaseUrl             = "https://nlp.%s.aliyuncs.com/nlp/api/"
	Method              = "POST"
	Accept              = "application/json"
	ContentType         = "application/json;chrset=utf-8"
	AcsSignnatureMethod = "HMAC-SHA1"
	AcsVersion          = "2018-04-08"
	HeaderSeparator     = "\n"
)

type headerMap map[string]string

type NlpClient struct {
	client          *http.Client
	accessKeyId     string
	accessKeySecret string
	baseUrl         string
}

func NewNlpClient(accessKeyId string, accessKeySecret string, regionId string) *NlpClient {
	httpClient := http.Client{}
	return &NlpClient{
		client:          &httpClient,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		baseUrl:         fmt.Sprintf(BaseUrl, regionId),
	}
}

func (n NlpClient) send(method string, urlStr string, bodyContent []byte) (*http.Response, error) {

	req, err := n.buildRequest(method, urlStr, bodyContent)
	if err != nil {
		return nil, err
	}

	resp, err := n.client.Do(req)
	if err != nil {
		panic(err)
		return nil, err
	}

	return resp, nil
}
func (n NlpClient) buildRequest(method string, urlStr string, body []byte) (*http.Request, error) {
	req, _ := http.NewRequest(method, urlStr, bytes.NewReader(body))

	headerMap, err := n.initSignature(req, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	return req, nil
}
func (n NlpClient) initSignature(req *http.Request, body []byte) (headerMap, error) {
	date := time.Now().UTC().Format(http.TimeFormat)

	bodyMd5 := md5Base64(body)

	path := req.URL.Path
	uuidStr := uuid.New().String()

	//对body做MD5+BASE64加密
	stringToSign := Method + HeaderSeparator +
		Accept + HeaderSeparator +
		bodyMd5 + HeaderSeparator +
		ContentType + HeaderSeparator +
		date + HeaderSeparator +
		"x-acs-signature-method:HMAC-SHA1" + HeaderSeparator +
		"x-acs-signature-nonce:" + uuidStr + "\n" +
		"x-acs-version:2018-04-08" + HeaderSeparator +
		path

	//计算HMAC-SHA1
	var signature = hmacSha1(n.accessKeySecret, stringToSign)

	//得到 authorization header
	var authHeader = "acs " + n.accessKeyId + ":" + signature

	headerMap := headerMap{
		"Accept":                 Accept,
		"Content-Type":           ContentType,
		"Content-MD5":            bodyMd5,
		"Date":                   date,
		"Host":                   req.URL.Host,
		"Authorization":          authHeader,
		"x-acs-signature-nonce":  uuidStr,
		"x-acs-signature-method": AcsSignnatureMethod,
		"x-acs-version":          AcsVersion,
	}

	return headerMap, nil
}
