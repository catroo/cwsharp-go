package simple

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"testing"
)

func Test(t *testing.T) {
	var input string = "Hello,World!你好，世界!"
	tokenizer := New()
	for iter := tokenizer.Traverse(cwsharp.ReadString(input)); iter.Next(); {
		fmt.Println(iter.Cur())
	}
}
