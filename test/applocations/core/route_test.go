package core

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	_ = godotenv.Load("../../../.test.env")
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
