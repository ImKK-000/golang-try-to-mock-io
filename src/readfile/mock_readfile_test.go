package readfile_test

import "errors"

func MockReaderSuccess() ([]byte, error) {
	return []byte("hi kk!"), nil
}

func MockReaderError() ([]byte, error) {
	return nil, errors.New("Reader Error!")
}
