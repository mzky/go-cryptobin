package rsa

import (
    "testing"
    "crypto/rand"

    cryptobin_test "github.com/deatil/go-cryptobin/tool/test"
)

// Test_PrimeKeyGeneration
func Test_PrimeKeyGeneration(t *testing.T) {
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)
    assertError := cryptobin_test.AssertErrorT(t)

    size := 768
    if testing.Short() {
        size = 256
    }

    obj := NewRSA().GenerateMultiPrimeKey(3, size)

    objPriKey := obj.CreatePKCS1PrivateKey()

    assertError(objPriKey.Error(), "objPriKey")
    assertNotEmpty(objPriKey.ToKeyString(), "objPriKey")

    objPubKey := obj.CreatePKCS1PublicKey()

    assertError(objPubKey.Error(), "objPubKey")
    assertNotEmpty(objPubKey.ToKeyString(), "objPubKey")
}

var (
    testSignPrikeyPkcs1En = `
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,c47d6aef98e01776d8519936adb449ad

2YSwbFUmWFWYlMLtkCWI2B8k8WxcmSyuu4vhLcHZrsIYm0EIe3xX7AkfXodZgO1g
vjjbX4w/qaB3WPg4Re9TxgaLqEEY0DRkuP4m/kDsu40ZQyMlGrXKl3Yp8D6Te2Wf
RcRBgibWwNJVgZYQ6nY/82SA2gIcvCQVo0VZVQQqnlULyhpMt2yR2N3HprD95BI3
yhoOae81qLikjVdWpoTwjSu8uHLw2qwWSrOjBgUSHjbXJBsvwIEorAGkvtTYWKCn
tUG0Rn40gDqvrEvowYdIxriZdvndcYbijsqsP3pOTPBP2rYiFT3Bj19FLC11R8bv
eXYs7/YgsVH+l7XhJabJPJSH4Zuz/TDkcVdrzBxtLxFsFVOqs68QfJ/xuM6SLYxy
6YG2Oq8MAdG96QHYUGxnIZBNXfBGsYbVGc4fSRv8FiCgXOX2l1F9dTlbX2FluKgb
om+SrCKZkWn/jEjSdnnCvkC22JzqyY2KAcLkSVUx+ZtQCl+0YHZHrk4Vkz9GvS6f
WOtBtj1YQ8BAhMpGy+eRxTIHBHHYBYosqBQ3B/5dexoFleznCHUGHzGGXRCd+ty9
9QJrur/3QZg4T/68JDypp3nIY4CcUdNOhoL1BjT0gvzaosZhAohcZaO7mCDCUQXG
oBvbC6wbNFNzm0iLBUIdNPoRujiSKgxk+m748IzJBpAgC9s3PjmNmnJQrywu9A/s
YkVBTYwccqhVzjSYBEnGSnBWcfvTSeSf2CoaSMnju9aNSt9gIuNTe5nact4VT0+T
VCx2NqZAU5so2yhpG8g2H393/3QF81nkeuFqkSDgDvPYenrhEOs1Lbn4mRRDBAJY
-----END RSA PRIVATE KEY-----
    `

    testSignPrikeyPkcs8En = `
-----BEGIN ENCRYPTED PRIVATE KEY-----
MIIC5TBfBgkqhkiG9w0BBQ0wUjAxBgkqhkiG9w0BBQwwJAQQj5kaaZMiw2Ha3s4z
F4e0LgICJxAwDAYIKoZIhvcNAgcFADAdBglghkgBZQMEARYEEN49eNpe9aDZN0/J
U05yLWoEggKAiBUPV+9o9iMxLdFTLzvSTSJAUvXxEnbY+SfeCW25qBvCsydO7Iuy
HfL+3YYKvxH4j115RG5uWA3BK8EpfrWGlAXAacTfPZim7AUbe6Omf+rYi3KWXTQ0
5StQHiMS5t0AhWGWk1vUwL33mJK6ceKR7BhfGliI4QGyWeTeGqJkq0P7m9Fo9ylY
MP2POg/0RutacEoB2BUvHxPu+RJZoj/3K2nvLtPJo5wWLBQW4F3TdPSn79VlGaWN
gTDY21sYPVz+5kIl3RSnPWwpOCUw23ZNGaHVX4xtnVzzBkC1cJeGeaBgZjht5uf1
wuh57g/neLBzjBUHKxAs3OdDUKYejg/HFrId0f+nDCuRwoJeyEwxkt7fJizIFftH
9Ks42I5ndTcB7zs1Eb5IF/HYYWHTzNLRm+Wt57LyoG++nOv0kMmlHiX9omgtsveA
Y6ONUeTxWwa2ljP5zpslGXYhmVyRg6hC/LnoqfUFf2jV5m+epvOOSQgQkzIrxN9j
QTquQhzUaaKpTBPPvkIHSb6568PLvhHqx/ilDGvp9UAWtO6bfXGwOIZcL3p9WetE
2sasOjp/VhgUJD37a0G7pwZpAoLsNC+tEUtaFkw5nNLMFM0urGYHBsVNZt/N2ibM
gKkq5AB/Zv4foXuqxsZM3k6V1uDpesQXt+/Xnqysi9MQkk16r+N9+uctugReej8M
nPaPvK54bnJ7HA9fnKt1Vf+x/NLGLdIX4xIzqBC4xOqLQlTbStRw2XRDGespR376
JnsP82g2PQN+dH0QNbW/C/eq5mfpAxh6x8JFBaha7/33czh465yg7Qpos+AMGK/D
TYEOgSzqZlhf5ISMB5uAwOW6k4TuJAUbrA==
-----END ENCRYPTED PRIVATE KEY-----
    `

    testSignPubkey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDbvgUxbNl35YrbZDPCVYOdJSxX
j69JLx3LwRUfGj34BdMb/DJuuoDzVA+fppISMul9mNaspHxguWsHgnvSpc4Swceq
t5Rr4E4VBnudvTe7BU9aHmbY4g7aldJ+S5gC3dT11wHz6Sre+2a74xsZxs51+M0Q
tGjMoEruLNj7fkPcNQIDAQAB
-----END PUBLIC KEY-----
    `
)

func Test_RSAPkcs1Sign(t *testing.T) {
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertTrue := cryptobin_test.AssertTrueT(t)

    data := "test-pass"

    obj := NewRSA()

    sign := obj.
        FromString(data).
        FromPKCS1PrivateKeyWithPassword([]byte(testSignPrikeyPkcs1En), "123").
        SetSignHash("SHA256").
        Sign()
    signData := sign.ToBase64String()

    assertError(sign.Error(), "RSAPkcs1Sign-sign")
    assertNotEmpty(signData, "RSAPkcs1Sign-sign")

    verify := obj.
        FromBase64String(signData).
        FromPublicKey([]byte(testSignPubkey)).
        SetSignHash("SHA256").
        Verify([]byte(data))
    verifyData := verify.ToVerify()

    assertError(verify.Error(), "RSAPkcs1Sign-verify")
    assertTrue(verifyData, "RSAPkcs1Sign-verify")
}

func Test_RSAPkcs8Sign(t *testing.T) {
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertTrue := cryptobin_test.AssertTrueT(t)

    data := "test-pass22222"

    obj := NewRSA()

    sign := obj.
        FromString(data).
        FromPKCS8PrivateKeyWithPassword([]byte(testSignPrikeyPkcs8En), "123").
        SetSignHash("SHA256").
        Sign()
    signData := sign.ToBase64String()

    assertError(sign.Error(), "RSAPkcs1Sign-sign")
    assertNotEmpty(signData, "RSAPkcs1Sign-sign")

    verify := obj.
        FromBase64String(signData).
        FromPublicKey([]byte(testSignPubkey)).
        SetSignHash("SHA256").
        Verify([]byte(data))
    verifyData := verify.ToVerify()

    assertError(verify.Error(), "RSAPkcs1Sign-verify")
    assertTrue(verifyData, "RSAPkcs1Sign-verify")
}

var (
    testPubN = "CCE3A1FA0E3EADEE4FE464F7D45F5009DBF2D77FF9DD9D822F41E8AD6F47762FE46569E2EE39906CE557328CF9CFE33906D4D0494CADEE2357B90178D3200DFF96EBB21053DC65AEFA458BC62C5540E3343F2968F934EAD87DAFCA6681C78CD3936E14808A74D5C7CD1EE10C7C3400C52358DF30B9383C70FF4E853ADD5D21D5"
    testPubE = "10001"
    testPubEnd = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDM46H6Dj6t7k/kZPfUX1AJ2/LX
f/ndnYIvQeitb0d2L+RlaeLuOZBs5VcyjPnP4zkG1NBJTK3uI1e5AXjTIA3/luuy
EFPcZa76RYvGLFVA4zQ/KWj5NOrYfa/KZoHHjNOTbhSAinTVx80e4Qx8NADFI1jf
MLk4PHD/ToU63V0h1QIDAQAB
-----END PUBLIC KEY-----
`
)

