package http

import (
	"io/ioutil"
	"net/http"
	"errors"
	"encoding/json"
	"mime/multipart"
	"bytes"
	"io"
	"time"
)

const CONTENT_TYPE_JSON = "application/json"

func Request(requestParam map[string]interface{}, matchRoute string) (map[string]interface{}, error) {
	// get request
	req, err := getParams(requestParam, matchRoute)
	if err != nil {
		return nil, err
	}

	// create http client and exec request
	client := http.Client{}
	reps, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer reps.Body.Close()

	re, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		return nil, err
	}

	repData := map[string]interface{}{}
	json.Unmarshal(re, &repData)
	return repData, nil
}

func getParams(requestParam map[string]interface{}, matchRoute string) (*http.Request, error) {

	method := ""
	query := ""
	var body, requestContent io.Reader // json data
	header := http.Header{}
	multipartData := &multipart.Form{}

	var err error
	if v, ok := requestParam["method"]; ok {
		method = v.(string)
	}
	if v, ok := requestParam["query"]; ok {
		query = v.(string)
	}
	if v, ok := requestParam["header"]; ok {
		header = v.(http.Header)
	}
	if v, ok := requestParam["multipart"]; ok {
		multipartData = v.(*multipart.Form)
	}
	if v, ok := requestParam["body"]; ok {
		body = v.(io.ReadCloser)
	}

	if method == "" || matchRoute == "" {
		return nil, errors.New("request api error")
	}

	contentType, ok := header["Content-Type"]
	if ok {
		switch contentType[0] {
		case CONTENT_TYPE_JSON:
			requestContent = body
		default:
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			// if has file
			if multipartData.File != nil {

				for k, v := range multipartData.File {

					for index, fh := range v {
						file, err := v[index].Open()
						defer file.Close()

						part, err := writer.CreateFormFile(k, fh.Filename)
						if err != nil {
							return nil, err
						}
						_, err = io.Copy(part, file)
					}
				}
			}

			for key, val := range multipartData.Value {
				_ = writer.WriteField(key, val[0])
			}
			err = writer.Close()
			if err != nil {
				return nil, err
			}
			header["Content-Type"] = []string{writer.FormDataContentType()}
			requestContent = body
		}
	}

	req, err := http.NewRequest(method, matchRoute+query, requestContent)
	if err != nil {
		return nil, errors.New("request error")
	}
	req.Header = header

	return req, err
}

func CheckNode(node string) bool {
	timeout := time.Duration(100 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(node)
	if err != nil {
		return false
	}
	return true
}
