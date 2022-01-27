package muid

import (
  "errors"
)

var (
	errInvalid  = errors.New("id: invalid Id")
	errRandom   = errors.New("id: cannot generate random number:")
	errScanning = errors.New("id: scanning unsupported type:")
)
