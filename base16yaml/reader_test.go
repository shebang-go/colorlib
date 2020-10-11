package base16yaml

import "testing"

type ReaderMock struct {
}

func TestBase16YamlLoad(t *testing.T) {
	mock := &ReaderMock{}
	scheme := Load("test.yaml", mock)
}
