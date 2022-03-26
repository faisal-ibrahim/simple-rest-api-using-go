package http

import (
	"encoding/json"
	"fmt"
	"github.com/github.com/faisal-ibrahim/simple-rest-api-using-go/internal/comment"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/update/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/delete/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive"}); err != nil {
			panic(err)
		}
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse uint from ID")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Error retriving comment by ID!")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	comments, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Fprintf(w, "Faild to retrive all comments!")
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprintf(w, "Faild to post new comment")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// UpdateComment - updates a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

}

// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

}
