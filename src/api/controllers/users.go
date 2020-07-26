package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
	err = json.Unmarshal(body, &user) //verifica si puede ser convertido a la struct User
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
	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body) //se encarga de obtener todo el contenido del body
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Newuser := models.User{}
	err = json.Unmarshal(body, &Newuser) //verifica si puede ser convertido a un <<User>> y si es asi, lo convierte
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err) //manejador de errores
		return
	}
	_ID, err := primitive.ObjectIDFromHex(vars["id"]) //obtiene el id para realizar una actualizacion de un registro x
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
		userID, err := userRepository.Update(_ID, Newuser)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, fmt.Sprintf("Updated document ID: [ %s ]  at %s ", userID.String(), time.Now()))
	}(repo) //--> se callea la funcion anonima mediante ()
}

//DeleteUser delete an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO: Crear controlador de DeleteUser
	w.Write([]byte("elimina usuarios"))
}
