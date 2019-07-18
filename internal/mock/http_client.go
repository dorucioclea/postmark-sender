package mock

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	SomeError = errors.New("some error")
)

type IoReaderError struct{}

type TransportStatusIncorrectResponse struct {
	Transport http.RoundTripper
}

type TransportStatusBadStatus struct {
	Transport http.RoundTripper
}

type TransportStatusErrorIoReader struct {
	Transport http.RoundTripper
}

func NewClientStatusIncorrectResponse() *http.Client {
	return &http.Client{
		Transport: &TransportStatusIncorrectResponse{},
	}
}

func NewClientStatusErrorIoReader() *http.Client {
	return &http.Client{
		Transport: &TransportStatusErrorIoReader{},
	}
}

func NewClientStatusBadStatus() *http.Client {
	return &http.Client{
		Transport: &TransportStatusBadStatus{},
	}
}

func NewIoReaderError() io.Reader {
	return &IoReaderError{}
}

func (h *TransportStatusIncorrectResponse) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("some_not_json_string")),
		Header:     make(http.Header),
	}, nil
}

func (h *TransportStatusErrorIoReader) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(NewIoReaderError()),
		Header:     make(http.Header),
	}, nil
}

func (h *TransportStatusBadStatus) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader(`{"ErrorCode": 123, "Message": "some error"}`)),
		Header:     make(http.Header),
	}, nil
}

func (r *IoReaderError) Read(p []byte) (int, error) {
	return 0, SomeError
}
