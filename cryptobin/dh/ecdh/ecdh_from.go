package ecdh

import (
    "crypto/rand"

    "github.com/deatil/go-cryptobin/dhd/ecdh"
    cryptobin_tool "github.com/deatil/go-cryptobin/tool"
)

// 私钥
func (this Ecdh) FromPrivateKey(key []byte) Ecdh {
    parsedKey, err := this.ParsePrivateKeyFromPEM(key)
    if err != nil {
        this.Error = err
        return this
    }

    this.privateKey = parsedKey.(*ecdh.PrivateKey)

    return this
}

// 私钥带密码
func (this Ecdh) FromPrivateKeyWithPassword(key []byte, password string) Ecdh {
    parsedKey, err := this.ParsePrivateKeyFromPEMWithPassword(key, password)
    if err != nil {
        this.Error = err
        return this
    }

    this.privateKey = parsedKey.(*ecdh.PrivateKey)

    return this
}

// 公钥
func (this Ecdh) FromPublicKey(key []byte) Ecdh {
    parsedKey, err := this.ParsePublicKeyFromPEM(key)
    if err != nil {
        this.Error = err
        return this
    }

    this.publicKey = parsedKey.(*ecdh.PublicKey)

    return this
}

// 根据私钥 x, y 生成
func (this Ecdh) FromKeyXYHexString(curve string, xString string, yString string) Ecdh {
    var c ecdh.Curve

    switch curve {
        case "P521":
            c = ecdh.P521()
        case "P384":
            c = ecdh.P384()
        case "P256":
            c = ecdh.P256()
        case "P224":
            c = ecdh.P224()
        default:
            c = ecdh.P224()
    }

    encoding := cryptobin_tool.NewEncoding()

    x, _ := encoding.HexDecode(xString)
    y, _ := encoding.HexDecode(yString)

    priv := &ecdh.PrivateKey{}
    priv.X = x
    priv.PublicKey.Y = y
    priv.PublicKey.Curve = c

    this.privateKey = priv
    this.publicKey  = &priv.PublicKey

    return this
}

// 生成密钥
func (this Ecdh) GenerateKey(curve string) Ecdh {
    var c ecdh.Curve

    switch curve {
        case "P521":
            c = ecdh.P521()
        case "P384":
            c = ecdh.P384()
        case "P256":
            c = ecdh.P256()
        case "P224":
            c = ecdh.P224()
        default:
            c = ecdh.P224()
    }

    this.privateKey, this.publicKey, this.Error = ecdh.GenerateKey(c, rand.Reader)

    return this
}
