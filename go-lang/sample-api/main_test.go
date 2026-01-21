package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
	}

	ct := rr.Header().Get("Content-Type")
	if ct != "application/json; charset=utf-8" {
		t.Fatalf("content-type = %q, want %q", ct, "application/json; charset=utf-8")
	}

	var resp helloResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if resp.Message != "hello world" {
		t.Fatalf("message = %q, want %q", resp.Message, "hello world")
	}
}

func TestDoubleHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/double/21", nil)
	rr := httptest.NewRecorder()

	doubleHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
	}

	ct := rr.Header().Get("Content-Type")
	if ct != "application/json; charset=utf-8" {
		t.Fatalf("content-type = %q, want %q", ct, "application/json; charset=utf-8")
	}

	var resp doubleResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if resp.Input != 21 || resp.Double != 42 {
		t.Fatalf("resp = %+v, want input=21 double=42", resp)
	}
}

func TestDoubleHandler_InvalidNumber(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/double/not-a-number", nil)
	rr := httptest.NewRecorder()

	doubleHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusBadRequest)
	}
}

func TestDoubleHandler_MissingNumber(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/double/", nil)
	rr := httptest.NewRecorder()

	doubleHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusBadRequest)
	}
}

func TestDoubleHandler_OutOfRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/double/9223372036854775807", nil)
	rr := httptest.NewRecorder()

	doubleHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusBadRequest)
	}
}

func TestDoubleHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/different/1", nil)
	rr := httptest.NewRecorder()

	doubleHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusNotFound)
	}
}
