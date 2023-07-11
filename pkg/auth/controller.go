package auth

import (
	"log"
	"net/http"

	"github.com/akifkadioglu/youtube-1/database"
	"github.com/akifkadioglu/youtube-1/ent/user"
	"github.com/akifkadioglu/youtube-1/models"
	"github.com/akifkadioglu/youtube-1/utils"
	"github.com/go-chi/render"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input BodyRegister
	db := database.DBManager()
	err := render.DecodeJSON(r.Body, &input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"message": "All fields required",
		})
		return
	}

	userCount, err := db.User.
		Query().
		Where(user.Username(input.Username)).
		Count(r.Context())

	if err != nil && userCount > 0 {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"message": "This username is already using by someone",
		})
		return
	}
	password := utils.Hash(input.Password)
	log.Println(password)
	_, err = db.User.
		Create().
		SetUsername(input.Username).
		SetPassword(password).
		Save(r.Context())

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"message": "Something went wrong",
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"message": "User created.",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input BodyLogin
	db := database.DBManager()
	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"message": "All fields required",
		})
		return
	}
	user, err := db.User.
		Query().
		Where(user.Username(input.Username)).
		First(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"message": "user not found",
		})
		return
	}
	isCompared := utils.CompareHash(user.Password, input.Password)
	if !isCompared {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"message": "user not found",
		})
		return
	}

	token, err := utils.GenerateToken(models.JwtModel{
		Name: user.Username,
		Age:  21,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"message": "Something went wrong",
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"token": token,
	})
}
