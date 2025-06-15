package clients

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
)

type HttpImpl struct {
	baseURL          string
	client           *http.Client
	getMethodTimeout *time.Duration
}

func NewHttpClient(baseUrl string, timeout *time.Duration) interfaces.HttpClientInterface {
	cli := &http.Client{}
	return HttpImpl{
		baseURL:          baseUrl,
		client:           cli,
		getMethodTimeout: timeout,
	}
}

func (i HttpImpl) Get(ctx context.Context, path string) (*http.Response, error) {
	if i.getMethodTimeout != nil {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, *i.getMethodTimeout)
		ctx = ctxWithTimeout
		defer cancel()
	}
	fullPath := i.baseURL + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, context.DeadlineExceeded
		}
		return nil, err
	}
	if resp.StatusCode == 408 {
		fmt.Println("erro: request finished with timeout")
	}
	return resp, nil
}
