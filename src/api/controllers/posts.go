package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pragmatically-dev/apirest/src/api/responses"
)

//GetPosts return all users from db
func GetPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	responses.JSON(w, http.StatusOK, "editing posts of user "+id)

	//responses.JSON(w, http.StatusOK, "FUNCIONA EL ROUTER PRROO (AGUANTE LOS PUNTEROS Y LA MARIEL!)")
	/*
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
		}(repo) //--> se llama la funcion anonima mediante ()
	*/
}

//GetPost return an users from db
func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	responses.JSON(w, http.StatusOK, "post of "+id)
}

//CreatePost create user in db
func CreatePost(w http.ResponseWriter, r *http.Request) {
}

//UpdatePost update an user
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
}

//DeletePost delete an user
func DeletePost(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)

}
