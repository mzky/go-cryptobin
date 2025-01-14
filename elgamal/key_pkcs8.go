package elgamal

import (
    "fmt"
    "errors"
    "math/big"
    "encoding/asn1"
    "crypto/x509/pkix"

    "golang.org/x/crypto/cryptobyte"
    cryptobyte_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

var (
    // elgamal 公钥 oid
    oidPublicKeyEIGamal = asn1.ObjectIdentifier{1, 3, 14, 7, 2, 1, 1}
)

// elgamal Parameters
type elgamalAlgorithmParameters struct {
    G, P *big.Int
}

// 私钥 - 包装
type pkcs8 struct {
    Version    int
    Algo       pkix.AlgorithmIdentifier
    PrivateKey []byte
}

// 公钥 - 包装
type pkixPublicKey struct {
    Algo      pkix.AlgorithmIdentifier
    BitString asn1.BitString
}

// 公钥信息 - 解析
type publicKeyInfo struct {
    Raw       asn1.RawContent
    Algorithm pkix.AlgorithmIdentifier
    PublicKey asn1.BitString
}

var (
    defaultPKCS8Key = NewPKCS8Key()
)

/**
 * elgamal pkcs8 密钥
 *
 * @create 2023-6-16
 * @author deatil
 */
type PKCS8Key struct {}

// 构造函数
func NewPKCS8Key() PKCS8Key {
    return PKCS8Key{}
}

// PKCS8 包装公钥
func (this PKCS8Key) MarshalPublicKey(key *PublicKey) ([]byte, error) {
    var publicKeyBytes []byte
    var publicKeyAlgorithm pkix.AlgorithmIdentifier
    var err error

    // params
    paramBytes, err := asn1.Marshal(elgamalAlgorithmParameters{
        G: key.G,
        P: key.P,
    })
    if err != nil {
        return nil, errors.New("elgamal: failed to marshal algo param: " + err.Error())
    }

    publicKeyAlgorithm.Algorithm = oidPublicKeyEIGamal
    publicKeyAlgorithm.Parameters.FullBytes = paramBytes

    var yInt cryptobyte.Builder
    yInt.AddASN1BigInt(key.Y)

    publicKeyBytes, err = yInt.Bytes()
    if err != nil {
        return nil, errors.New("elgamal: failed to builder PrivateKey: " + err.Error())
    }

    pkix := pkixPublicKey{
        Algo: publicKeyAlgorithm,
        BitString: asn1.BitString{
            Bytes:     publicKeyBytes,
            BitLength: 8 * len(publicKeyBytes),
        },
    }

    return asn1.Marshal(pkix)
}

// PKCS8 包装公钥
func MarshalPKCS8PublicKey(pub *PublicKey) ([]byte, error) {
    return defaultPKCS8Key.MarshalPublicKey(pub)
}

// PKCS8 解析公钥
func (this PKCS8Key) ParsePublicKey(der []byte) (*PublicKey, error) {
    var pki publicKeyInfo
    rest, err := asn1.Unmarshal(der, &pki)
    if err != nil {
        return nil, err
    }

    if len(rest) > 0 {
        return nil, asn1.SyntaxError{Msg: "trailing data"}
    }

    algoEq := pki.Algorithm.Algorithm.Equal(oidPublicKeyEIGamal)
    if !algoEq {
        return nil, errors.New("elgamal: unknown public key algorithm")
    }

    // 解析
    keyData := &pki

    yDer := cryptobyte.String(keyData.PublicKey.RightAlign())

    y := new(big.Int)
    if !yDer.ReadASN1Integer(y) {
        return nil, errors.New("elgamal: invalid EIGamal public key")
    }

    pub := &PublicKey{
        G: new(big.Int),
        P: new(big.Int),
        Y: y,
    }

    paramsDer := cryptobyte.String(keyData.Algorithm.Parameters.FullBytes)
    if !paramsDer.ReadASN1(&paramsDer, cryptobyte_asn1.SEQUENCE) ||
        !paramsDer.ReadASN1Integer(pub.G) ||
        !paramsDer.ReadASN1Integer(pub.P) {
        return nil, errors.New("elgamal: invalid EIGamal public key")
    }

    if pub.Y.Sign() <= 0 ||
        pub.G.Sign() <= 0 ||
        pub.P.Sign() <= 0 {
        return nil, errors.New("elgamal: zero or negative EIGamal parameter")
    }

    return pub, nil
}

// PKCS8 解析公钥
func ParsePKCS8PublicKey(derBytes []byte) (*PublicKey, error) {
    return defaultPKCS8Key.ParsePublicKey(derBytes)
}

// ====================

// PKCS8 包装私钥
func (this PKCS8Key) MarshalPrivateKey(key *PrivateKey) ([]byte, error) {
    var privKey pkcs8

    // params
    paramBytes, err := asn1.Marshal(elgamalAlgorithmParameters{
        G: key.G,
        P: key.P,
    })
    if err != nil {
        return nil, errors.New("elgamal: failed to marshal algo param: " + err.Error())
    }

    privKey.Algo = pkix.AlgorithmIdentifier{
        Algorithm:  oidPublicKeyEIGamal,
        Parameters: asn1.RawValue{
            FullBytes: paramBytes,
        },
    }

    var xInt cryptobyte.Builder
    xInt.AddASN1BigInt(key.X)

    privateKeyBytes, err := xInt.Bytes()
    if err != nil {
        return nil, errors.New("elgamal: failed to builder PrivateKey: " + err.Error())
    }

    privKey.PrivateKey = privateKeyBytes

    return asn1.Marshal(privKey)
}

// PKCS8 包装私钥
func MarshalPKCS8PrivateKey(key *PrivateKey) ([]byte, error) {
    return defaultPKCS8Key.MarshalPrivateKey(key)
}

// PKCS8 解析私钥
func (this PKCS8Key) ParsePrivateKey(der []byte) (key *PrivateKey, err error) {
    var privKey pkcs8
    _, err = asn1.Unmarshal(der, &privKey)
    if err != nil {
        return nil, err
    }

    if !privKey.Algo.Algorithm.Equal(oidPublicKeyEIGamal) {
        return nil, fmt.Errorf("elgamal: PKCS#8 wrapping contained private key with unknown algorithm: %v", privKey.Algo.Algorithm)
    }

    xDer := cryptobyte.String(string(privKey.PrivateKey))

    x := new(big.Int)
    if !xDer.ReadASN1Integer(x) {
        return nil, errors.New("elgamal: invalid EIGamal public key")
    }

    priv := &PrivateKey{
        PublicKey: PublicKey{
            G: new(big.Int),
            P: new(big.Int),
            Y: new(big.Int),
        },
        X: x,
    }

    // 找出 g, p 数据
    paramsDer := cryptobyte.String(privKey.Algo.Parameters.FullBytes)
    if !paramsDer.ReadASN1(&paramsDer, cryptobyte_asn1.SEQUENCE) ||
        !paramsDer.ReadASN1Integer(priv.G) ||
        !paramsDer.ReadASN1Integer(priv.P) {
        return nil, errors.New("elgamal: invalid EIGamal private key")
    }

    // 算出 Y 值
    priv.Y.Exp(priv.G, priv.X, priv.P)

    if priv.Y.Sign() <= 0 || priv.G.Sign() <= 0 ||
        priv.P.Sign() <= 0 || priv.X.Sign() <= 0 {
        return nil, errors.New("elgamal: zero or negative EIGamal parameter")
    }

    return priv, nil
}

// PKCS8 解析私钥
func ParsePKCS8PrivateKey(derBytes []byte) (key *PrivateKey, err error) {
    return defaultPKCS8Key.ParsePrivateKey(derBytes)
}
