package main

import (
	"github.com/readit-tw/readit-api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockResourceRepository struct {
	mock.Mock
}

func (m *MockResourceRepository) Create(r *model.Resource) (*model.Resource, error) {
	args := m.Called(r)
	return args.Get(0).(*model.Resource), args.Error(1)
}

func TestCreateResourceHandler(t *testing.T) {
	mockRepo := new(MockResourceRepository)

	mockResource := &model.Resource{Id: "", Link: "http://www.google.com"}
	mockRepo.On("Create", mockResource).Return(mockResource, nil)

	ts := httptest.NewServer(http.HandlerFunc(createResourceHandler(mockRepo)))
	defer ts.Close()

	request := "{\"id\":\"\",\"link\":\"http://www.google.com\"}"
	body := strings.NewReader(request)
	res, err := http.Post(ts.URL, "application/json", body)
	assert.Nil(t, err)

	actualBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err)

	assert.Equal(t, string(actualBytes), request)

	assert.Equal(t, res.StatusCode, http.StatusCreated)

	contentType := res.Header.Get("Content-Type")
	assert.Equal(t, contentType, "application/json")
	mockRepo.AssertExpectations(t)
}
