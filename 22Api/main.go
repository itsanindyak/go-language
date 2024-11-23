package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID   string  `json:"courseid"`
	CourseName string  `json:"name"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses = []Course{
	{CourseID: "3", CourseName: "React JS", Price: 299, Author: &Author{Fullname: "Anindya", Website: "www.anindya.kol"}},
	{CourseID: "4", CourseName: "Angular JS", Price: 399, Author: &Author{Fullname: "Aniket", Website: "www.aniket.kol"}},
}

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/course", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PATCH")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1234", r))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to AKA Course </h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the id")

}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// if no data is send
	if r.Body == nil {
		json.NewEncoder(w).Encode("Not data found.")
	}

	var course Course

	json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("no data is recived")
		return
	}

	for _, eachCourse := range courses {
		if eachCourse.CourseName == course.CourseName {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Course with same name is already present.")
			return
		}
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	var flag bool = false
	var updatedcourse Course
	for index, course := range courses {
		if course.CourseID == id {
			flag = true
			courses = append(courses[:index], courses[index+1:]...) //imp
			err := json.NewDecoder(r.Body).Decode(&updatedcourse)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Invalid request body")
				return
			}
			updatedcourse.CourseID = id
			courses = append(courses, updatedcourse)
			break
		}
	}
	if !flag {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id is not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedcourse)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for index, course := range courses {
		if course.CourseID == id {
			courses = append(courses[:index], courses[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Course delete succesfully.")
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("No course found.")

}
