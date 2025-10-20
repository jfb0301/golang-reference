package main 

import "fmt"

// Definng own errors that include additional info for logging 


// 1. Define enumeration to represent status codes 

type Status int

const(
	InvalidLogin Status = iota + 1
	NotFound
)



// 2. Define a StatusErr to hold this value 

type StatusErr struct {
	Status Status 
	Message string 
}

func se(se StatusErr) Error() string {
	return se.Message 
}


// 3. StatusErr to provide more details 


func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr {
			Status : InvalidLogin, 
			Message : fmt.Sprintf("Invalid credentials for user %s", uid),  
		}
	}
	data, err := getData(file) 
	if err != nil {
		return nil, StatusErr{
			status : NotFound,
			Message : fmt.Sprintf("file %s not found", file), 
		}
	}
	return data, nil 
}


