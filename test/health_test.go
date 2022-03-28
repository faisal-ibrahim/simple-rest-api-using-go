//go:build e2e

package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/health")

	if err != nil {
		t.Fail()
	}

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
