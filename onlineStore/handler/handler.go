package handler

import (
	"Test_task1/onlineStore/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) MapRoutes() http.Handler {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(h.BasicAuthMiddleware)

	seller := api.PathPrefix("/seller").Subrouter()
	h.setupRoutes(seller, map[string]http.HandlerFunc{
		"create":  h.CreateSeller,
		"getAll":  h.GetAllSellers,
		"getById": h.GetSellerById,
		"update":  h.UpdateSeller,
		"delete":  h.DeleteSeller,
	})

	return r
}

func (h *Handler) setupRoutes(subRouter *mux.Router, handlers map[string]http.HandlerFunc) {
	subRouter.HandleFunc("/", handlers["create"]).Methods("POST")
	subRouter.HandleFunc("/", handlers["getAll"]).Methods("GET")
	subRouter.HandleFunc("/{id}", handlers["getById"]).Methods("GET")
	subRouter.HandleFunc("/{id}", handlers["updateById"]).Methods("PUT")
	subRouter.HandleFunc("/{id}", handlers["deleteById"]).Methods("DELETE")
}
