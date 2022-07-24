package cryptobin

import (
    "errors"
    "crypto/ecdsa"
    "crypto/x509"
    "encoding/pem"
)

// 私钥
func (this Ecdsa) CreatePrivateKey() Ecdsa {
    if this.privateKey == nil {
        this.Error = errors.New("privateKey error.")

        return this
    }

    x509PrivateKey, err := x509.MarshalECPrivateKey(this.privateKey)
    if err != nil {
        this.Error = err
        return this
    }

    privateBlock := &pem.Block{
        Type: "EC PRIVATE KEY",
        Bytes: x509PrivateKey,
    }

    this.keyData = pem.EncodeToMemory(privateBlock)

    return this
}

// 公钥
func (this Ecdsa) CreatePublicKey() Ecdsa {
    var publicKey *ecdsa.PublicKey

    if this.publicKey == nil {
        if this.privateKey == nil {
            this.Error = errors.New("privateKey error.")

            return this
        }

        publicKey = &this.privateKey.PublicKey
    } else {
        publicKey = this.publicKey
    }

    x509PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        this.Error = err
        return this
    }

    publicBlock := &pem.Block{
        Type: "PUBLIC KEY",
        Bytes: x509PublicKey,
    }

    this.keyData = pem.EncodeToMemory(publicBlock)

    return this
}
