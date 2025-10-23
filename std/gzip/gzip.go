func buildGZipReader(filename string) (*gzip.Reader, func(), error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.close()
		r.close()
	}, nil 
}

r, closer, err := buildGZipReader("my_data_text.gz")
if err != nil {
	return err 
}

defer close()
counts, err := countLetters(r)
if err != nil {
	return err 
}

fmt.Println(counts)