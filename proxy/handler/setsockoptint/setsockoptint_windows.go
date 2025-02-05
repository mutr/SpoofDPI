//go:build windows

package setsockoptint

import "syscall"

func SetsockoptInt(fd uintptr, ttl int) {
	syscall.SetsockoptInt(syscall.Handle(fd), syscall.IPPROTO_IP, syscall.IP_TTL, ttl)
}
