package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Article struct {
	Title string "json:title"
	Desc  string "json:desc"
}

type Articles []Article

var articles = Articles{
	Article{Title: "Hello", Desc: "Desc 1"},
	Article{Title: "Hello2", Desc: "Desc 2"},
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/articles-post", withLogging(postArticles))
	http.ListenAndServe(":3000", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func getArticles(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(articles)
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newArticle Article
		err := json.NewDecoder(r.Body).Decode(&newArticle)
		if err != nil {
			http.Error(w, "Can't read body", http.StatusInternalServerError)
		}

		articles = append(articles, newArticle)
		json.NewEncoder(w).Encode(articles)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged koneksi dari", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
