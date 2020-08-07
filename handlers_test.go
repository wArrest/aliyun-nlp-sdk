package aliyun_nlp_sdk

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const   (
	AccessKeyId  = "Your AccessKeyId"
	AccessKeySerect = "Your AccessKeySecret"
	RegionId = "cn-shanghai"
)
func TestNlpHandlers(t *testing.T) {
	nlpClient := NewNlpClient(AccessKeyId, AccessKeySerect, RegionId)

	t.Run("WordPos", func(t *testing.T) {
		resp, _ := nlpClient.WordPos("真丝韩都衣舍连衣裙", "general")
		respData, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(
			t,
			`{"data":[{"pos":"NN","word":"真丝"},{"pos":"NR","word":"韩都衣舍"},{"pos":"NN","word":"连衣裙"}]}`,
			string(respData),
		)
	})

	t.Run("WordSegment", func(t *testing.T) {
		resp, _ := nlpClient.WordSegment("Iphone专用数据线", "ZH", "general")
		respData, _ := ioutil.ReadAll(resp.Body)

		assert.Equal(t,
			[]byte(`{"data":[{"id":0,"word":"Iphone"},{"id":1,"word":"专用"},{"id":2,"word":"数据线"}]}`),
			respData)
	})

	t.Run("Entity", func(t *testing.T) {
		resp, _ := nlpClient.Entity("真丝韩都衣舍连衣裙", "full", "ecommerce")
		respData, _ := ioutil.ReadAll(resp.Body)

		want := []byte(`{"data":[{"synonym":"","tag":"材质","weight":"0.6","word":"真丝"},{"synonym":"","tag":"品牌","weight":"0.8","word":"韩都衣舍"},{"synonym":"连身裙;联衣裙","tag":"品类","weight":"1.0","word":"连衣裙"}]}`)

		assert.Equal(t, string(want), string(respData))
	})

}
