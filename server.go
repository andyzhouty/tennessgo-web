package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/z-t-y/tennessgo"
)

type RequestModel struct {
	ToTranslate string `json:"to_translate"`
}

type ResponseModel struct {
	ToTranslate string `json:"to_translate"`
	Translated  string `json:"translated"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/api", handleAPIRequest)
	server.ListenAndServe()
}

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleAPIGet(w, r)
	case "POST":
		err = handleAPIPost(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleAPIGet(w http.ResponseWriter, r *http.Request) (err error) {
	response := map[string]string{
		"name":        "Tennessine-Go API",
		"description": "处理不规范中文句子的Web API",
	}
	output, err := json.MarshalIndent(&response, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handleAPIPost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var request RequestModel
	json.Unmarshal(body, &request)
	translation := tennessgo.NewTranslation(request.ToTranslate)
	translated, err := translation.Translate()
	if err != nil {
		return err
	}
	response := ResponseModel{
		ToTranslate: request.ToTranslate,
		Translated:  translated,
	}
	output, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return errors.New("error while marshaling json: " + err.Error())
	}
	w.Write(output)
	return
}
