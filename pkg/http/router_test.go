// pkg/http/router_test.go

package http

import (
	"net/http"
	"net/http/httptest"
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
			req, err := http.NewRequest(tc.method, tc.path, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)
		})
	}
}
