package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"getEmployees",
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"addEmployee",
		"PUT",
		"employee/add",
		addEmployee,
	},
	Route{
		"deleteteEmployee",
		"DELETE",
		"employee/delete",
		deleteEmployee,
	},
}

type Employee struct {
	Id        string `json:"Id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees []Employee

func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Qux"},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		log.Printf("error occurred while decoding employee data :: ", err)
		return
	}

	log.Printf("adding employee id :: %s with FirstName :: %s and LastName :: %s", employee.Id, employee.FirstName, employee.LastName)
	employees = append(employees, Employee{Id: employee.Id, FirstName: employee.FirstName, LastName: employee.LastName})
	json.NewEncoder(w).Encode(employees)
}

func getIndex(id string) int {
	for i := 0; i < len(employees); i++ {
		if employees[i].Id == id {
			return i
		}
	}
	return -1
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		log.Printf("error occurred while decoding employee data :: ", err)
		return
	}

	log.Printf("deleting employee id :: %s with FirstName :: %s and LastName :: %s", employee.Id, employee.FirstName, employee.LastName)
	index := getIndex(employee.Id)
	employees = append(employees[:index], employees[index+1:]...)
	json.NewEncoder(w).Encode(&employees)
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)

	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
