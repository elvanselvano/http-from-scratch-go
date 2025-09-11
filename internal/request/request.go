package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

var ERROR_MALFORMED_REQUEST_LINE = fmt.Errorf("malformed request line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported http version")
var HTTP_SEPARATOR = "\r\n"

func parseRequestLine(s string) (*RequestLine, string, error) {
	idx := strings.Index(s, HTTP_SEPARATOR)
	if idx == -1 {
		return nil, s, nil
	}

	startLine := s[:idx]
	restOfMessage := s[idx+len(HTTP_SEPARATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, restOfMessage, ERROR_MALFORMED_REQUEST_LINE
	}

	httpParts := strings.Split(parts[2], "/")
	if len(httpParts) != 2 || httpParts[0] != "HTTP" || httpParts[1] != "1.1" {
		return nil, restOfMessage, ERROR_MALFORMED_REQUEST_LINE
	}

	rl := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   httpParts[1],
	}

	return rl, restOfMessage, nil
}

func RequestFromHeader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("unable to io.ReadAll"),
			err,
		)
	}

	str := string(data)
	rl, _, err := parseRequestLine(str)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, err
}
