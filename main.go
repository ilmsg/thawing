package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilmsg/thawing/handler"
)

func main() {

	indexHandler := handler.NewHomeHandler()
	pageHandler := handler.NewPageHandler()
	postHandler := handler.NewPostHandler()
	userHandler := handler.NewUserHandler()
	userPageHandler := handler.NewUserPageHandler()
	userPostHandler := handler.NewUserPostHandler()

	// conn := database.NewDatabase()
	// repoUserContact := repo.NewRepoUserContact(conn)
	// srvUserContact := service.NewServiceUserContact(repoUserContact)

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	r.Handle("/public/", http.StripPrefix("/public/", fs))

	r.HandleFunc("/{:year}", indexHandler.GetListByYear).Methods(http.MethodGet)
	r.HandleFunc("/{:year}/{:month}", indexHandler.GetListByYear).Methods(http.MethodGet)
	r.HandleFunc("/{:year}/{:month}/{:day}", indexHandler.GetListByYear).Methods(http.MethodGet)

	rPage := r.PathPrefix("/pages").Subrouter()
	rPage.HandleFunc("/", pageHandler.List).Methods(http.MethodGet)
	rPage.HandleFunc("/", postHandler.List).Methods(http.MethodGet)
	rPage.HandleFunc("/{:id}", pageHandler.Detail).Methods(http.MethodGet)
	rPage.HandleFunc("/{:id}", postHandler.Detail).Methods(http.MethodGet)

	rUser := r.PathPrefix("/users").Subrouter()
	rUser.HandleFunc("/register", userHandler.GetRegister).Methods(http.MethodGet)
	rUser.HandleFunc("/register", userHandler.PostRegister).Methods(http.MethodPost)
	rUser.HandleFunc("/login", userHandler.GetLogin).Methods(http.MethodGet)
	rUser.HandleFunc("/login", userHandler.PostLogin).Methods(http.MethodPost)
	rUser.HandleFunc("/profile", userHandler.GetProfile).Methods(http.MethodGet)
	rUser.HandleFunc("/profile", userHandler.PostProfile).Methods(http.MethodPost)
	rUser.HandleFunc("/logout", userHandler.GetLogout).Methods(http.MethodGet)

	rUserPage := rUser.PathPrefix("/pages").Subrouter()
	rUserPage.HandleFunc("/", userPageHandler.List).Methods(http.MethodGet)
	rUserPage.HandleFunc("/", userPageHandler.Save).Methods(http.MethodPost)
	rUserPage.HandleFunc("/{:id}", userPageHandler.Detail).Methods(http.MethodGet)
	rUserPage.HandleFunc("/{:id}/edit", userPageHandler.Edit).Methods(http.MethodGet)
	rUserPage.HandleFunc("/{:id}/delete", userPageHandler.Delete).Methods(http.MethodGet)
	rUserPage.HandleFunc("/create", userPageHandler.Create).Methods(http.MethodGet)

	rUserPost := rUser.PathPrefix("/posts").Subrouter()
	rUserPost.HandleFunc("/", userPostHandler.List).Methods(http.MethodGet)
	rUserPost.HandleFunc("/", userPostHandler.Save).Methods(http.MethodPost)
	rUserPost.HandleFunc("/{:id}", userPostHandler.Detail).Methods(http.MethodGet)
	rUserPost.HandleFunc("/{:id}/edit", userPostHandler.Edit).Methods(http.MethodGet)
	rUserPost.HandleFunc("/{:id}/delete", userPostHandler.Delete).Methods(http.MethodGet)
	rUserPost.HandleFunc("/create", userPostHandler.Create).Methods(http.MethodGet)

	r.HandleFunc("/contact", indexHandler.GetAbout).Methods(http.MethodGet)
	r.HandleFunc("/contact", indexHandler.PostContact).Methods(http.MethodPost)
	r.HandleFunc("/about", indexHandler.GetAbout).Methods(http.MethodGet)
	r.HandleFunc("/", indexHandler.GetIndex).Methods(http.MethodGet)

	if err := http.ListenAndServe(":3070", r); err != nil {
		log.Fatal(err)
	}
}
