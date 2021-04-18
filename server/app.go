package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static
var static embed.FS

type student struct {
	ID   string `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type App struct {
	db *gorm.DB
	r  *mux.Router
}

func (a *App) start() {
	a.db.AutoMigrate(&student{})
	a.r.HandleFunc("/students", a.getAllStudents).Methods("GET")
	a.r.HandleFunc("/students", a.addStudent).Methods("POST")
	a.r.HandleFunc("/students/{id}", a.updateStudent).Methods("PUT")
	a.r.HandleFunc("/students/{id}", a.deleteStudent).Methods("DELETE")
	webapp, err := fs.Sub(static, "static")
	if err != nil {
		fmt.Println(err)
	}
	// We need to use Gorilla Mux' PathPrefix function here, because the Pathprefix
	// adds a wildcard to the route eg: /*, otherwise we would only route to "/"
	// Hence the error with 404-returning JS files before got thrown, because
	// Gorilla Mux had no route to these JS files.
	a.r.PathPrefix("/").Handler(http.FileServer(http.FS(webapp)))
	log.Fatal(http.ListenAndServe(":8080", a.r))
}

func (a *App) getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var all []student
	err := a.db.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	s.ID = uuid.New().String()
	err = a.db.Save(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (a *App) updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	s.ID = mux.Vars(r)["id"]
	err = a.db.Save(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := a.db.Unscoped().Delete(student{ID: mux.Vars(r)["id"]}).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
