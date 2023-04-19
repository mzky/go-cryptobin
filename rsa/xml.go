package rsa

import (
    "errors"
    "math/big"
    "crypto/rsa"
    "encoding/xml"
    "encoding/base64"
)

// 私钥
type xmlPrivateKey struct {
    XMLName  xml.Name `xml:"RSAKeyValue"`
    Modulus  string   `xml:"Modulus"`
    Exponent string   `xml:"Exponent"`
    D        string   `xml:"D"`
    P        string   `xml:"P"`
    Q        string   `xml:"Q"`
    DP       string   `xml:"DP"`
    DQ       string   `xml:"DQ"`
    InverseQ string   `xml:"InverseQ"`
}

// 公钥
type xmlPublicKey struct {
    XMLName  xml.Name `xml:"RSAKeyValue"`
    Modulus  string   `xml:"Modulus"`
    Exponent string   `xml:"Exponent"`
}

// 构造函数
func NewXMLKey() XMLKey {
    return XMLKey{}
}

var defaultXMLKey = NewXMLKey()

/**
 * rsa xml密钥
 *
 * @create 2023-4-10
 * @author deatil
 */
type XMLKey struct {}

// 包装公钥
func (this XMLKey) MarshalPublicKey(key *rsa.PublicKey) ([]byte, error) {
    publicKey := xmlPublicKey{
        Modulus:  this.bigintToB64(key.N),
        Exponent: this.bigintToB64(big.NewInt(int64(key.E))),
    }

    return xml.MarshalIndent(publicKey, "", "    ")
}

func MarshalXMLPublicKey(key *rsa.PublicKey) ([]byte, error) {
    return defaultXMLKey.MarshalPublicKey(key)
}

// 解析公钥
func (this XMLKey) ParsePublicKey(der []byte) (*rsa.PublicKey, error) {
    var pub xmlPublicKey
    err := xml.Unmarshal(der, &pub)
    if err != nil {
        return nil, err
    }

    publicKey := &rsa.PublicKey{
        N: this.b64ToBigint(pub.Modulus),
        E: int(this.b64ToBigint(pub.Exponent).Int64()),
    }

    return publicKey, nil
}

func ParseXMLPublicKey(der []byte) (*rsa.PublicKey, error) {
    return defaultXMLKey.ParsePublicKey(der)
}

// ====================

// 包装私钥
func (this XMLKey) MarshalPrivateKey(key *rsa.PrivateKey) ([]byte, error) {
    key.Precompute()

    // 构造私钥信息
    priv := xmlPrivateKey{
        Modulus:  this.bigintToB64(key.N),
        Exponent: this.bigintToB64(big.NewInt(int64(key.E))),
        D:        this.bigintToB64(key.D),
        P:        this.bigintToB64(key.Primes[0]),
        Q:        this.bigintToB64(key.Primes[1]),
        DP:       this.bigintToB64(key.Precomputed.Dp),
        DQ:       this.bigintToB64(key.Precomputed.Dq),
        InverseQ: this.bigintToB64(key.Precomputed.Qinv),
    }

    return xml.MarshalIndent(priv, "", "    ")
}

func MarshalXMLPrivateKey(key *rsa.PrivateKey) ([]byte, error) {
    return defaultXMLKey.MarshalPrivateKey(key)
}

// 解析私钥
func (this XMLKey) ParsePrivateKey(der []byte) (*rsa.PrivateKey, error) {
    var priv xmlPrivateKey
    err := xml.Unmarshal(der, &priv)
    if err != nil {
        return nil, err
    }

    e := int(this.b64ToBigint(priv.Exponent).Int64())
    n := this.b64ToBigint(priv.Modulus)
    d := this.b64ToBigint(priv.D)
    p := this.b64ToBigint(priv.P)
    q := this.b64ToBigint(priv.Q)

    if n.Sign() <= 0 || d.Sign() <= 0 || p.Sign() <= 0 || q.Sign() <= 0 {
        return nil, errors.New("rsa xml: private key contains zero or negative value")
    }

    key := new(rsa.PrivateKey)
    key.PublicKey = rsa.PublicKey{
        N: n,
        E: e,
    }

    key.D = d
    key.Primes = make([]*big.Int, 2)
    key.Primes[0] = p
    key.Primes[1] = q

    err = key.Validate()
    if err != nil {
        return nil, err
    }

    key.Precompute()

    return key, nil
}

func ParseXMLPrivateKey(der []byte) (*rsa.PrivateKey, error) {
    return defaultXMLKey.ParsePrivateKey(der)
}

// ====================

func (this XMLKey) b64d(str string) []byte {
    decoded, _ := base64.StdEncoding.DecodeString(str)

    return []byte(decoded)
}

func (this XMLKey) b64e(src []byte) string {
    return base64.StdEncoding.EncodeToString(src)
}

func (this XMLKey) b64ToBigint(str string) *big.Int {
    bInt := &big.Int{}
    bInt.SetBytes(this.b64d(str))
    return bInt
}

// big.NewInt(int64)
func (this XMLKey) bigintToB64(encoded *big.Int) string {
    return this.b64e(encoded.Bytes())
}
