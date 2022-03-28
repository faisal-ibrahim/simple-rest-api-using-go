//go:build e2e

package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"id": 12345, "slug": "/", "author": "123455", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")

	if err != nil {
		t.Fail()
	}

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestGetComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment/12345")
	if err != nil {
		t.Fatal()
	}

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUpdateComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "123455", "body": "hello world update"}`).
		Put(BASE_URL + "/api/comment/12345")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestDeleteComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Delete(BASE_URL + "/api/comment/12345")
	if err != nil {
		t.Fatal()
	}

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestGetAllComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")

	if err != nil {
		t.Fatal()
	}

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
