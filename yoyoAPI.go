package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Desc    string  `json:"desc"`
	Content string  `json:"content"`
	Author  *Author `json:"author"`
}
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
}

var Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Articles)
	fmt.Println("get all articles hit")
	// json.NewEncoder(w).Encode(articles)

}

func postArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(rand.Intn(1000))
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	fmt.Fprintf(w, "post method hit")
}
func getSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) //parameters
	for _, item := range Articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
	fmt.Fprintf(w, "getsingle method hit")
}
func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range Articles {
		if item.ID == params["id"] {
			updateID := item.ID
			Articles = append(Articles[:index], Articles[index+1:]...)
			var article Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.ID = updateID
			Articles = append(Articles, article)
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	json.NewEncoder(w).Encode(Articles)
	fmt.Fprintf(w, "update method hit")
}
func deletSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range Articles {
		if item.ID == params["id"] {

			Articles = append(Articles[:index], Articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Articles)
	fmt.Fprintf(w, "delete method hit")
}
func handlerRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", getSingleArticle).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", deletSingleArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":881", myRouter))
}
func main() {

	Articles = append(Articles, Article{ID: "1", Title: "random title 1", Desc: "testing desc", Content: "testing content", Author: &Author{FirstName: "di", LastName: "zhao"}})
	Articles = append(Articles, Article{ID: "2", Title: "random title 2", Desc: "testing desc", Content: "testing content", Author: &Author{FirstName: "linlin", LastName: "linlin"}})
	Articles = append(Articles, Article{ID: "3", Title: "random title 3", Desc: "testing desc", Content: "testing content", Author: &Author{FirstName: "gloria", LastName: "gloria"}})
	handlerRequests()
	// router := mux.NewRouter().StrictSlash(true)
	// setupRouter(router)
	// log.Fatal(http.ListenAndServe(":8080", router))
}
