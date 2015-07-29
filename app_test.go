package main

import (
	"github.com/readit-tw/readit-api/model"
	"github.com/readit-tw/readit-api/repository"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateResourceHandler(t *testing.T) {
	mockRepo := new(repository.MockResourceRepository)

	mockResource := &model.Resource{Id: "", Link: "http://www.google.com"}
	mockRepo.On("Create", mockResource).Return(mockResource, nil)

	ts := httptest.NewServer(http.HandlerFunc(createResourceHandler(mockRepo)))
	defer ts.Close()

	request := "{\"id\":\"\",\"title\":\"\",\"link\":\"http://www.google.com\"}"
	body := strings.NewReader(request)
	res, err := http.Post(ts.URL, "application/json", body)
	assert.Nil(t, err)

	actualBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err)

	assert.Equal(t, string(actualBytes), request)

	assert.Equal(t, res.StatusCode, http.StatusCreated)

	contentType := res.Header.Get("Content-Type")
	assert.Equal(t, "application/json", contentType)
	mockRepo.AssertExpectations(t)
}

func TestListResourcesHandler(t *testing.T) {
	mockRepo := new(repository.MockResourceRepository)
	mockResources := []*model.Resource{&model.Resource{Link: "http://google.com"}, &model.Resource{Link: "http://yahoo.com"}}
	mockRepo.On("GetAll").Return(mockResources, nil)

	ts := httptest.NewServer(http.HandlerFunc(listResourcesHandler(mockRepo)))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.Nil(t, err)

	actualBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err)

	assert.Equal(t, "[{\"id\":\"\",\"title\":\"\",\"link\":\"http://google.com\"},{\"id\":\"\",\"title\":\"\",\"link\":\"http://yahoo.com\"}]", string(actualBytes))

	contentType := res.Header.Get("Content-Type")
	assert.Equal(t, "application/json", contentType)
	assert.Equal(t, res.StatusCode, http.StatusOK)
	mockRepo.AssertExpectations(t)
}
