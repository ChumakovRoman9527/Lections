package link

import (
	"11-JWTAUTH/pkg/middleware"
	"11-JWTAUTH/pkg/req"
	"11-JWTAUTH/pkg/res"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type linkHandler struct {
	LinkRepository *LinkRepository
}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &linkHandler{LinkRepository: deps.LinkRepository}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update()))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())

}

func (handler *linkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		link := NewLink(body.URL)
		for {

			existedlink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedlink == nil {
				break
			}
			link.GeneratedHash()
		}
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLink, http.StatusCreated)

	}
}

func (handler *linkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.URL,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, link, 201)
	}
}

func (handler *linkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepository.GetByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.LinkRepository.Delete(link.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, 200)
	}
}

func (handler *linkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
