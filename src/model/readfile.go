package model

type ReadFile struct {
	Reader func() ([]byte, error)
}
