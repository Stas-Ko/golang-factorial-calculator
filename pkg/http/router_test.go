package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter()

	testCases := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "Valid Route",
			method:       "POST",
			path:         "/calculate",
			expectedCode: http.StatusOK,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var jsonStr = []byte(`{"a": 3, "b": 5}`)  // Updated test data
			req, err := http.NewRequest(tc.method, tc.path, strings.NewReader(string(jsonStr)))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)
		})
	}
}
