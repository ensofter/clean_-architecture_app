package book

import (
	"clean_architecture_app/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//books := h.bookService.GetAllBooks(context.Background(), 10, 10)
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}
