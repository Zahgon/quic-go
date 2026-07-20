package main

import (
	"crypto/tls"
	"io"
	"log"
)

const addr = "localhost:4242"

const message = "foobar"

func main() {
	go func() { log.Fatal(echoServer()) }()

	if err := clientMain(); err != nil {
		panic(err)
	}
}

func echoServer() error { _ = "STUB: not implemented"; return nil }

func clientMain() error { _ = "STUB: not implemented"; return nil }

type loggingWriter struct{ io.Writer }

func (w loggingWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func generateTLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }
