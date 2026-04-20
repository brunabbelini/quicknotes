package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/brunabbelini/quicknotes/internal/apperror"
)

var ErrNotFound = apperror.WithStatus(errors.New("not found"), http.StatusNotFound)
var ErrInternal = apperror.WithStatus(errors.New("aconteceu um erro ao executar essa operação"), http.StatusInternalServerError)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func (f HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var statusErr apperror.StatusError
		if errors.As(err, &statusErr) {
			if statusErr.StatusCode() == http.StatusNotFound {
				files := []string{
					"views/template/base.html",
					"views/template/pages/404.html",
				}
				t, err := template.ParseFiles(files...)
				if err != nil {
					http.Error(w, err.Error(), statusErr.StatusCode())
				}
				t.ExecuteTemplate(w, "base", statusErr.Error())
				return
			}
			http.Error(w, err.Error(), statusErr.StatusCode())
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
