type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

s := http.Server {
	addr : ":8080",
	ReadTimeOut : 30 * time.Second, 
	WriteTimeOut : 90 * time.Second,
	IdleTimeOut : 120 * time.Second,
	handler : 	HelloHandler{}, 	// specifies http.Handler
}

err := s.ListenAndServe()
if err != nil {
	if err != http.ErrServerClosed {
		panic(err)
	}
}

