package offsum

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestQuery(t *testing.T) {
	conf := viper.New()
	conf.SetConfigFile("../conf/demo.yaml")
	err := conf.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	offsumClient := NewOffsumClient(conf)
	page := offsumClient.Query("http://www.zhihu.com/")

	fmt.Println("Url:", page.Url)
	fmt.Println("Header:")
	for k, v := range page.Header {
		fmt.Println(k, ":", v)
	}

	for k, v := range page.Datas {
		fmt.Println("page:", k)
		fmt.Println("CompressedSize:", v.CompressedSize)
		fmt.Println("OriginalSize:", v.OriginalSize)
		fmt.Println(v.Data)
	}
}
