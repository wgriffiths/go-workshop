package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"
)

func TestOClock(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()
	OClock(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200 Go %d", w.Code)
	}

}

func TestOClock_Body(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()
	OClock(w, req)
	value := time.Now().Format("2006-01-02T15:04")
	expected := fmt.Sprintf(`{"time":"%s"}`, value)
	result := w.Body.String()
	if result != expected {
		t.Errorf("Expected %q Result %q", expected, result)
	}

}
