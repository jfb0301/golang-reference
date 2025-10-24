func handler(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.error()))
		return
	}

	data := req.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.WriteHeader([]byte(err.error()))
		return 
	}
	rw.Write([]byte(result))
}

