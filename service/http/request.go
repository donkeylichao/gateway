package http

import (
	"github.com/dghubble/sling"
	"net/http"
	"fmt"
)

type Params struct {
	Count int `url:"count,omitempty"`
}

// IssueService provides methods for creating and reading issues.
type IssueService struct {
	sling *sling.Sling
}

// Client is a tiny Github client
type Client struct {
	IssueService *IssueService
	// other service endpoints...
}

type Issue struct {
	//ID     int    `json:"id"`
	//URL    string `json:"url"`
	//Number int    `json:"number"`
	//State  string `json:"state"`
	//Title  string `json:"title"`
	//Body   string `json:"body"`
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}


type GithubError struct {
	Message string `json:"message"`
	Errors  []struct {
		Resource string `json:"resource"`
		Field    string `json:"field"`
		Code     string `json:"code"`
	} `json:"errors"`
	DocumentationURL string `json:"documentation_url"`
}

func (e GithubError) Error() string {
	return fmt.Sprintf("github: %v %+v %v", e.Message, e.Errors, e.DocumentationURL)
}

func Request(path string,param map[string]interface{}) ([]Issue, *http.Response, error) {

	issues := new([]Issue)
	githubError := new(GithubError)
	resp, err := sling.New().Get(path).Receive(issues, githubError)

	//fmt.Printf("%s", err)
	//fmt.Printf("%s", issues)
	if err == nil {
		err = githubError
	}
	return *issues, resp, err
}