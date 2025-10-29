// Middleware extract user ID from a coockie 

func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Coockie("user")
	if err != nil {
		return "", err
	}
	return userCoockie.Value, nil
}

func MiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		user, err := extractUser(req)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := req.Context()
		ctx := ContextWithUsers(ctx, user)
		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}

