package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id`
	Name string `json:"name,omitempty"`
}

var users = []User{{1, "Alice"}, {2, "Bob"}}

func main() {
	http.HandleFunc("/users", auth(logger(handleUsers)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("x-id")
		if id == "" {
			log.Printf("| %s | %s | no ID", r.Method, r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", id)
		r = r.WithContext(ctx)
		next(w, r)
	}

}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxID := r.Context().Value("id")
		id, ok := ctxID.(string)
		if !ok {
			log.Printf("%s | %s | wrong id", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("| %s | %s | %s\n", r.Method, r.URL, id)
		next(w, r)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User
	if err = json.Unmarshal(req, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users = append(users, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsers(w, r)
	case http.MethodPost:
		AddUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
