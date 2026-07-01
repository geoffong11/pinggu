package sender

import (
	"bytes"
	"io"
	"math/rand/v2"
	"net/http"
)

type NormalUser struct {
	GetEndpoints  []Get
	PostEndpoints []Post
}

func (normalUser NormalUser) Action() ([]byte, error) {
	var randType int
	if len(normalUser.PostEndpoints) == 0 {
		randType = 0
	} else if len(normalUser.GetEndpoints) == 0 {
		randType = 1
	} else {
		randType = rand.IntN(2)
	}
	var resp *http.Response
	var err error
	if randType == 0 {
		randomGet := rand.IntN(len(normalUser.GetEndpoints))
		GetRequest := normalUser.GetEndpoints[randomGet]
		resp, err = http.Get(GetRequest.Endpoint)
	} else {
		randomPost := rand.IntN(len(normalUser.PostEndpoints))
		PostRequest := normalUser.PostEndpoints[randomPost]
		resp, err = http.Post(PostRequest.Endpoint, "application/json", bytes.NewBuffer(PostRequest.Body))
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
