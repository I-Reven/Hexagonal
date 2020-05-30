package core

import (
	"fmt"
	"github.com/I-Reven/Hexagonal/src/applications/core"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutePing(t *testing.T) {
	router := core.Http{}.Route()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
	assert.Equal(t, "PONG", w.Body.String())
}
