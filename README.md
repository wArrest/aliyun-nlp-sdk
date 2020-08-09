# aliyun-nlp-sdk
sdk for aliyun nlp service:1.0

> 阿里云官方没有提供nlp:1.0版本的golang SDK，自己简单封装了一个

## Feature
- [x] 多语言分词 WordSegment
- [x] 词性标注 WordPos
- [x] 命名实体 Entity
- [x] 情感分析
- [x] 中心词提取
- [x] 智能文本分类
- [x] 文本信息抽取
- [x] 商品评价解析 

## Install
~~~ go
go get -u github.com/wArrest/aliyun-nlp-sdk
~~~
## Start
~~~go
client := NewNlpClient(AccessKeyId, AccessKeySerect, RegionId)
resp, err := client.ExecApi(
	        "wordsegment",
		[]byte(`{"lang":"ZH","text":"Iphone专用数据线"}`),
		"general",
		)
//handle resp ....
~~~
