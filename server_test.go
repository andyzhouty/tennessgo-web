package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func checkEquation(got, want interface{}, t *testing.T) {
	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}

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

	t.Run("reserved keyword", func(t *testing.T) {
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
	})
	t.Run("empty string", func(t *testing.T) {
		writer := httptest.NewRecorder()
		jsonData := strings.NewReader(`{"to_translate":""}`)
		request, _ := http.NewRequest("POST", "/api", jsonData)
		var resp map[string]string
		mux.ServeHTTP(writer, request) // 发送请求
		json.Unmarshal(writer.Body.Bytes(), &resp)
		checkEquation(resp["error"], "empty string to translate", t)
	})
}

func TestFrontend(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request) // 发送请求

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
