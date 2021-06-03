package offsum

import (
	"github.com/cfhamlet/os-go-docid/docid"
)

func Docid(url string) (string, error) {
	docid, err := docid.New(url)
	if err != nil {
		return "", err
	}
	return docid.String(), err
}

func UrlID(url string) (string, error) {
	docid, err := docid.New(url)
	if err != nil {
		return "", err
	}
	return docid.URLID().String(), err
}
