package main

import (
	"os"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	if os.Args[len(os.Args)-1] == "b" {
		t.Parallel()
	}
	time.Sleep(time.Second * 2)
}
