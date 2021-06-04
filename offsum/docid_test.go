package offsum

import (
	"fmt"
	"testing"
)

func TestDocid(t *testing.T) {
	docid, _ := Docid("http://www.zhihu.com/")
	fmt.Println("docid:", docid)
}

func TestUrlID(t *testing.T) {
	urlid, _ := UrlID("http://www.zhihu.com/")
	fmt.Println("urlid:", urlid)
}
