package learn

import (
	"fmt"
	"os"
)

//====================================
// zero initialization
//
//When you declare a new variable, or
//when you create a value with the new() builtin
//function, it is initialized to the zero value for
//the type.

type Logger struct {
	out *os.File
}

func (l Logger) Log(s string) {
	out := l.out
	if out == nil {
		out = os.Stderr
	}
	fmt.Fprintf(out, "%s [%d]: %s\n", os.Args[0],
		os.Getpid(), s)
}

func (l *Logger) SetOutput(out *os.File) {
	l.out = out
}
