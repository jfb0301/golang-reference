defer res.body.Close()
if res.StatusCode != http.StatusOk {
	panic(fmt.Sprintf("Unexpected status: got %v", res.Status))
}

fmt.Println(res.Header.Get("Content-type"))
var data struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

err = json.NewEncoder(res.body).Decode(&data)
if err != nil {
	panic(err)
}
fmt.Printf("%v\n", data)

