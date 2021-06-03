package offsum

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type PageData struct {
	Data           string `json:"data"`
	CompressedSize int    `json:"compressed_size"`
	OriginalSize   int    `json:"original_size"`
}

type OffPage struct {
	Header map[string]interface{} `json:"header"`
	Datas  map[string]*PageData   `json:"datas"`
}

func NewPageData(page *OffPage, pageType string) *PageData {
	data := new(PageData)
	contentMap := page.Header["Content-Type"].(map[string]string)
	compressedSize, _ := strconv.Atoi(contentMap[pageType])
	data.CompressedSize = compressedSize
	originalMap := page.Header["Original-Size"].(map[string]string)
	originalSize, _ := strconv.Atoi(originalMap[pageType])
	data.OriginalSize = originalSize
	return data
}

func ReadSize(b []byte) map[string]string {
	r := make(map[string]string)
	arrs := bytes.Split(bytes.TrimRight(bytes.TrimSpace(b), ";"), []byte(";"))
	for _, arr := range arrs {
		index := bytes.IndexByte(arr, ',')
		if index < 0 {
			continue
		}
		key := string(bytes.TrimSpace(arr[0:index]))
		value := string(bytes.TrimSpace(arr[index+1 : len(arr)]))
		r[key] = value
	}
	return r
}

func ReadData(r io.Reader, size int) (io.Reader, error) {
	buf := make([]byte, size)
	n, err := io.ReadFull(r, buf)
	if err != nil {
		fmt.Println("Read error!", err)
	}
	fmt.Println("length:", n)
	b := bytes.NewReader(buf)
	depressData, err := zlib.NewReader(b)
	defer depressData.Close()
	return depressData, err
}

func UTF16ToUTF8(r io.Reader) io.Reader {
	decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	decodeData := decoder.Reader(r)
	return decodeData
}

func ReadPage(r io.Reader) *OffPage {
	page := new(OffPage)
	page.Header = make(map[string]interface{})
	page.Datas = make(map[string]*PageData)
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil || io.EOF == err {
			break
		}
		index := bytes.IndexByte(line, ':')
		if index > 0 {
			key := bytes.TrimSpace(line[0:index])
			value := bytes.TrimSpace(line[index+1 : len(line)])
			fmt.Println("key:", strings.TrimSpace(string(key)), "value:", strings.TrimSpace(string(value)))
			if bytes.Equal(key, []byte("Content-Type")) || bytes.Equal(key, []byte("Original-Size")) {
				fmt.Println("find Content-Type")
				page.Header[string(key)] = ReadSize(value)
			} else {
				page.Header[string(key)] = string(value)
			}
		}
		if line[0] == '\n' {
			fmt.Println("header end")
			break
		}
		fmt.Print(string(line))
	}
	fmt.Println("Content-Type:")
	for k, v := range page.Header["Content-Type"].(map[string]string) {
		fmt.Println("key:", k, "value:", v)
	}
	fmt.Println("body:")
	contentMap := page.Header["Content-Type"].(map[string]string)
	// xmlpage
	xmlpageSize, _ := strconv.Atoi(contentMap["xmlpage"])
	depressXmlData, err := ReadData(reader, xmlpageSize)
	if err != nil {
		fmt.Println("error:", err)
	}
	decodeXmlData := UTF16ToUTF8(depressXmlData)

	xmlData, err := ioutil.ReadAll(decodeXmlData)
	fmt.Println("xmlpage depressed len", len(xmlData), "cap", cap(xmlData))
	xmlPage := NewPageData(page, "xmlpage")
	xmlPage.Data = string(xmlData)
	page.Datas["xmlpage"] = xmlPage

	// snapshot
	snapshotSize, _ := strconv.Atoi(contentMap["snapshot"])
	depressSnapshotData, err := ReadData(reader, snapshotSize)
	if err != nil {
		fmt.Println("error:", err)
	}

	snapshotData, err := ioutil.ReadAll(depressSnapshotData)
	fmt.Println("snapshot depressed len", len(snapshotData), "cap", cap(snapshotData))
	snapshotPage := NewPageData(page, "snapshot")
	snapshotPage.Data = string(snapshotData)
	page.Datas["snapshot"] = snapshotPage

	// renderpage
	renderpageSize, _ := strconv.Atoi(contentMap["renderpage"])
	depressrenderpageData, err := ReadData(reader, renderpageSize)
	if err != nil {
		fmt.Println("error:", err)
	}

	renderpageData, err := ioutil.ReadAll(depressrenderpageData)
	fmt.Println("renderpage depressed len", len(renderpageData), "cap", cap(renderpageData))
	renderPage := NewPageData(page, "renderpage")
	renderPage.Data = string(renderpageData)
	page.Datas["renderpage"] = renderPage


	return page
}
