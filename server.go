package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/rs/cors"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/api", handleAPIRequest)
	corsHandler := cors.Default().Handler(mux)
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, corsHandler)
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/index.html", "templates/layout.html"}
	t := template.Must(template.ParseFiles(files...))
	switch r.Method {
	case "GET":
		t.ExecuteTemplate(w, "layout", nil)
	case "POST":
		tr := tennessgo.NewTranslation(r.FormValue("text"))
		result, err := tr.Translate()
		t.ExecuteTemplate(w, "layout", map[string]string{
			"err":  err.Error(),
			"data": result,
		})
	}
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
		output, _ := json.MarshalIndent(&map[string]string{
			"translated": "",
			"error":      err.Error(),
		}, "", "\t")
		w.Write(output)
		return
	}
}

func handleAPIGet(w http.ResponseWriter, r *http.Request) (err error) {
	response := map[string]string{
		"name":        "Tennessine-Go API",
		"description": "处理不规范中文句子的Web API",
	}
	output, _ := json.MarshalIndent(&response, "", "\t")
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
	output, _ := json.MarshalIndent(response, "", "\t")
	w.Write(output)
	return
}
