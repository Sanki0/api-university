package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"

)

type Student struct {
	ID   string    `json:"id"`
	Name string `json:"name"`
}

func getStudents() []*Student {
	// Open up our database connection.
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM students")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var students []*Student
	for results.Next() {
		var s Student
		// for each row, scan the result into our tag composite object
		err = results.Scan(&s.ID, &s.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		students = append(students, &s)
	}

	return students
}


/*func getSingleStudent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key,ok := vars["id"]
	if !ok {
		fmt.Fprintf(w, "Failed to get ID")
		}		
	fmt.Println("ID: ", key)
}*/


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func studentPage(w http.ResponseWriter, r *http.Request) {
	students := getStudents()

	fmt.Println("Endpoint Hit: studentsPage")
	json.NewEncoder(w).Encode(students)
}

/*func singleStudentPage(w http.ResponseWriter, r *http.Request){
	getSingleStudent(w, r)
}*/


func main() {
	http.HandleFunc("/", homePage)


	//GET STUDENTS
	http.HandleFunc("/students", studentPage)

	//GET SINGLE STUDENT
	//r := mux.NewRouter()
	//r.HandleFunc("/student/{id}", singleStudentPage)

	
	log.Fatal(http.ListenAndServe(":8080", nil))
}