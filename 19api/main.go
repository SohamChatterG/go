package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course
type Course struct {
	CourseId    string  `json:"couseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {

	r := mux.NewRouter()
	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "reactjs", CoursePrice: 233, Author: &Author{Fullname: "soham", Website: "soham-tech-portfolio.net"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>your 2nd go end point soham</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-type", "application/json")

	// get id frp, request
	params := mux.Vars(r)

	// loop through
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var newCourse Course
	json.NewDecoder(r.Body).Decode(&newCourse)
	if newCourse.IsEmpty() {
		json.NewEncoder(w).Encode("Invalid Input")
		return
	}

	rand.Seed(time.Now().UnixNano())
	newCourse.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, newCourse)
	json.NewEncoder(w).Encode(newCourse)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			course.CourseId = params["id"]
			// make other necessary changes as u wish
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	json.NewEncoder(w).Encode("not found courseId")

}
