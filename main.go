// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Animal struct represents the data structure for an animal.
type Animal struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Legs  int    `json:"legs"`
}

// App struct holds the application state.
type App struct {
	animals map[int]Animal
	mu      sync.RWMutex
}

// NewApp initializes a new App instance.
func NewApp() *App {
	return &App{
		animals: make(map[int]Animal),
	}
}

// CreateAnimal handles the creation of a new animal entry.
func (a *App) CreateAnimal(w http.ResponseWriter, r *http.Request) {
	var animal Animal
	if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	// Check if the ID already exists
	if _, exists := a.animals[animal.ID]; exists {
		http.Error(w, "Duplicate entry", http.StatusBadRequest)
		return
	}

	// Save the animal
	a.animals[animal.ID] = animal
	w.WriteHeader(http.StatusCreated)
}

// UpdateAnimal handles the update of an existing animal or creates a new one.
func (a *App) UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var animal Animal
	if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	// Update the existing animal or create a new one
	a.animals[id] = animal
	w.WriteHeader(http.StatusOK)
}

// DeleteAnimal handles the deletion of an existing animal.
func (a *App) DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	// Check if the ID exists
	if _, exists := a.animals[id]; !exists {
		http.Error(w, "Animal not found", http.StatusNotFound)
		return
	}

	// Delete the animal
	delete(a.animals, id)
	w.WriteHeader(http.StatusOK)
}

// GetAnimals returns a list of all currently existing animals.
func (a *App) GetAnimals(w http.ResponseWriter, r *http.Request) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	// Convert map values to a slice
	animals := make([]Animal, 0, len(a.animals))
	for _, animal := range a.animals {
		animals = append(animals, animal)
	}

	// Check if any animals exist
	if len(animals) == 0 {
		http.Error(w, "No animals found", http.StatusNotFound)
		return
	}

	// Return the list of animals
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animals)
}

// GetAnimalByID returns an animal by its ID.
func (a *App) GetAnimalByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	// Check if the ID exists
	if animal, exists := a.animals[id]; exists {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(animal)
		return
	}

	http.Error(w, "Animal not found", http.StatusNotFound)
}

func main() {
	app := NewApp()

	router := mux.NewRouter()
	router.HandleFunc("/v1/animal", app.CreateAnimal).Methods("POST")
	router.HandleFunc("/v1/animal/{id:[0-9]+}", app.UpdateAnimal).Methods("PUT")
	router.HandleFunc("/v1/animal/{id:[0-9]+}", app.DeleteAnimal).Methods("DELETE")
	router.HandleFunc("/v1/animal", app.GetAnimals).Methods("GET")
	router.HandleFunc("/v1/animal/{id:[0-9]+}", app.GetAnimalByID).Methods("GET")

	fmt.Println("Server listening on :8069")
	http.ListenAndServe(":8069", router)
}
