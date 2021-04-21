package easynet

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	simplejson "github.com/nsingxia/go-simplejson"
)

const (
	POST = "POST"
	GET  = "GET"
)

// HttpEx is HttpEx
func HttpRaw(method, url string, data io.Reader, headerFun func(request *http.Request)) ([]byte, error) {
	client := &http.Client{}

	reqest, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	if headerFun != nil {
		headerFun(reqest)
	}

	resp, err := client.Do(reqest)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	encStr := resp.Header.Get("content-encoding")
	if strings.Contains(encStr, "gzip") {

	}

	return unGzip(body)
}

func HttpEx(method, url string, data io.Reader, headerFun func(request *http.Request)) ([]byte, error) {
	client := &http.Client{}

	reqest, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	if headerFun != nil {
		headerFun(reqest)
	}

	resp, err := client.Do(reqest)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func HttpJson(method, url string, data io.Reader, headerFun func(request *http.Request)) (*simplejson.Json, error) {
	b, err := HttpRaw(method, url, data, headerFun)
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(b)
}

func HttpExAll(method, url string, data io.Reader, headerFun func(request *http.Request)) (*http.Response, error) {
	client := &http.Client{}

	reqest, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	if headerFun != nil {
		headerFun(reqest)
	}

	resp, err := client.Do(reqest)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func unGzip(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(r)
}
