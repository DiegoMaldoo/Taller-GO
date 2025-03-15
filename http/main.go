package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"name"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(usuarios)

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error al leer el body", http.StatusBadRequest)
			return
		}
		var user Usuario
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "Error parseando el JSON", http.StatusBadRequest)
			return
		}
		user.ID = len(usuarios) + 1
		usuarios = append(usuarios, user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	case http.MethodPut:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID invalido", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error al leer el body", http.StatusBadRequest)
			return
		}
		var updatedUser Usuario
		err = json.Unmarshal(body, &updatedUser)
		if err != nil {
			http.Error(w, "error parseando el JSON", http.StatusBadRequest)
			return
		}
		for i, user := range usuarios {
			if user.ID == id {
				usuarios[i].Nombre = updatedUser.Nombre
				usuarios[i].Email = updatedUser.Email
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(usuarios[i])
				return
			}
		}
		http.Error(w, "usuario no encontrado", http.StatusNotFound)

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID invalido", http.StatusBadRequest)
			return
		}
		for i, user := range usuarios {
			if user.ID == id {
				usuarios = append(usuarios[:i], usuarios[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "usuario no encontrado", http.StatusNotFound)

	default:
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "pong")
	} else {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("./public/index.html")
	if err != nil {
		http.Error(w, "error leyendo el html", http.StatusInternalServerError)
		return
	}
	w.Write(content)
}

func main() {
	usuarios = append(usuarios, Usuario{ID: 1, Nombre: "Diego", Email: "diegmaldo@mail.com"})

	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/v1/users", Users)
	http.HandleFunc("/", Index)

	fmt.Println("servidor escuchando en el puerto 3000")
	http.ListenAndServe(":3000", nil)
}
