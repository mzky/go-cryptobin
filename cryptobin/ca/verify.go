package ca

import (
    "errors"
    "crypto/x509"
    "encoding/pem"

    sm2X509 "github.com/deatil/go-cryptobin/gm/x509"
)

// 验证
func (this CA) Verify(rootPEM string, certPEM string, opts x509.VerifyOptions) (bool, error) {
    roots := x509.NewCertPool()
    ok := roots.AppendCertsFromPEM([]byte(rootPEM))
    if !ok {
        return false, errors.New("CA: failed to parse root certificate")
    }

    block, _ := pem.Decode([]byte(certPEM))
    if block == nil {
        return false, errors.New("CA: failed to parse certificate PEM")
    }

    cert, err := x509.ParseCertificate(block.Bytes)
    if err != nil {
        return false, errors.New("CA: failed to parse certificate: " + err.Error())
    }

    // 重设
    opts.Roots = roots

    if _, err := cert.Verify(opts); err != nil {
        return false, errors.New("CA: failed to verify certificate: " + err.Error())
    }

    return true, nil
}

// SM2 验证
func (this CA) SM2Verify(rootPEM string, certPEM string, opts sm2X509.VerifyOptions) (bool, error) {
    roots := sm2X509.NewCertPool()
    ok := roots.AppendCertsFromPEM([]byte(rootPEM))
    if !ok {
        return false, errors.New("failed to parse root certificate")
    }

    block, _ := pem.Decode([]byte(certPEM))
    if block == nil {
        return false, errors.New("failed to parse certificate PEM")
    }

    cert, err := sm2X509.ParseCertificate(block.Bytes)
    if err != nil {
        return false, errors.New("failed to parse certificate: " + err.Error())
    }

    // 重设
    opts.Roots = roots

    if _, err := cert.Verify(opts); err != nil {
        return false, errors.New("failed to verify certificate: " + err.Error())
    }

    return true, nil
}
