package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anthonydevelops/team-gen/api"
	"github.com/magiconair/properties/assert"
)

var a api.Server

func TestListUsers(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/api/users", nil)
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	assert.Equal(t, response.Code, http.StatusOK)
}
