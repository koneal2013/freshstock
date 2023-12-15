package test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	testProduce := model.Produce{
		Code:      "A12T-4GH7-QPL9-3N4M",
		Name:      "Lettuce",
		UnitPrice: 3.46,
	}
	testProduceBytes, err := json.Marshal(testProduce)
	require.NoError(t, err)

	err = model.RegisterValidator()
	require.NoError(t, err)

	h := api.NewHandlers(testStore)
	r := api.SetupRoutes(h)

	tests := []struct {
		name   string
		path   string
		body   []byte
		method string
	}{
		{name: "add produce", path: "/produce/", body: testProduceBytes,
			method: http.MethodPost},
		{name: "get produce by code", path: fmt.Sprintf("/produce/%s", testProduce.Code),
			method: http.MethodGet},
		{name: "search produce", path: fmt.Sprintf("/produce/?q=%s", testProduce.Name),
			method: http.MethodGet},
		{name: "delete produce", path: fmt.Sprintf("/produce/%s", testProduce.Code),
			method: http.MethodDelete},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, test.path, bytes.NewReader(test.body))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			switch test.name {
			case "add produce":
				assert.Equal(t, http.StatusCreated, w.Code)
				compareResults(t, testProduce, w)
			case "delete produce":
				assert.Equal(t, http.StatusNoContent, w.Code)
			default:
				assert.Equal(t, http.StatusOK, w.Code)
				compareResults(t, testProduce, w)
			}
		})
	}
}

func compareResults(t *testing.T, expected model.Produce, rr *httptest.ResponseRecorder) {
	t.Helper()

	var result model.Produce
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		var results []model.Produce
		err := json.Unmarshal(rr.Body.Bytes(), &results)
		require.NoError(t, err)
		assert.Equal(t, 1, len(results))
		assert.Equal(t, expected.Name, results[0].Name)
		assert.Equal(t, expected.Code, results[0].Code)
		assert.Equal(t, expected.UnitPrice, results[0].UnitPrice)
	} else {
		assert.Equal(t, expected.Name, result.Name)
		assert.Equal(t, expected.Code, result.Code)
		assert.Equal(t, expected.UnitPrice, result.UnitPrice)
	}
}
