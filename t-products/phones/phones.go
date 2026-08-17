package phones

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Phone struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Model string `json:"model"`
}

var phones = []Phone{
	{ID: 1, Name: "Tecno", Color: "Green", Model: "Spack30C"},
	{ID: 2, Name: "Iphone", Color: "Black", Model: "Iphone 15 Promax"},
	{ID: 3, Name: "Samsung", Color: "Gray", Model: "Galaxy505 pro"},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			json.NewEncoder(w).Encode(phones)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		for _, phone := range phones {
			if phone.ID == id {
				json.NewEncoder(w).Encode(phone)
				return
			}
			http.Error(w, "bad request", http.StatusBadRequest)
		}

	case http.MethodPost:
		var newProduct Phone
		json.NewDecoder(r.Body).Decode(&newProduct)
		phones = append(phones, newProduct)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("created"))
		json.NewEncoder(w).Encode(newProduct)

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		for i, p := range phones {
			if p.ID == id {
				phones = append(phones[:i], phones[i+1:]...)
				w.Write([]byte("deleted"))
				return
			}
		}
		http.Error(w, "phone not found", http.StatusNotFound)

	case http.MethodPatch:
		var updatedPhone Phone

		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&updatedPhone); err != nil {
			http.Error(w, "error Decoding json", http.StatusBadRequest)
			return
		}
		for i, phone := range phones {
			if phone.ID == id {
				phones[i] = updatedPhone

				json.NewEncoder(w).Encode(updatedPhone)
				return
			}
		}
		http.Error(w, "unfound phone", http.StatusNotFound)

	default:
		http.Error(w, "phone not found", http.StatusNotFound)
	}

}
