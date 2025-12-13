package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListGrades(t *testing.T) {
	rec := httptest.NewRecorder()
	ListGrades(rec, httptest.NewRequest("GET", "/api/grades", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
