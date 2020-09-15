package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pragmatically-dev/apirest/src/api/models"
	"github.com/pragmatically-dev/apirest/src/api/repository"

	"github.com/gorilla/mux"
	"github.com/pragmatically-dev/apirest/src/api/responses"
)

//GetPosts return all posts from db
func GetPosts(w http.ResponseWriter, r *http.Request) {

	repo, err := repository.GetRepositoryPostCrud(w, r) //obtiene la estructura *RepositoryPostCRUD
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//funcion anonima que se encarga de tomar como parametro a cualquier estructura que implemente la interfaz UserRepository
	func(postRepository repository.PostRepository) {
		posts, err := postRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, posts)
	}(repo) //--> se llama la funcion anonima mediante ()

}

//GetPost return an posts from db
func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	responses.JSON(w, http.StatusOK, "post of "+id)
}

//CreatePost create post in db
func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post) //verifica si puede ser convertido a la struct User
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err) //manejador de errores
		return
	}
	repo, err := repository.GetRepositoryPostCrud(w, r) //obtiene la estructura *RepositoryUsersCRUD
	//funcion anonima que se encarga de tomar como parametro a cualquier estructura que implemente la interfaz UserRepository
	func(postRepository repository.PostRepository) {
		post, err := postRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, post.ID))
		responses.JSON(w, http.StatusOK, post)
	}(repo)
} //--> se callea la funcion }

//UpdatePost update an post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
}

//DeletePost delete an post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)

}
