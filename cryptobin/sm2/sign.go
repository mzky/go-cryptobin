package sm2

import (
    "errors"
    "math/big"
    "crypto/rand"
    "encoding/asn1"

    "github.com/deatil/go-cryptobin/gm/sm2"
    cryptobin_tool "github.com/deatil/go-cryptobin/tool"
)

// 私钥签名
func (this SM2) Sign() SM2 {
    if this.privateKey == nil {
        err := errors.New("SM2: privateKey error.")
        return this.AppendError(err)
    }

    parsedData, err := this.privateKey.Sign(rand.Reader, this.data, nil)
    if err != nil {
        return this.AppendError(err)
    }

    this.parsedData = parsedData

    return this
}

// 公钥验证
// 使用原始数据[data]对比签名后数据
func (this SM2) Verify(data []byte) SM2 {
    if this.publicKey == nil {
        err := errors.New("SM2: publicKey error.")
        return this.AppendError(err)
    }

    this.verify = this.publicKey.Verify(data, this.data, nil)

    return this
}

// ===============

type sm2Signature struct {
    R, S *big.Int
}

// 私钥签名
func (this SM2) SignAsn1(uid []byte) SM2 {
    if this.privateKey == nil {
        err := errors.New("SM2: privateKey error.")
        return this.AppendError(err)
    }

    r, s, err := sm2.Sm2Sign(rand.Reader, this.privateKey, this.data, uid)
    if err != nil {
        return this.AppendError(err)
    }

    parsedData, err := asn1.Marshal(sm2Signature{r, s})
    if err != nil {
        return this.AppendError(err)
    }

    this.parsedData = parsedData

    return this
}

// 公钥验证
// 使用原始数据[data]对比签名后数据
func (this SM2) VerifyAsn1(data []byte, uid []byte) SM2 {
    if this.publicKey == nil {
        err := errors.New("SM2: publicKey error.")
        return this.AppendError(err)
    }

    var sm2Sign sm2Signature
    _, err := asn1.Unmarshal(this.data, &sm2Sign)
    if err != nil {
        return this.AppendError(err)
    }

    this.verify = sm2.Sm2Verify(this.publicKey, data, uid, sm2Sign.R, sm2Sign.S)

    return this
}

// ===============

// 私钥签名
// 兼容[招行]
func (this SM2) SignHex(uid []byte) SM2 {
    if this.privateKey == nil {
        err := errors.New("SM2: privateKey error.")
        return this.AppendError(err)
    }

    r, s, err := sm2.Sm2Sign(rand.Reader, this.privateKey, this.data, uid)
    if err != nil {
        return this.AppendError(err)
    }

    encoding := cryptobin_tool.NewEncoding()

    rHex := encoding.HexEncode(r.Bytes())
    sHex := encoding.HexEncode(s.Bytes())

    sign := encoding.HexPadding(rHex, 64) + encoding.HexPadding(sHex, 64)

    parsedData, err := encoding.HexDecode(sign)
    if err != nil {
        return this.AppendError(err)
    }

    this.parsedData = parsedData

    return this
}

// 公钥验证
// 兼容[招行]
// 使用原始数据[data]对比签名后数据
func (this SM2) VerifyHex(data []byte, uid []byte) SM2 {
    if this.publicKey == nil {
        err := errors.New("SM2: publicKey error.")
        return this.AppendError(err)
    }

    signData := cryptobin_tool.NewEncoding().HexEncode(this.data)

    r, _ := new(big.Int).SetString(signData[:64], 16)
    s, _ := new(big.Int).SetString(signData[64:], 16)

    this.verify = sm2.Sm2Verify(this.publicKey, data, uid, r, s)

    return this
}
