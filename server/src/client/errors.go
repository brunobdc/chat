package client

import "errors"

var errSocketCantBeNil = errors.New("socket can't be nil")
var errFailedToReadAMessage = errors.New("failed to read a message")
var errFailedToCloseTheSocket = errors.New("failed to close the socket")
