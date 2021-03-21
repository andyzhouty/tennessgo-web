package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
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
		mux.ServeHTTP(writer, request) // 发送请求
		checkEquation(writer.Code, 400, t)
	})
}

func TestMain(t *testing.T) {
	go main()                    // 在新goroutine中启动服务器
	time.Sleep(time.Millisecond) // 等待服务器启动
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	resp, err := http.Get("http://localhost:" + port + "/api")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
		return
	}

	var dataMap map[string]string
	json.Unmarshal(bytes, &dataMap)
	if dataMap["name"] != "Tennessine-Go API" {
		t.Error("incorrect data received")
	}
}
