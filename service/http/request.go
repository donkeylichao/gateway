package http

import (
	"io/ioutil"
	"net/http"
	"errors"
	"encoding/json"
	"io"
	"net/url"
	"mime/multipart"
)

func Request(requestParam map[string]interface{},matchRoute string) (map[string]interface{}, error) {

	method := ""
	query := ""
	var body io.ReadCloser
	header := http.Header{}
	multipartData := &multipart.Form{}
	form := url.Values{}
	var err error
	if v,ok := requestParam["method"];ok {
		method = v.(string)
	}
	if v,ok := requestParam["query"];ok {
		query = v.(string)
	}
	if v,ok := requestParam["header"];ok {
		header = v.(http.Header)
	}
	if v,ok := requestParam["multipart"];ok {
		multipartData = v.(*multipart.Form)
	}
	if v,ok := requestParam["body"];ok {
		body = v.(io.ReadCloser)
	}
	if v,ok := requestParam["form"];ok {
		form = v.(url.Values)
	}

	if method == "" || matchRoute == "" {
		return nil, errors.New("request api error")
	}

	req, err := http.NewRequest(method, matchRoute + query, body)
	if err != nil {
		return nil, errors.New("request error")
	}

	req.Header = header
	req.MultipartForm = multipartData
	req.Form = form

	// create http client and exec request
	client := http.Client{}
	reps, err := client.Do(req)
	if err != nil {
		return nil ,err
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
