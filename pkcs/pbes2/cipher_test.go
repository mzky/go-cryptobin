package pbes2

import (
    "testing"

    cryptobin_test "github.com/deatil/go-cryptobin/tool/test"
)

func test_cipher(t *testing.T, cipher Cipher, name string, key []byte) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    plaintext := []byte("test data")

    endata, parm, err := cipher.Encrypt(key, plaintext)
    assertError(err, name + "-Encrypt")
    assertNotEmpty(endata, name + "-endata")
    assertNotEmpty(parm, name + "-parm")

    dedata, err := cipher.Decrypt(key, parm, endata)
    assertError(err, name + "-Decrypt")
    assertNotEmpty(dedata, name + "-dedata")

    assertEqual(dedata, plaintext, name + "-Equal")
}

func Test_Ciphers(t *testing.T) {
    test_cipher(t, DESCBC, "DESCBC", []byte("ssdfrt5t"))
    test_cipher(t, DESEDE3CBC, "DESEDE3CBC", []byte("ssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, RC2CBC, "RC2CBC", []byte("ssdfrt5tssdfrt5t"))
    test_cipher(t, RC5CBC, "RC5CBC", []byte("ssdfrt5tssdfrt5t"))

    test_cipher(t, AES128CBC, "AES128CBC", []byte("ssdfrt5tssdfrt5t"))
    test_cipher(t, AES128OFB, "AES128OFB", []byte("ssdfrt5tssdfrt5t"))
    test_cipher(t, AES128CFB, "AES128CFB", []byte("ssdfrt5tssdfrt5t"))
    test_cipher(t, AES128GCM, "AES128GCM", []byte("ssdfrt5tssdfrt5t"))
    test_cipher(t, AES128GCMb, "AES128GCMb", []byte("ssdfrt5tssdfrt5t"))
    test_cipher(t, AES128CCM, "AES128CCM", []byte("ssdfrt5tssdfrt5t"))
    test_cipher(t, AES128CCMb, "AES128CCMb", []byte("ssdfrt5tssdfrt5t"))

    test_cipher(t, AES192CBC, "AES192CBC", []byte("ssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES192OFB, "AES192OFB", []byte("ssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES192CFB, "AES192CFB", []byte("ssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES192GCM, "AES192GCM", []byte("ssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES192GCMb, "AES192GCMb", []byte("ssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES192CCM, "AES192CCM", []byte("ssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES192CCMb, "AES192CCMb", []byte("ssdfrt5tssdfrt5tssdfrt5t"))

    test_cipher(t, AES256CBC, "AES256CBC", []byte("ghiolkjmssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES256OFB, "AES256OFB", []byte("ghiolkjmssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES256CFB, "AES256CFB", []byte("ghiolkjmssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES256GCM, "AES256GCM", []byte("ghiolkjmssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES256GCMb, "AES256GCMb", []byte("ghiolkjmssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES256CCM, "AES256CCM", []byte("ghiolkjmssdfrt5tssdfrt5tssdfrt5t"))
    test_cipher(t, AES256CCMb, "AES256CCMb", []byte("ghiolkjmssdfrt5tssdfrt5tssdfrt5t"))

    test_cipher(t, SM4CBC, "SM4CBC", []byte("ghdfrt5tssdfrt5t"))
    test_cipher(t, SM4OFB, "SM4OFB", []byte("ghdfrt5tssdfrt5t"))
    test_cipher(t, SM4CFB, "SM4CFB", []byte("ghdfrt5tssdfrt5t"))
    test_cipher(t, SM4CFB1, "SM4CFB1", []byte("ghdfrt5tssdfrt5t"))
    test_cipher(t, SM4CFB8, "SM4CFB8", []byte("ghdfrt5tssdfrt5t"))
    test_cipher(t, SM4GCM, "SM4GCM", []byte("ghdfrt5tssdfrt5t"))
    test_cipher(t, SM4GCMb, "SM4GCMb", []byte("ghdfrt5tssdfrt5t"))
    test_cipher(t, SM4CCM, "SM4CCM", []byte("ghdfrt5tssdfrt5t"))
    test_cipher(t, SM4CCMb, "SM4CCMb", []byte("ghdfrt5tssdfrt5t"))

}

