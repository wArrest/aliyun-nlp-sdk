package aliyun_nlp_sdk

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNlpClient_initSignature(t *testing.T) {
	client := NewNlpClient(AccessKeyId, AccessKeySerect, RegionId)
	body := []byte(`{"text":"Iphone专用数据线"}`)
	req, _ := http.NewRequest(Method, "https://nlp.cn-shanghai.aliyuncs.com/nlp/api/xxx", bytes.NewReader(body))
	_, err := client.initSignature(req,body)

	assert.Empty(t, err)
}
