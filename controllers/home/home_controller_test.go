package home

import (
	"strconv"
	"testing"
)

/**
* @Author junfenghe
* @Description 测试 parInt
* @Date 2021-03-20 6:10
* @Param
* @return
**/
func TestParseInt(t *testing.T) {
	testStr := "junfenghe"
	testUint64, _ := strconv.ParseInt(testStr, 10, 32)
	t.Log(testUint64)
}
