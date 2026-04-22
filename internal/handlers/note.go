package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/brunabbelini/quicknotes/internal/apperror"
	"github.com/brunabbelini/quicknotes/internal/repositories"
)

type noteHandler struct {
	repo repositories.NoteRepository
}

func NewNoteHandler(repo repositories.NoteRepository) *noteHandler {
	return &noteHandler{repo: repo}
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return ErrNotFound
	}

	files := []string{
		"views/template/base.html",
		"views/template/pages/home.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return ErrInternal
	}
	notes, err := nh.repo.List(r.Context())
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, "base", newNoteResponseFromNoteList(notes))
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		return apperror.WithStatus(errors.New("anotação é obrigatória"), http.StatusBadRequest)
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	files := []string{
		"views/template/base.html",
		"views/template/pages/note-view.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		return ErrInternal
	}
	note, err := nh.repo.GetById(r.Context(), id)
	if err != nil {
		return err
	}
	buff := &bytes.Buffer{}
	err = t.ExecuteTemplate(buff, "base", newNoteResponseFromNote(note))
	if err != nil {
		return ErrInternal
	}
	buff.WriteTo(w)
	return nil
}

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		// rejeitar a requisição
		return apperror.WithStatus(errors.New("operação não permitida"), http.StatusMethodNotAllowed)
	}
	fmt.Fprint(w, "Criando uma nova nota...")
	return nil
}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) error {
	files := []string{
		"views/template/base.html",
		"views/template/pages/note-new.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		return ErrInternal
	}
	return t.ExecuteTemplate(w, "base", nil)
}
