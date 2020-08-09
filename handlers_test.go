package aliyun_nlp_sdk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	AccessKeyId     = "xxx"
	AccessKeySerect = "xxx"
	RegionId        = "cn-shanghai"
)

func TestNlpClient_ExecApi(t *testing.T) {
	client := NewNlpClient(AccessKeyId, AccessKeySerect, RegionId)
	t.Run("wordsegment", func(t *testing.T) {
		resp, err := client.ExecApi("wordsegment",
			[]byte(`{"lang":"ZH","text":"Iphone专用数据线"}`),
			"general",
		)
		data := readRespData(resp)
		assert.Empty(t, err)
		assert.Equal(
			t,
			`{"data":[{"id":0,"word":"Iphone"},{"id":1,"word":"专用"},{"id":2,"word":"数据线"}]}`,
			string(data))
	})

	t.Run("wordpos", func(t *testing.T) {
		resp, err := client.ExecApi("wordpos",
			[]byte(`{"text":"真丝韩都衣舍连衣裙"}`),
			"general",
		)
		data := readRespData(resp)
		assert.Empty(t, err)
		assert.Equal(
			t,
			`{"data":[{"pos":"NN","word":"真丝"},{"pos":"NR","word":"韩都衣舍"},{"pos":"NN","word":"连衣裙"}]}`,
			string(data))
	})

	t.Run("entity", func(t *testing.T) {
		resp, err := client.ExecApi("entity",
			[]byte(`{"text":"真丝韩都衣舍连衣裙","type":"full"}`),
			"ecommerce",
		)
		data := readRespData(resp)
		assert.Empty(t, err)
		assert.Equal(
			t,
			`{"data":[{"synonym":"","tag":"材质","weight":"0.6","word":"真丝"},{"synonym":"","tag":"品牌","weight":"0.8","word":"韩都衣舍"},{"synonym":"连身裙;联衣裙","tag":"品类","weight":"1.0","word":"连衣裙"}]}`,
			string(data))
	})

	t.Run("sentiment", func(t *testing.T) {
		resp, err := client.ExecApi("sentiment",
			[]byte(`{"text": "虽然有点贵，不是很修身，但颜色很亮，布料摸起来挺舒服的，图案也好看，挺喜欢的。"}`),
			"ecommerce",
		)
		data := readRespData(resp)
		assert.Empty(t, err)
		assert.Equal(
			t,
			`{"data":{"text_polarity":"0"}}`,
			string(data))
	})

	t.Run("kwe", func(t *testing.T) {
		resp, err := client.ExecApi("kwe",
			[]byte(`{"lang":"ZH","text":"新鲜桔子"}`),
			"ecommerce",
		)
		data := readRespData(resp)
		assert.Empty(t, err)
		assert.Equal(
			t,
			`{"data":["桔子"]}`,
			string(data))
	})

	t.Run("textstructure", func(t *testing.T) {
		resp, err := client.ExecApi("textstructure",
			[]byte(`{"text":"脚蹬Mra，帅里帅气Mra是2013年新崛起的新锐品牌，作为Mra的大Boss福叔他爱鞋如痴，从皮料到包装他都严格把关，当收集到足够多的意见后，他总会用纯手工打扮出第一双样鞋，然后再不断 的调整改进，>每双都能精益求精！福叔常说要用品质为顾客撑腰，因此Mra都是选用的上等小牛皮加工，由经验丰富的老工匠在一边亲自操刀。","tag_flag":"true"}`),
			"ecommerce",
		)
		data := readRespData(resp)
		assert.Empty(t, err)
		assert.Equal(
			t,
			`{"data":{"tags":"鞋:产品词;上等:产品修饰词;Mra:普通词;选用:产品修饰词;崛起:机构实体;新锐:产品修饰词;老工匠:产品词;精益求精:机构实体;福叔:普通词;帅气:产品修饰词","label_name":"运动/户外"}}`,
			string(data))
	})

	t.Run("reviewanalysis", func(t *testing.T) {
		resp, err := client.ExecApi("reviewanalysis",
			[]byte(`{"text":"面料舒适，款式好，只是尺码偏小，好在我看了其他买家的评价，在原尺码上加了一号，正合适，很满意！给满分！服务好，发货快！","cate":"clothing"}`),
			"ecommerce",
		)
		data := readRespData(resp)
		assert.Empty(t, err)
		assert.Equal(
			t,
			`{"data":{"aspectItem":[{"aspectCategory":"面料/材质","aspectIndex":"0 2","aspectTerm":"面料","opinionTerm":"舒适","aspectPolarity":"正"},{"aspectCategory":"版型/款式","aspectIndex":"5 7","aspectTerm":"款式","opinionTerm":"好","aspectPolarity":"正"},{"aspectCategory":"尺码","aspectIndex":"11 13","aspectTerm":"尺码","opinionTerm":"偏小","aspectPolarity":"负"},{"aspectCategory":"卖家服务","aspectIndex":"51 53","aspectTerm":"服务","opinionTerm":"服务好","aspectPolarity":"正"}],"textIntensity":"0.5330","textPolarity":"正"}}`,
			string(data))
	})


}
func readRespData(resp *http.Response) []byte {
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(data))
	return data
}
