package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/project-detail/{id}", project_detail).Methods("GET")
	route.HandleFunc("/add-project", Add_Project).Methods("POST")
	fmt.Println("Server berjalan pada port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("view/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("view/project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("view/Input_Form.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func project_detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("view/blog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	data := map[string]interface{}{
		"Project_Name": "Belajar Go Languages",
		"Description":  "ini deskripsi (testing)",
		"Id":           id,
	}
	tmpt.Execute(w, data)
}

func Add_Project(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Project Name : " + r.PostForm.Get("project_name"))
	fmt.Println("Start Date : " + r.PostForm.Get("start-date"))
	fmt.Println("End Date : " + r.PostForm.Get("end-date"))
	fmt.Println("Description : " + r.PostForm.Get("description"))
	fmt.Println("Tecnologies : " + r.PostForm.Get("node-js"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
