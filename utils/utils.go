package utils

import (
	"net/url"
	"strings"
)

func FormatMessage(msg string) url.Values {
	// 1.根据" "将原始日志拆分成数组
	msgArr := strings.Split(msg, " ")

	// 2.所需数据是日志数组下标为6的字符串里"?"后边的值, 即根据"?"将该字符串拆分成数组, 取下标为1的值
	data := strings.Split(msgArr[6], "?")[1]

	// 3.由于所需数据当前格式为"k1=v1&k2=v2", 考虑将该字符串转化为键值对
	// 方法1: 使用自身url包ParseQuery方法
	ret, err := url.ParseQuery(data)
	if err != nil {
		panic(err)
	}
	return ret

	// 方法2: 根据"&"拆分字符串,再根据"="拆分具体数据构造键值对
	//entries := strings.Split(data, "&")
	//m := make(map[string]string)
	//for _, e := range entries {
	//	parts := strings.Split(e, "=")
	//	m[parts[0]] = parts[1]
	//}
	//return m
}
