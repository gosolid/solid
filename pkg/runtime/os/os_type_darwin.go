//go:build darwin

package os

func os_type() (string, error) {
	return "Darwin", nil
}
