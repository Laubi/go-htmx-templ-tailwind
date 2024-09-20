package main

import (
	"cmp"
	"embed"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/Laubi/go-htmx-templ-tailwind/internal/persistence"
	"github.com/Laubi/go-htmx-templ-tailwind/internal/views"
)

//go:embed static
var staticAssets embed.FS

func main() {
	err := persistence.ConnectToDb()
	if err != nil {
		panic(err)
	}
	defer persistence.CloseDb()

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(staticAssets)))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		users, err := persistence.GetAllUsers()
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Internal server error. The error has been logged", http.StatusInternalServerError)
			return
		}

		component := views.Users(users)
		err = component.Render(r.Context(), w)
		if err != nil {
			slog.Error("failed rendering component", "error", err)
		}
	})

	mux.HandleFunc("POST /user/", func(writer http.ResponseWriter, request *http.Request) {
		u := persistence.User{
			Email:    request.FormValue("email"),
			Username: request.FormValue("username"),
		}

		u, err := persistence.CreateUser(u)
		if err != nil {
			slog.Error(err.Error())
			http.Error(writer, "Internal server error. The error has been logged", http.StatusInternalServerError)
			return
		}

		component := views.User(u)
		err = component.Render(request.Context(), writer)
		if err != nil {
			slog.Error("failed rendering component", "error", err)
			http.Error(writer, "Internal server error. The error has been logged", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("DELETE /user/{id}", func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.Atoi(request.PathValue("id"))
		if err != nil {
			slog.Error("invalid id", "error", err, "id", request.PathValue("id"))
			http.Error(writer, "Internal server error. The error has been logged", http.StatusInternalServerError)
			return
		}

		if persistence.DeleteUser(id) != nil {
			slog.Error("failed to delete user", "id", id)
			http.Error(writer, "Internal server error. The error has been logged", http.StatusInternalServerError)
			return
		}

	})

	port := cmp.Or(os.Getenv("PORT"), "8080")
	slog.Info("starting server", "port", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		panic(err)
	}
}
