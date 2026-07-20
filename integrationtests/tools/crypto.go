package tools

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"time"
)

const ALPN = "quic-go integration tests"

var (
	notBefore = time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	notAfter  = time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)
)

func GenerateCA() (*x509.Certificate, crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.PrivateKey), nil
}

func GenerateLeafCert(ca *x509.Certificate, caPriv crypto.PrivateKey) (*x509.Certificate, crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.PrivateKey), nil
}

func GenerateTLSConfigWithLongCertChain(ca *x509.Certificate, caPrivateKey crypto.PrivateKey) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
