package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// model for the course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fakeDB

var courses []Course

// middleware or helper - file

func (c *Course) IsEmpty() bool {

	return c.CourseName == ""

}

func main() {
	fmt.Println("Building the rest apis")

}

//controller -file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Welcome to API </h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getcoursebyId(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id for request
	params := mux.vars(r)

	// loop through courses , find matching id and return the resposne

	for _, course := range courses {
		if course.CourseId == params["id"] {

			json.NewEncoder(w).Encode(course)
			return

		}
	}

	json.NewEncoder(w).Encode("No Course found with given id : ")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one courses")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send the some data")
	}

	// what about {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {

		json.NewEncoder(w).Encode("Please send the some data")
		return

	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one courses")
	w.Header().Set("Content-Type", "application/json")

	// id from prequest

	params := mux.vars(r)

	// loop , id , remove index , add with myid
	for index, course := range courses {

		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Value has been updated : ")
			return

		} else {
			json.NewEncoder(w).Encode("course not forund by Id : ")
			return
		}

	}

}

func delete(w http.ResponseWriter, r *http.Request) {

	fmt.Println("delete one courses")
	w.Header().Set("Content-Type", "application/json")

	// id from prequest

	params := mux.vars(r)
	var course Course

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(course)

}
