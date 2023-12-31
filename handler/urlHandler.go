package handler

import (
	"log"
	"net/http"
	"url-shortener/helper"
	"url-shortener/model"
	"url-shortener/repository"
	"url-shortener/views"

	"github.com/gorilla/mux"
)

type UrlHandler struct {
	urlRepository *repository.UrlRepository
	views         *views.View
}

func NewUrlHandler(urlRepository *repository.UrlRepository, views *views.View) *UrlHandler {
	return &UrlHandler{
		urlRepository: urlRepository,
		views:         views,
	}
}

func (h *UrlHandler) Add(w http.ResponseWriter, r *http.Request) {
	var addRequest model.AddRequest
	if err := helper.ReadRequestBody(r, &addRequest); err != nil {
		helper.WriteResponse(w, http.StatusBadRequest, "BAD REQUEST", nil)
		return
	}

	if err := helper.Validate(w, addRequest); err != nil {
		return
	}

	url := model.Url{
		ID:  helper.RandStringBytes(16),
		Url: addRequest.Url,
	}

	if err := h.urlRepository.Save(r.Context(), url); err != nil {
		helper.WriteResponse(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "something went wrong")
		log.Println(err)
		return
	}

	helper.WriteResponse(w, http.StatusOK, "OK", model.AddResponse{ShortUrl: "localhost:8080/" + url.ID})
}

func (h *UrlHandler) Visit(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	id := parameter["url"]

	url, err := h.urlRepository.FindById(r.Context(), id)
	if err != nil {
		h.views.Write(w, "notfound", nil)
		return
	}

	http.Redirect(w, r, "https://"+url.Url, http.StatusSeeOther)
}
