func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		// wrap the context with stuff
		req = req.WithContext(ctx)
		handler.ServeHTTP(rw, req)
	})
}

