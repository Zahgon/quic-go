package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "net/http/pprof"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/quic-go/http3/qlog"
	"github.com/quic-go/quic-go/internal/testdata"
)

type binds []string

func (b binds) String() string { _ = "STUB: not implemented"; return "" }

func (b *binds) Set(v string) error { _ = "STUB: not implemented"; return nil }

type Size interface {
	Size() int64
}

func generatePRData(l int) []byte { _ = "STUB: not implemented"; return nil }

func setupHandler(www string) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	bs := binds{}
	flag.Var(&bs, "bind", "bind to")
	www := flag.String("www", "", "www data")
	tcp := flag.Bool("tcp", false, "also listen on TCP")
	key := flag.String("key", "", "TLS key (requires -cert option)")
	cert := flag.String("cert", "", "TLS certificate (requires -key option)")
	flag.Parse()

	if len(bs) == 0 {
		bs = binds{"localhost:6121"}
	}

	handler := setupHandler(*www)

	var wg sync.WaitGroup
	var certFile, keyFile string
	if *key != "" && *cert != "" {
		keyFile = *key
		certFile = *cert
	} else {
		certFile, keyFile = testdata.GetCertificatePaths()
	}
	for _, b := range bs {
		fmt.Println("listening on", b)
		bCap := b
		wg.Go(func() {
			var err error
			if *tcp {
				err = http3.ListenAndServeTLS(bCap, certFile, keyFile, handler)
			} else {
				server := http3.Server{
					Handler: handler,
					Addr:    bCap,
					QUICConfig: &quic.Config{
						Tracer: qlog.DefaultConnectionTracer,
					},
				}
				err = server.ListenAndServeTLS(certFile, keyFile)
			}
			if err != nil {
				fmt.Println(err)
			}
		})
	}
	wg.Wait()
}
