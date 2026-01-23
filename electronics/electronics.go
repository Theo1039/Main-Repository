package electronics

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Electronics struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Color string  `json:"color"`
}

var elects = []Electronics{
	{ID: 1, Name: "Television", Price: 4566.00, Color: "black"},
	{ID: 2, Name: "Radio", Price: 1233.00, Color: "pink"},
	{ID: 3, Name: "Fridge", Price: 8900.00, Color: "white"},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			json.NewEncoder(w).Encode(elects)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		for _, elect := range elects {
			if elect.ID == id {
				json.NewEncoder(w).Encode(elect)
			}
		}
		http.Error(w, "id not found", http.StatusNotFound)

	case http.MethodPost:
		var createdProduct Electronics
		if err := json.NewDecoder(r.Body).Decode(&createdProduct); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		elects = append(elects, createdProduct)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("created"))

	case http.MethodPut:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		var updatedProduct Electronics
		if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
			http.Error(w, "error, bad request", http.StatusBadRequest)
			return
		}
		for i, p := range elects {
			if p.ID == id {
				elects[i] = updatedProduct
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("updated"))
				json.NewEncoder(w).Encode(updatedProduct)
			}
		}
		http.Error(w, "request not found", http.StatusNotFound)

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		for i, elect := range elects {
			if elect.ID == id {
				elects = append(elects[:i], elects[i+1:]...)
				w.Write([]byte("deleted"))
				json.NewEncoder(w).Encode(elects)
			}
		}
		http.Error(w, "bad Request", http.StatusBadRequest)

	default:
		http.Error(w, "invalid request", http.StatusNotFound)
	}
}
