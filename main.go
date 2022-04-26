//go:build !windows
// +build !windows

package main

import (
	"go-timer/timer"
)

func main() {
	timer.New()
}
