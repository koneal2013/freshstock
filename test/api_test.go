package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/koneal2013/freshstock/internal/api"
	"github.com/koneal2013/freshstock/internal/model"
	"github.com/koneal2013/freshstock/internal/store"
)

func TestHandlers(t *testing.T) {
	testStore := store.NewProduceStore()
	testProduce := `{"code":"A12T-4GH7-QPL9-3N4M","name":"Lettuce","unit_price":3.46}`
	expectedResult := `{"code":"A12T-4GH7-QPL9-3N4M","name":"Lettuce","unit_price":"$3.46"}`
	expectedResultSearch := `[{"code":"A12T-4GH7-QPL9-3N4M","name":"Lettuce","unit_price":"$3.46"}]`

	err := model.RegisterValidator()
	require.NoError(t, err)

	h := api.NewHandlers(testStore)
	r := api.SetupRoutes(h)

	tests := []struct {
		name   string
		path   string
		body   []byte
		method string
	}{
		{name: "add produce", path: "/api/v1/produce/", body: []byte(testProduce), method: http.MethodPost},
		{name: "get produce by code", path: "/api/v1/produce/A12T-4GH7-QPL9-3N4M", method: http.MethodGet},
		{name: "search produce", path: "/api/v1/produce/?q=let", method: http.MethodGet},
		{name: "delete produce", path: "/api/v1/produce/A12T-4GH7-QPL9-3N4M", method: http.MethodDelete},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, test.path, bytes.NewReader(test.body))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			switch test.name {
			case "add produce":
				assert.Equal(t, http.StatusCreated, w.Code)
				assert.Equal(t, expectedResult, string(w.Body.Bytes()))
			case "delete produce":
				assert.Equal(t, http.StatusNoContent, w.Code)
			case "search produce":
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, expectedResultSearch, string(w.Body.Bytes()))
			default:
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, expectedResult, string(w.Body.Bytes()))
			}
		})
	}
}
