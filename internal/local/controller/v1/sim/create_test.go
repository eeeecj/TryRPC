package sim

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestSimulation_Create(t *testing.T) {
	r := setupRouter()
	s := &Simulation{}
	r.POST("/create", s.Create)
	w := httptest.NewRecorder()
	data := map[string]string{
		"data": "xiaoming",
	}
	b, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(b))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response = make(map[string]string)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "", response["message"])
}
