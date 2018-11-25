package readfile_test

import "errors"

func MockReaderSuccess(path string) ([]byte, error) {
	return []byte("hi kk!"), nil
}

func MockReaderError(path string) ([]byte, error) {
	return nil, errors.New("Reader Error!")
}
