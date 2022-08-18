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

// Model for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Middleware, helper - file
func (course *Course) IsEmpty() bool {
	// return course.CourseId == "" || course.CourseName == ""
	return course.CourseName == ""
}

func main() {
	fmt.Println("API - Courses")

	// Seed data

	courses = append(courses, Course{CourseId: "1", CourseName: "React Js",
		CoursePrice: 1201, Author: &Author{Fullname: "Clint", Website: "c.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Vue Js",
		CoursePrice: 1800, Author: &Author{Fullname: "Clint", Website: "c.com"}})

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - file

// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API service</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course")
	w.Header().Set("Content-Type", "application/json")

	// Get if from request
	params := mux.Vars(r)
	fmt.Println(params)

	// Find the course
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found for the couse id : " + params["id"])
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)

	// Body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please fill required fields")
		return
	}

	for index, _ := range courses {
		if courses[index].CourseId == params["id"] {
			courses[index].Author.Fullname = course.Author.Fullname
			courses[index].Author.Website = course.Author.Website
			courses[index].CourseName = course.CourseName
			courses[index].CoursePrice = course.CoursePrice
			json.NewEncoder(w).Encode("Course updated")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found for the couse id : " + params["id"])
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")

	// Body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please fill required fields")
		return
	}

	for _, eachCourse := range courses {
		if eachCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Duplicate entry")
			return
		}
	}

	// Generate unique id, string
	// Append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(10000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Course Not Found")
	return

}
