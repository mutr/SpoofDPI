//go:build darwin

package setsockoptint

import "syscall"

func SetsockoptInt(fd uintptr, ttl int) {
	syscall.SetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_TTL, ttl)
}
