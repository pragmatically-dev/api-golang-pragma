package routes

import (
	"net/http"

	"github.com/pragmatically-dev/apirest/src/api/controllers"
)

var postsRoutes = []Route{

	Route{
		URI:     "/posts",
		Method:  http.MethodGet,
		Handler: controllers.GetPosts,
	},

	Route{
		URI:     "/posts/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetPost,
	},

	Route{
		URI:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
	},

	Route{
		URI:     "/posts/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdatePost,
	},

	Route{
		URI:     "/posts/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeletePost,
	},
}
