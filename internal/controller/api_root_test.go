package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"devchallenge.it/conversation/internal/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type TestContext struct {
	service *Service
	router  *mux.Router
	dao     *model.Dao
}

func NewTestContext() *TestContext {
	r := mux.NewRouter()

	dao := model.NewDao(nil)

	service := New(r, dao)

	return &TestContext{
		service: service,
		router:  r,
		dao:     dao,
	}
}

func TestGetRoot(t *testing.T) {
	t.Skip("Not implemented")
	tctx := NewTestContext()

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	tctx.router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}
