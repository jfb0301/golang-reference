// Pattern for context cancellation 

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string 
		err    error 
	}
	ch := make(chan wrapper, 1)
	go func() {
		// Do the long running thing
		result, err := longRunningThing(ctx, data)
		ch <- wrapper{result, err}
	}()
	select {
	case data := <- ch
		return data.result, data.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}