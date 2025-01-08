package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult Response
	}{
		{
			name:           "first",
			expression:     `{"expression":"5*7"}`,
			expectedResult: Response{"35"},
		},
		{
			name:           "second",
			expression:     `{"expression":"2+7-5*(3-4)"}`,
			expectedResult: Response{"14"},
		},
		{
			name:           "third",
			expression:     `{"expression":"(2+2)*(3+3)"}`,
			expectedResult: Response{"24"},
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer([]byte(testCase.expression)))
			req.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(calculateHandle)
			handler.ServeHTTP(recorder, req)

			rar := Response{}
			json.Unmarshal(recorder.Body.Bytes(), &rar)
			if recorder.Code != http.StatusOK {
				t.Errorf("Wrong status code: %d, wanted 200", recorder.Code)
			} else if testCase.expectedResult != rar {
				t.Errorf("Wrong data: %s, wanted %s", testCase.expectedResult, rar)
			}
		})
	}
}

func TestHandlerWrong(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult Errorr
	}{
		{
			name:           "four",
			expression:     ``,
			expectedResult: Errorr{"Internal server error"},
		},
		{
			name:           "five",
			expression:     `{"expression":"2+"}`,
			expectedResult: Errorr{"Expression Is not Valid"},
		},
		{
			name:           "six",
			expression:     `[2, 3]`,
			expectedResult: Errorr{"Internal server error"},
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer([]byte(testCase.expression)))
			req.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(calculateHandle)
			handler.ServeHTTP(recorder, req)

			rar := Errorr{}
			json.Unmarshal(recorder.Body.Bytes(), &rar)
			if testCase.expectedResult != rar {
				t.Errorf("Wrong data: got %s, wanted %s", rar, testCase.expectedResult)
			}
		})
	}
}
