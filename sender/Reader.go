package sender

import (
	"io"
	"math/rand/v2"
	"net/http"
)

type Reader struct {
	Endpoints []Get
}

func (reader Reader) Action() ([]byte, error) {
	randomPos := rand.IntN(len(reader.Endpoints))
	getEndpoint := reader.Endpoints[randomPos]
	resp, err := http.Get(getEndpoint.Endpoint)
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
