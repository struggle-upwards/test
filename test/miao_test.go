package test

import (
	"test/function"
	"testing"
)

func TestMoreSplit(t *testing.T) {
	t.Log("开始测试miao叫")
	if function.Miao() != "Miao?" {
		t.Errorf("miao叫不正常")
	}
}
