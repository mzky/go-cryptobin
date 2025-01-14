package gost

import (
    "fmt"
    "errors"
    "encoding/asn1"
    "crypto/x509/pkix"

    "golang.org/x/crypto/cryptobyte"
)

var (
    oidPublicKeyGOST = asn1.ObjectIdentifier{1, 2, 643, 2, 2, 19}

    oidGostR3410_2001_TestParamSet         = asn1.ObjectIdentifier{1, 2, 643, 2, 2, 35, 0}

    oidTc26_gost_3410_12_256_paramSetA     = asn1.ObjectIdentifier{1, 2, 643, 7, 1, 2, 1, 1, 1}
    oidGostR3410_2001_CryptoPro_A_ParamSet = asn1.ObjectIdentifier{1, 2, 643, 2, 2, 35, 1}
    oidGostR3410_2001_CryptoPro_B_ParamSet = asn1.ObjectIdentifier{1, 2, 643, 2, 2, 35, 2}
    oidGostR3410_2001_CryptoPro_C_ParamSet = asn1.ObjectIdentifier{1, 2, 643, 2, 2, 35, 3}

    oidTc26_gost_3410_12_512_paramSetA = asn1.ObjectIdentifier{1, 2, 643, 7, 1, 2, 1, 2, 1}
    oidTc26_gost_3410_12_512_paramSetB = asn1.ObjectIdentifier{1, 2, 643, 7, 1, 2, 1, 2, 2}
    oidTc26_gost_3410_12_512_paramSetC = asn1.ObjectIdentifier{1, 2, 643, 7, 1, 2, 1, 2, 3}

    /* OID for EC DH */
    oidGostR3410_2001_CryptoPro_XchA_ParamSet = asn1.ObjectIdentifier{1, 2, 643, 2, 2, 36, 0}
    oidGostR3410_2001_CryptoPro_XchB_ParamSet = asn1.ObjectIdentifier{1, 2, 643, 2, 2, 36, 1}
)

func init() {
    AddNamedCurve(CurveIdGostR34102001TestParamSet(), oidGostR3410_2001_TestParamSet)

    AddNamedCurve(CurveIdtc26gost341012256paramSetA(), oidTc26_gost_3410_12_256_paramSetA)
    AddNamedCurve(CurveIdGostR34102001CryptoProAParamSet(), oidGostR3410_2001_CryptoPro_A_ParamSet)
    AddNamedCurve(CurveIdGostR34102001CryptoProBParamSet(), oidGostR3410_2001_CryptoPro_B_ParamSet)
    AddNamedCurve(CurveIdGostR34102001CryptoProCParamSet(), oidGostR3410_2001_CryptoPro_C_ParamSet)

    AddNamedCurve(CurveIdtc26gost341012512paramSetA(), oidTc26_gost_3410_12_512_paramSetA)
    AddNamedCurve(CurveIdtc26gost341012512paramSetB(), oidTc26_gost_3410_12_512_paramSetB)
    AddNamedCurve(CurveIdtc26gost341012512paramSetC(), oidTc26_gost_3410_12_512_paramSetC)

    AddNamedCurve(CurveIdGostR34102001CryptoProXchAParamSet(), oidGostR3410_2001_CryptoPro_XchA_ParamSet)
    AddNamedCurve(CurveIdGostR34102001CryptoProXchBParamSet(), oidGostR3410_2001_CryptoPro_XchB_ParamSet)
}

const gostPrivKeyVersion = 1

// pkcs8 data
type pkcs8 struct {
    Version    int
    Algo       pkix.AlgorithmIdentifier
    PrivateKey []byte
}

// PublicKey data
type pkixPublicKey struct {
    Algo      pkix.AlgorithmIdentifier
    BitString asn1.BitString
}

// publicKeyInfo parse
type publicKeyInfo struct {
    Raw       asn1.RawContent
    Algorithm pkix.AlgorithmIdentifier
    PublicKey asn1.BitString
}

