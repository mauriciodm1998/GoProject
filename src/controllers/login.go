package controllers

import (
	"API/src/auth"
	"API/src/bank"
	"API/src/models"
	"API/src/repositorys"
	"API/src/responses"
	"API/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyResquest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyResquest, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := bank.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositorys.NewUserRepository(db)
	userBase, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(userBase.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateToken(userBase.ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(token)

}
