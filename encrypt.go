package aliyun_nlp_sdk

import (
  "crypto/hmac"
  "crypto/md5"
  "crypto/sha1"
  "encoding/base64"
)

func md5Base64(body []byte) string {
  h := md5.New()
  h.Write(body)
  md5Str := h.Sum(nil)
  r := base64.StdEncoding.EncodeToString([]byte(md5Str))
  return r
}

func hmacSha1(keyStr string, value string) string {
  key := []byte(keyStr)
  mac := hmac.New(sha1.New, key)
  mac.Write([]byte(value))
  //进行base64编码
  res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
  return res
}
