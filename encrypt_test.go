package aliyun_nlp_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_md5Base64(t *testing.T) {
	got := md5Base64([]byte("123"))
	want := "ICy5YqxZB1uWSwcVLSNLcA=="
	assert.Equal(t, want, got)
}


