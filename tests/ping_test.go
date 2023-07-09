package nain_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ma1y0/go-bank/routes"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
    r := routes.CreateRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/ping", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
}
