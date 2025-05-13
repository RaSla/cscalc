package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "text/plain"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}
}

func TestAddHandler(t *testing.T) {
	tests := []struct {
		a, b     int
		expected string
		status   int
	}{
		{1, 2, "a = 1, b = 2\na + b = 3\n", http.StatusOK},
		{0, 0, "a = 0, b = 0\na + b = 0\n", http.StatusOK},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/plus?a="+url.QueryEscape(strconv.Itoa(test.a))+"&b="+url.QueryEscape(strconv.Itoa(test.b)), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(addHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.status {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.status)
		}

		if rr.Body.String() != test.expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), test.expected)
		}
	}
}

func TestAddHandlerInvalidParams(t *testing.T) {
	tests := []struct {
		url      string
		expected string
		status   int
	}{
		{"/plus?a=abc&b=2", "Invalid parameter 'a'\n", http.StatusBadRequest},
		{"/plus?a=1&b=xyz", "Invalid parameter 'b'\n", http.StatusBadRequest},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(addHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.status {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.status)
		}

		if rr.Body.String() != test.expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), test.expected)
		}
	}
}

func TestAddApiHandler(t *testing.T) {
	tests := []struct {
		a, b     int
		expected string
		status   int
	}{
		{1, 2, `{"result":3}`, http.StatusOK},
		{0, 0, `{"result":0}`, http.StatusOK},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/api/plus?a="+url.QueryEscape(strconv.Itoa(test.a))+"&b="+url.QueryEscape(strconv.Itoa(test.b)), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(addApiHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.status {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.status)
		}

		if rr.Body.String() != test.expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), test.expected)
		}

		expectedContentType := "application/json"
		if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
			t.Errorf("handler returned wrong content type: got %v want %v",
				contentType, expectedContentType)
		}
	}
}
