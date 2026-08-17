package cars

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Car struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Make  string `json:"make"`
	Color string `json:"color"`
	Year  int    `json:"year"`
}

var cars = []Car{
	{ID: 1, Name: "Camry", Make: "Toyota", Color: "Red", Year: 2012},
	{ID: 2, Name: "Jetta", Make: "Jetta", Color: "Blue", Year: 2001},
	{ID: 3, Name: "Spider", Make: "Toyota", Color: "White", Year: 2021},
	{ID: 4, Name: "Pathfinder", Make: "Toyota", Color: "Green", Year: 2023},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(cars)

	case http.MethodPost:
		var car Car
		json.NewDecoder(r.Body).Decode(&car)
		cars = append(cars, car)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(car)

	case http.MethodPut:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id not found", http.StatusBadRequest)
			return
		}

		var updatecar Car
		if err := json.NewDecoder(r.Body).Decode(&updatecar); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		for i, car := range cars {
			if car.ID == id {
				cars[i] = updatecar

				json.NewEncoder(w).Encode(updatecar)
				return
			}
		}
		http.Error(w, "car not found", http.StatusNotFound)

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		for i, car := range cars {
			if car.ID == id {
				cars = append(cars[:i], cars[i+1:]...)
				w.Write([]byte("car deleted"))
				return
			}
		}
		http.Error(w, "car not found", http.StatusNotFound)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
