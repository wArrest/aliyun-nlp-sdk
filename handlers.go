package aliyun_nlp_sdk

import (
  "encoding/json"
  "fmt"
  "net/http"
)

func (n NlpClient) WordSegment(text string, lang string, domain string) (*http.Response, error) {
  body := struct {
    Text string `json:"text"`
    Lang string `json:"lang"`
  }{
    Text: text,
    Lang: lang,
  }
  bodyBytes, err := json.Marshal(body)
  if err != nil {
    return nil, err
  }

  var urlStr = n.baseUrl + "wordsegment/" + domain

  resp, err := n.send(Method, urlStr, bodyBytes)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func (n NlpClient) WordPos(text string, domain string) (*http.Response, error) {
  wordPosBody := struct {
    Text string `json:"text"`
  }{
    Text: text,
  }
  wordPosBodyBytes, _ := json.Marshal(wordPosBody)

  urlStr := n.baseUrl + "wordpos/" + domain

  resp, err := n.send(Method, urlStr, wordPosBodyBytes)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func (n NlpClient) Entity(text string, style string, domain string) (*http.Response, error) {
  body := struct {
    Text  string `json:"text"`
    Style string `json:"type"`
  }{
    Text:  text,
    Style: style,
  }
  bodyBytes, err := json.Marshal(body)
  if err != nil {
    return nil, err
  }

  urlStr := n.baseUrl + fmt.Sprintf("entity/%s", domain)
  resp, err := n.send(Method,urlStr,bodyBytes)
  if err != nil {
    return nil, err
  }
  return resp, nil

}
