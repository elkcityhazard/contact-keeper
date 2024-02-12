package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/elkcityhazard/contact-keeper/internal/models"
	"github.com/elkcityhazard/contact-keeper/internal/repository"
	"github.com/elkcityhazard/contact-keeper/pkg/errors"
)

func (m *Repository) HandleAddUser(w http.ResponseWriter, r *http.Request) {

	type tmpUser struct {
		*models.User
		Password1 string `json:"password1"`
		Password2 string `json:"password2"`
	}

	errors := errors.NewErrors()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		return
	}

	defer r.Body.Close()

	var tUser = tmpUser{}

	err = json.Unmarshal(body, &tUser)
	if err != nil {
		log.Println(err)
		return
	}

	// check the errors

	if tUser.FirstName == "" {
		errors.AddError("firstName", "first name is required")
	}

	if tUser.LastName == "" {
		errors.AddError("lastName", "last name is required")
	}

	if tUser.Email == "" {
		errors.AddError("email", "email is required")
	}

	if tUser.Password1 == "" {
		errors.AddError("password1", "password is required")
	}

	if tUser.Password2 == "" {
		errors.AddError("password2", "password is required")
	}

	if !strings.EqualFold(tUser.Password1, tUser.Password2) {

		errors.AddError("password1", "passwords do not match")
		errors.AddError("password2", "passwords do not match")
	}

	_, err = mail.ParseAddress(tUser.Email)

	if err != nil {
		errors.AddError("email", "invalid email")
	}

	if errors.HasErrors() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		errors.AddError("body", "could not unmarshal body")
	}

	user.Password.PlainText = tUser.Password1
	user.Password.GenerateHashedPassword()

	log.Printf("%+v\n", user)

	user.Required("FirstName", "LastName", "Email")

	var dbRepo = repository.NewSqlUserRepository(Repo.app)

	if errors.HasErrors() {
		errors.SendJSON(http.StatusBadRequest, w, "error")
		return
	}

	insertedUser, err := dbRepo.CreateUser(r.Context(), user)
	if err != nil {
		errors.AddError("insertedUser", err.Error())
		errors.SendJSON(http.StatusInternalServerError, w, "error")
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(insertedUser); err != nil {
		errors.AddError("encodeJSON", err.Error())
		errors.SendJSON(http.StatusInternalServerError, w, "error")
		return
	}

}
