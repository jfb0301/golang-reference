package main

import (
	"fmt"
	"net/http"
	"errors"
)

// Logger
func LogOutput(message string) {
	fmt.Println(message)
}

// Data Store
type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// Factory function for instance of NewSimpleDataStore
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Robert",
			"2": "Robert Martin",
			"3": "Robert C",
		},
	}
}

// BL - Interfaces describe what it depends on
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Logger interface
type Logger interface {
	Log(message string)
}

// Make LogOutput meet the Logger interface: define function type with a method on it
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// Dependencies defined - Implement BL
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("In Say Hello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("Unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("In Say Goodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("Unknown user")
	}
	return "Goodbye, " + name, nil
}

// SimpleLogic instance
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// API
type Logic interface {
	SayHello(userID string) (string, error)
	SayGoodbye(userID string) (string, error) // Added to match implementation
}

// Controller
type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(message))
}

// Factory function for Controller
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

// Start server

/*
	Main: Only part of the code that knows what all the concrete types actually are 
*/

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}