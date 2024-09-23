package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"devchallenge.it/conversation/internal/model"
	"github.com/go-redis/redismock/v9"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type TestContext struct {
	service *Service
	router  *mux.Router
	dao     *model.Dao
	mock    redismock.ClientMock
}

func NewTestContext() *TestContext {
	r := mux.NewRouter()

	rdb, mock := redismock.NewClientMock()
	dao := model.NewDao(rdb)

	service := New(r, dao)

	return &TestContext{
		service: service,
		router:  r,
		dao:     dao,
		mock:    mock,
	}
}

func TestGetRoot(t *testing.T) {
	t.Skip("Not implemented")
	tctx := NewTestContext()

	/*
		Create expected redis	requests
		tctx.mock.ExpectHGet("devchallenge-xx", "var1").SetVal("1")
	*/

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	tctx.router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}
