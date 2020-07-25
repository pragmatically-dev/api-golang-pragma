package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"github.com/pragmatically-dev/apirest/src/api/models"
	"github.com/pragmatically-dev/apirest/src/api/repository"
	"github.com/pragmatically-dev/apirest/src/api/responses"
)

//GetUsers return all users from db
func GetUsers(w http.ResponseWriter, r *http.Request) {
	repo, err := repository.GetRepositoryCrud(w, r) //obtiene la estructura *RepositoryUsersCRUD
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//funcion anonima que se encarga de tomar como parametro a cualquier estructura que implemente la interfaz UserRepository
	func(userRepository repository.UserRepository) {
		users, err := userRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, users)
	}(repo) //--> se callea la funcion anonima mediante ()
}

//GetUser return an users from db
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo, err := repository.GetRepositoryCrud(w, r) //obtiene la estructura *RepositoryUsersCRUD
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//funcion anonima que se encarga de tomar como parametro a cualquier estructura que implemente la interfaz UserRepository
	func(userRepository repository.UserRepository) {
		user, err := userRepository.FindByID(_ID)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, user)
	}(repo) //--> se callea la funcion anonima mediante ()
}

//CreateUser create user in db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user) //verifica si puede ser convertido a json
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err) //manejador de errores
		return
	}
	repo, err := repository.GetRepositoryCrud(w, r) //obtiene la estructura *RepositoryUsersCRUD
	//funcion anonima que se encarga de tomar como parametro a cualquier estructura que implemente la interfaz UserRepository
	func(userRepository repository.UserRepository) {
		user, err := userRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusOK, user)
	}(repo) //--> se callea la funcion anonima
}

//UpdateUser update an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//TODO crear controlador UpdateUser
	w.Write([]byte("actualiza de usuarios"))
}

//DeleteUser delete an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO: Crear controlador de DeleteUser
	w.Write([]byte("elimina usuarios"))
}
