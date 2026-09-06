package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/quic-go/quic-go/interop/http09"
	"github.com/quic-go/quic-go/interop/utils"
)

var errUnsupported = errors.New("unsupported test case")

var tlsConf *tls.Config

func main() {
	logFile, err := os.Create("/logs/log.txt")
	if err != nil {
		fmt.Printf("Could not create log file: %s\n", err.Error())
		os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	keyLog, err := utils.GetSSLKeyLog()
	if err != nil {
		fmt.Printf("Could not create key log: %s\n", err.Error())
		os.Exit(1)
	}
	if keyLog != nil {
		defer keyLog.Close()
	}

	tlsConf = &tls.Config{
		InsecureSkipVerify: true,
		KeyLogWriter:       keyLog,
	}
	testcase := os.Getenv("TESTCASE")
	if err := runTestcase(testcase); err != nil {
		if err == errUnsupported {
			fmt.Printf("unsupported test case: %s\n", testcase)
			os.Exit(127)
		}
		fmt.Printf("Downloading files failed: %s\n", err.Error())
		os.Exit(1)
	}
}

func runTestcase(testcase string) error { _ = "STUB: not implemented"; return nil }

func runVersionNegotiationTest(r *http09.RoundTripper, urls []string) error {
	_ = "STUB: not implemented"
	return nil
}

func runMultiConnectTest(r *http09.RoundTripper, urls []string) error {
	_ = "STUB: not implemented"
	return nil
}

type sessionCache struct {
	tls.ClientSessionCache
	put chan<- struct{}
}

func newSessionCache(c tls.ClientSessionCache) (tls.ClientSessionCache, <-chan struct{}) {
	_ = "STUB: not implemented"
	return *new(tls.ClientSessionCache), nil
}

func (c *sessionCache) Put(key string, cs *tls.ClientSessionState) {
	_ = "STUB: not implemented"
	return
}

func runResumptionTest(r *http09.RoundTripper, urls []string, use0RTT bool) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadFiles(cl http.RoundTripper, urls []string, use0RTT bool) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadFile(cl http.RoundTripper, url string, use0RTT bool) error {
	_ = "STUB: not implemented"
	return nil
}
