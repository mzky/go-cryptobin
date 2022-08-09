package curve25519

import (
    "crypto/rand"

    "github.com/deatil/go-cryptobin/dhd/curve25519"
    cryptobin_tool "github.com/deatil/go-cryptobin/tool"
)

// 私钥
func (this Curve25519) FromPrivateKey(key []byte) Curve25519 {
    parsedKey, err := this.ParsePrivateKeyFromPEM(key)
    if err != nil {
        this.Error = err
        return this
    }

    this.privateKey = parsedKey.(*curve25519.PrivateKey)

    return this
}

// 私钥带密码
func (this Curve25519) FromPrivateKeyWithPassword(key []byte, password string) Curve25519 {
    parsedKey, err := this.ParsePrivateKeyFromPEMWithPassword(key, password)
    if err != nil {
        this.Error = err
        return this
    }

    this.privateKey = parsedKey.(*curve25519.PrivateKey)

    return this
}

// 公钥
func (this Curve25519) FromPublicKey(key []byte) Curve25519 {
    parsedKey, err := this.ParsePublicKeyFromPEM(key)
    if err != nil {
        this.Error = err
        return this
    }

    this.publicKey = parsedKey.(*curve25519.PublicKey)

    return this
}

// 根据私钥 x, y 生成
func (this Curve25519) FromKeyXYHexString(xString string, yString string) Curve25519 {
    encoding := cryptobin_tool.NewEncoding()

    x, _ := encoding.HexDecode(xString)
    y, _ := encoding.HexDecode(yString)

    priv := &curve25519.PrivateKey{}
    priv.X = x
    priv.PublicKey.Y = y

    this.privateKey = priv
    this.publicKey  = &priv.PublicKey

    return this
}

// 生成密钥
func (this Curve25519) GenerateKey() Curve25519 {
    this.privateKey, this.publicKey, this.Error = curve25519.GenerateKey(rand.Reader)

    return this
}
