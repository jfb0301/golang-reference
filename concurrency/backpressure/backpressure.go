type PressureGauge struct {
	ch chan struct{}
}

func new(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case <- pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default: 
		return errors.New("No more capacity")
	}
}

