package sys_json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

/*
三方插件 json 高效率解决go解析不足
注意：本工程用这个

	added by yh @ 2023-04-25
*/
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Init() {
	extra.RegisterFuzzyDecoders() //开启兼容PHP模式
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
func UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}
