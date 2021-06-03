# go-offsum
Http client to access offsum

## Usage

```go
conf := viper.New()
conf.SetConfigFile("conf/demo.yaml")
conf.ReadInConfig()
offsumClient := NewOffsumClient(conf)
page := offsumClient.Query("http://www.zhihu.com/")
```