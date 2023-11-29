// pkg/calculate/handler_test.go

package calculate

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHandler(t *testing.T) {
	testCases := []struct {
		name          string
		inputJSON     string
		expectedCode  int
		expectedError string
	}{
		{
			name:          "Valid Input",
			inputJSON:     `{"a": 3, "b": 5}`,
			expectedCode:  http.StatusOK,
			expectedError: "",
		},
		{
			name:          "Invalid Input (Negative Values)",
			inputJSON:     `{"a": -3, "b": 5}`,
			expectedCode:  http.StatusBadRequest,
			expectedError: "Incorrect input",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/calculate", strings.NewReader(tc.inputJSON))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CalculateHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)

			if tc.expectedError != "" {
				assert.Contains(t, rr.Body.String(), tc.expectedError)
			}
		})
	}
}
