package offsum

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestQuery(t *testing.T) {
	conf := viper.New()
	conf.SetConfigFile("../../conf/demo.yaml")
	err := conf.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	offsumClient := NewOffsumClient(conf)
	page := offsumClient.Query("http://www.zhihu.com/")

	for k, v := range page.Datas {
		fmt.Println("pages:", k)
		fmt.Println(v.Data)
	}
}
