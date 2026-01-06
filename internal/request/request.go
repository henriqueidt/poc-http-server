package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion string
	RequestTarget string
	Method string
}
type Request struct {
	RequestLine RequestLine
}

var ERROR_BAD_REQUEST = fmt.Errorf("Bad Request Line")
var ERROR_HTTP_VERSION_NOT_SUPPORTED = fmt.Errorf("unsupported http version")
var SEPARATOR = "\r\n"

func parseRequestLine(line string) (*RequestLine, string, error) {
	index := strings.Index(line, SEPARATOR)
	if index == -1 {
		return nil, line, nil
	}

	startLine := line[:index]
	rest := line[index + len(SEPARATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, rest, ERROR_BAD_REQUEST
	}

	httpParts := strings.Split(parts[2], "/")

	if len(httpParts) != 2 || httpParts[0] != "HTTP" || httpParts[1] != "1.1" {
		return nil, rest, ERROR_HTTP_VERSION_NOT_SUPPORTED
	}

	requestLine := &RequestLine{
		HttpVersion: httpParts[1],
		RequestTarget: parts[1],
		Method: parts[0],
	}

	return requestLine, rest, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)

	if err != nil {
		return nil, errors.Join(fmt.Errorf("unable to read"), err)
	}

	str := string(data)

	requestLine, _, err := parseRequestLine(str)

	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *requestLine,
	}, err
}