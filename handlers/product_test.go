package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	reqBody := `{"id":1,"user_id":1,"product_name":"Test Product","product_description":"Test description","product_price":50.0}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	CreateProduct(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}