// Per RFC 5915 the NamedCurveOID is marked as ASN.1 OPTIONAL, however in
// most cases it is not.
type gostPrivateKey struct {
    Version       int
    PrivateKey    []byte
    NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
    PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

// Marshal PublicKey
func MarshalPublicKey(pub *PublicKey) ([]byte, error) {
    var publicKeyBytes []byte
    var publicKeyAlgorithm pkix.AlgorithmIdentifier
    var err error

    oid, ok := OidFromNamedCurve(pub.Curve)
    if !ok {
        return nil, errors.New("gost: unsupported gost curve")
    }

    var paramBytes []byte
    paramBytes, err = asn1.Marshal(oid)
    if err != nil {
        return nil, err
    }

    publicKeyAlgorithm.Algorithm = oidPublicKeyGOST
    publicKeyAlgorithm.Parameters.FullBytes = paramBytes

    if !pub.Curve.IsOnCurve(pub.X, pub.Y) {
        return nil, errors.New("gost: invalid gost curve public key")
    }

    publicKeyBytes = ToPublicKey(pub)

    pkix := pkixPublicKey{
        Algo: publicKeyAlgorithm,
        BitString: asn1.BitString{
            Bytes:     publicKeyBytes,
            BitLength: 8 * len(publicKeyBytes),
        },
    }

    return asn1.Marshal(pkix)
}

// Parse PublicKey
func ParsePublicKey(derBytes []byte) (pub *PublicKey, err error) {
    var pki publicKeyInfo
    rest, err := asn1.Unmarshal(derBytes, &pki)
    if err != nil {
        return
    } else if len(rest) != 0 {
        err = errors.New("gost: trailing data after ASN.1 of public-key")
        return
    }

    if len(rest) > 0 {
        err = asn1.SyntaxError{Msg: "trailing data"}
        return
    }

    // parse
    keyData := &pki

    oid := keyData.Algorithm.Algorithm
    params := keyData.Algorithm.Parameters
    der := cryptobyte.String(keyData.PublicKey.RightAlign())

    algoEq := oid.Equal(oidPublicKeyGOST)
    if !algoEq {
        err = errors.New("gost: unknown public key algorithm")
        return
    }

    paramsDer := cryptobyte.String(params.FullBytes)
    namedCurveOID := new(asn1.ObjectIdentifier)
    if !paramsDer.ReadASN1ObjectIdentifier(namedCurveOID) {
        return nil, errors.New("gost: invalid ECDH parameters")
    }

    namedCurve := NamedCurveFromOid(*namedCurveOID)
    if namedCurve == nil {
        err = errors.New("gost: unsupported gost curve")
        return
    }

    pub, err = NewPublicKey(namedCurve, der)
    if err != nil {
        err = errors.New("gost: failed to unmarshal gost curve point")
        return
    }

    return
}

// ====================

// Marshal PrivateKey
func MarshalPrivateKey(key *PrivateKey) ([]byte, error) {
    var privKey pkcs8

    oid, ok := OidFromNamedCurve(key.Curve)
    if !ok {
        return nil, errors.New("gost: unsupported gost curve")
    }

    // Marshal oid
    oidBytes, err := asn1.Marshal(oid)
    if err != nil {
        return nil, errors.New("gost: failed to marshal algo param: " + err.Error())
    }

    privKey.Algo = pkix.AlgorithmIdentifier{
        Algorithm:  oidPublicKeyGOST,
        Parameters: asn1.RawValue{
            FullBytes: oidBytes,
        },
    }

    if !key.Curve.IsOnCurve(key.X, key.Y) {
        return nil, errors.New("invalid elliptic key public key")
    }

    privKey.PrivateKey, err = marshalGostPrivateKeyWithOID(key, oid)

    if err != nil {
        return nil, errors.New("gost: failed to marshal EC private key while building PKCS#8: " + err.Error())
    }

    return asn1.Marshal(privKey)
}

// Parse PrivateKey
func ParsePrivateKey(derBytes []byte) (*PrivateKey, error) {
    var privKey pkcs8
    var err error

    _, err = asn1.Unmarshal(derBytes, &privKey)
    if err != nil {
        return nil, err
    }

    algoEq := privKey.Algo.Algorithm.Equal(oidPublicKeyGOST)
    if !algoEq {
        err = errors.New("gost: unknown private key algorithm")
        return nil, err
    }

    bytes := privKey.Algo.Parameters.FullBytes

    namedCurveOID := new(asn1.ObjectIdentifier)
    if _, err := asn1.Unmarshal(bytes, namedCurveOID); err != nil {
        namedCurveOID = nil
    }

    key, err := parseGostPrivateKey(namedCurveOID, privKey.PrivateKey)
    if err != nil {
        return nil, errors.New("gost: failed to parse EC private key embedded in PKCS#8: " + err.Error())
    }

    return key, nil
}

func marshalGostPrivateKeyWithOID(key *PrivateKey, oid asn1.ObjectIdentifier) ([]byte, error) {
    if !key.Curve.IsOnCurve(key.X, key.Y) {
        return nil, errors.New("invalid gost key public key")
    }

    privateKey := ToPrivateKey(key)
    publicKey  := ToPublicKey(&key.PublicKey)

    return asn1.Marshal(gostPrivateKey{
        Version:       gostPrivKeyVersion,
        PrivateKey:    privateKey,
        NamedCurveOID: oid,
        PublicKey:     asn1.BitString{
            Bytes: publicKey,
        },
    })
}

func parseGostPrivateKey(namedCurveOID *asn1.ObjectIdentifier, der []byte) (key *PrivateKey, err error) {
    var privKey gostPrivateKey
    if _, err := asn1.Unmarshal(der, &privKey); err != nil {
        return nil, errors.New("gost: failed to parse EC private key: " + err.Error())
    }

    if privKey.Version != gostPrivKeyVersion {
        return nil, fmt.Errorf("gost: unknown EC private key version %d", privKey.Version)
    }

    var curve *Curve
    if namedCurveOID != nil {
        curve = NamedCurveFromOid(*namedCurveOID)
    } else {
        curve = NamedCurveFromOid(privKey.NamedCurveOID)
    }

    if curve == nil {
        return nil, errors.New("gost: unknown gost curve")
    }

    priv, err := NewPrivateKey(curve, privKey.PrivateKey)
    if err != nil {
        return nil, err
    }

    return priv, nil
}
