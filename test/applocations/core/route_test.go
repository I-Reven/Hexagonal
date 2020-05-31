package core

import (
	"encoding/json"
	"fmt"
	"github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	_ = godotenv.Load("../../../.test.env")
}

func TestRoutePing(t *testing.T) {
	router := core.Http{}.Route()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
	assert.Equal(t, "PONG", w.Body.String())
}

func TestRequestIAmAlive(t *testing.T) {
	router := core.Http{}.Route()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/i-am-alive", nil)
	router.ServeHTTP(w, req)

	i := entity.IAmAlive{}
	_ = json.Unmarshal(w.Body.Bytes(), &i)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, i.GetDbSuccess())
	assert.Equal(t, true, i.GetCashSuccess())
	assert.Equal(t, true, i.GetConsumerSuccess())
	assert.Equal(t, true, i.GetGrpcSuccess())
	assert.Equal(t, true, i.GetHttpSuccess())
	assert.Equal(t, "I Am Alive", i.GetContent())
}
