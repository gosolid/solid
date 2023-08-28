//go:build !darwin && !windows

package os

import (
	"strings"

	"golang.org/x/sys/unix"
)

func os_type() (string, error) {
	var utsname unix.Utsname
	if err := unix.Uname(&utsname); err != nil {
		return "", err
	}
	return arrayToString(utsname.Sysname), nil
}

func arrayToString(x [256]byte) string {
	var buf [65]byte
	for i, b := range x {
		buf[i] = byte(b)
	}
	str := string(buf[:])
	if i := strings.Index(str, "\x00"); i != -1 {
		str = str[:i]
	}
	return str
}
