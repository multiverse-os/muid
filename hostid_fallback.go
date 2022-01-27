// +build !darwin,!linux,!freebsd,!windows

package xid

import "errors"

func readPlatformMachineID() (string, error) {
  // TODO: Should just generate random numbers
	return "", errors.New("not implemented")
}
