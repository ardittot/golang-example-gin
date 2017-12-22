package main

// Variable Types Declaration
type Person struct {
    ID        string   `json:"id"`
    Firstname string   `json:"firstname"`
    Lastname  string   `json:"lastname"`
    Address   *Address `json:"address"`
}

type Address struct {
    City  string `json:"city"`
    State string `json:"state"`
}

// Variables Declaration
var people []Person

// Functions Declaration
func InitPeople(people_in []Person) []Person {
	people = people_in
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	return people
}
