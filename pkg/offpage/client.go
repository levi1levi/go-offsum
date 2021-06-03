package offpage

import (
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/spf13/viper"
	"time"
)

type OffsumClient struct {
	Client  *httpclient.Client
	BaseUrl string
	Timeout int64
}

func NewOffsumClient(conf *viper.Viper) *OffsumClient {
	client := new(OffsumClient)
	client.BaseUrl = conf.GetString("offsum.baseUrl")
	client.Timeout = conf.GetInt64("offsum.timeout")
	timeout := time.Duration(client.Timeout) * time.Millisecond
	client.Client = httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	return client
}

func (this *OffsumClient) Query(url string) *OffPage {
	urlid, _ := UrlID(url)
	res, err := this.Client.Get(this.BaseUrl+urlid, nil)
	if err != nil {
		panic(err)
	}
	return ReadPage(res.Body)
}
