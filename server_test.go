package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPIGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", handleAPIRequest)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api", nil)
	mux.ServeHTTP(writer, request) // 发送请求

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	data := make(map[string]string)
	json.Unmarshal(writer.Body.Bytes(), &data)
	if data["name"] != "Tennessine-Go API" {
		t.Errorf("Cannot retrieve JSON")
	}
}

func TestAPIPost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", handleAPIRequest)

	writer := httptest.NewRecorder()
	jsonData := strings.NewReader(`{"to_translate":"bilibili"}`)
	request, _ := http.NewRequest("POST", "/api", jsonData)
	mux.ServeHTTP(writer, request) // 发送请求

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	data := ResponseModel{}
	json.Unmarshal(writer.Body.Bytes(), &data)
	if data.ToTranslate != "bilibili" || data.Translated != "bilibili" {
		t.Errorf("error occured while parsing JSON")
	}
}
