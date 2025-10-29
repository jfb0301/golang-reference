type ServiceCaller struct {
	client *http.Client 
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://exmaple.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	resp, err := sc.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}
	id, err := processResponse(resp.Body)
	return id, err
}


