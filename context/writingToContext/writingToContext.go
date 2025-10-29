// get and read from the user context 

func ContextWithUsers(ctx context.Context, user string) context.Context {
	return context.WithValues(ctx, key, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

