package aliyun_nlp_sdk

import (
  "fmt"
  "net/http"
)

func (n NlpClient) ExecApi(name string, body []byte, domain string) (*http.Response, error) {
  //build requst url
  url := n.baseUrl + fmt.Sprintf("%s/%s", name, domain)

  resp, err := n.send(Method, url, body)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

//func (n NlpClient) WordSegment(text string, lang string, domain string) (*http.Response, error) {
//  body := struct {
//    Text string `json:"text"`
//    Lang string `json:"lang"`
//  }{
//    Text: text,
//    Lang: lang,
//  }
//  bodyBytes, err := json.Marshal(body)
//  if err != nil {
//    return nil, err
//  }
//
//  var urlStr = n.baseUrl + "wordsegment/" + domain
//
//  resp, err := n.send(Method, urlStr, bodyBytes)
//  if err != nil {
//    return nil, err
//  }
//  return resp, nil
//}
//
//func (n NlpClient) WordPos(text string, domain string) (*http.Response, error) {
//  wordPosBody := struct {
//    Text string `json:"text"`
//  }{
//    Text: text,
//  }
//  wordPosBodyBytes, _ := json.Marshal(wordPosBody)
//
//  urlStr := n.baseUrl + "wordpos/" + domain
//
//  resp, err := n.send(Method, urlStr, wordPosBodyBytes)
//  if err != nil {
//    return nil, err
//  }
//  return resp, nil
//}
//
//func (n NlpClient) Entity(text string, style string, domain string) (*http.Response, error) {
//  body := struct {
//    Text  string `json:"text"`
//    Style string `json:"type"`
//  }{
//    Text:  text,
//    Style: style,
//  }
//  bodyBytes, err := json.Marshal(body)
//  if err != nil {
//    return nil, err
//  }
//
//  urlStr := n.baseUrl + fmt.Sprintf("entity/%s", domain)
//  resp, err := n.send(Method, urlStr, bodyBytes)
//  if err != nil {
//    return nil, err
//  }
//  return resp, nil
//}
//
//func (n NlpClient) TextStructure(text string, domain string) (*http.Response, error) {
//  body := struct {
//    Text    string `json:"text"`
//    TagFlag string `json:"tag_flag"`
//  }{
//    Text:    text,
//    TagFlag: "true",
//  }
//
//  bodyBytes, err := json.Marshal(body)
//  if err != nil {
//    return nil, err
//  }
//
//  urlStr := n.baseUrl + fmt.Sprintf("textstructure/%s", domain)
//  resp, err := n.send(Method, urlStr, bodyBytes)
//  if err != nil {
//    return nil, err
//  }
//
//  return resp, nil
//}