func Test_PubNE(t *testing.T) {
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertEqual := cryptobin_test.AssertEqualT(t)

    en := NewRSA().
        FromPublicKeyNE(testPubN, testPubE).
        CreatePKCS8PublicKey()
    enData := en.ToKeyString()

    assertError(en.Error(), "PubNE-make")
    assertNotEmpty(enData, "PubNE-make")

    assertEqual(enData, testPubEnd, "PubNE-make")
}

var testPEMCiphers = []string{
    "DESCBC",
    "DESEDE3CBC",
    "AES128CBC",
    "AES192CBC",
    "AES256CBC",

    "DESCFB",
    "DESEDE3CFB",
    "AES128CFB",
    "AES192CFB",
    "AES256CFB",

    "DESOFB",
    "DESEDE3OFB",
    "AES128OFB",
    "AES192OFB",
    "AES256OFB",

    "DESCTR",
    "DESEDE3CTR",
    "AES128CTR",
    "AES192CTR",
    "AES256CTR",
}

func Test_CreatePKCS1PrivateKeyWithPassword(t *testing.T) {
    for _, cipher := range testPEMCiphers{
        test_CreatePKCS1PrivateKeyWithPassword(t, cipher)
    }
}

func test_CreatePKCS1PrivateKeyWithPassword(t *testing.T, cipher string) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    t.Run(cipher, func(t *testing.T) {
        pass := make([]byte, 12)
        _, err := rand.Read(pass)
        if err != nil {
            t.Fatal(err)
        }

        gen := New().GenerateKey(2048)

        prikey := gen.GetPrivateKey()

        pri := gen.
            CreatePKCS1PrivateKeyWithPassword(string(pass), cipher).
            ToKeyString()

        assertError(gen.Error(), "Test_CreatePKCS1PrivateKeyWithPassword")
        assertNotEmpty(pri, "Test_CreatePKCS1PrivateKeyWithPassword-pri")

        newPrikey := New().
            FromPKCS1PrivateKeyWithPassword([]byte(pri), string(pass)).
            GetPrivateKey()

        assertNotEmpty(newPrikey, "Test_CreatePKCS1PrivateKeyWithPassword-newPrikey")

        assertEqual(newPrikey, prikey, "Test_CreatePKCS1PrivateKeyWithPassword")
    })
}
