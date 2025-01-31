package gost

import (
    "crypto/cipher"

    "github.com/deatil/go-cryptobin/tool/alias"
)

// GOST 28147-89 defines a block size of 64 bits
const BlockSize = 8

// Internal state of the GOST block cipher
type gostCipher struct {
    key []uint32 // Encryption key
    s   [][]byte // S-box provided as parameter
    k   [][]byte // Expanded s-box
}

// NewCipher creates and returns a new cipher.Block. The key argument
// should be the 32 byte GOST 28147-89 key. The sbox argument should be a
// 64 byte substitution table, represented as a two-dimensional array of 8 rows
// of 16 4-bit integers.
func NewCipher(key []byte, sbox [][]byte) (cipher.Block, error) {
    if len(key) != 32 {
        return nil, KeySizeError(len(key))
    }

    if len(sbox) != 8 {
        return nil, SboxSizeError(len(sbox))
    }

    for i := 0; i < len(sbox); i++ {
        if len(sbox[i]) != 16 {
            return nil, SboxSizeError(len(sbox[i]))
        }
    }

    newKey := bytesToUint32s(key)
    kbox := sboxExpansion(sbox)

    c := &gostCipher{
        key: newKey,
        s:   sbox,
        k:   kbox,
    }

    return c, nil
}

func (this *gostCipher) BlockSize() int {
    return BlockSize
}

func (this *gostCipher) Encrypt(dst, src []byte) {
    if len(src) < BlockSize {
        panic("cryptobin/gost: input not full block")
    }

    if len(dst) < BlockSize {
        panic("cryptobin/gost: output not full block")
    }

    if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
        panic("cryptobin/gost: invalid buffer overlap")
    }

    encSrc := bytesToUint32s(src)
    encDst := make([]uint32, len(encSrc))

    this.encrypt32(encDst, encSrc)

    resBytes := uint32sToBytes(encDst)
    copy(dst, resBytes)
}

func (this *gostCipher) Decrypt(dst, src []byte) {
    if len(src) < BlockSize {
        panic("cryptobin/gost: input not full block")
    }

    if len(dst) < BlockSize {
        panic("cryptobin/gost: output not full block")
    }

    if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
        panic("cryptobin/gost: invalid buffer overlap")
    }

    encSrc := bytesToUint32s(src)
    encDst := make([]uint32, len(encSrc))

    this.decrypt32(encDst, encSrc)

    resBytes := uint32sToBytes(encDst)
    copy(dst, resBytes)
}

// GOST block cipher round function
func (this *gostCipher) f(x uint32) uint32 {
    x = uint32(this.k[0][(x >> 24) & 255]) << 24 |
        uint32(this.k[1][(x >> 16) & 255]) << 16 |
        uint32(this.k[2][(x >>  8) & 255]) <<  8 |
        uint32(this.k[3][x & 255])

    // rotate result left by 11 bits
    return (x << 11) | (x >> (32 - 11))
}

// Encrypt one block from src into dst.
func (this *gostCipher) encrypt32(dst, src []uint32) {
    n1, n2 := src[0], src[1]

    n2 = n2 ^ this.f(n1 + this.key[0])
    n1 = n1 ^ this.f(n2 + this.key[1])
    n2 = n2 ^ this.f(n1 + this.key[2])
    n1 = n1 ^ this.f(n2 + this.key[3])
    n2 = n2 ^ this.f(n1 + this.key[4])
    n1 = n1 ^ this.f(n2 + this.key[5])
    n2 = n2 ^ this.f(n1 + this.key[6])
    n1 = n1 ^ this.f(n2 + this.key[7])

    n2 = n2 ^ this.f(n1 + this.key[0])
    n1 = n1 ^ this.f(n2 + this.key[1])
    n2 = n2 ^ this.f(n1 + this.key[2])
    n1 = n1 ^ this.f(n2 + this.key[3])
    n2 = n2 ^ this.f(n1 + this.key[4])
    n1 = n1 ^ this.f(n2 + this.key[5])
    n2 = n2 ^ this.f(n1 + this.key[6])
    n1 = n1 ^ this.f(n2 + this.key[7])

    n2 = n2 ^ this.f(n1 + this.key[0])
    n1 = n1 ^ this.f(n2 + this.key[1])
    n2 = n2 ^ this.f(n1 + this.key[2])
    n1 = n1 ^ this.f(n2 + this.key[3])
    n2 = n2 ^ this.f(n1 + this.key[4])
    n1 = n1 ^ this.f(n2 + this.key[5])
    n2 = n2 ^ this.f(n1 + this.key[6])
    n1 = n1 ^ this.f(n2 + this.key[7])

    n2 = n2 ^ this.f(n1 + this.key[7])
    n1 = n1 ^ this.f(n2 + this.key[6])
    n2 = n2 ^ this.f(n1 + this.key[5])
    n1 = n1 ^ this.f(n2 + this.key[4])
    n2 = n2 ^ this.f(n1 + this.key[3])
    n1 = n1 ^ this.f(n2 + this.key[2])
    n2 = n2 ^ this.f(n1 + this.key[1])
    n1 = n1 ^ this.f(n2 + this.key[0])

    dst[0], dst[1] = n2, n1
}

// Decrypt one block from src into dst.
func (this *gostCipher) decrypt32(dst, src []uint32) {
    n1, n2 := src[0], src[1]

    n2 = n2 ^ this.f(n1 + this.key[0])
    n1 = n1 ^ this.f(n2 + this.key[1])
    n2 = n2 ^ this.f(n1 + this.key[2])
    n1 = n1 ^ this.f(n2 + this.key[3])
    n2 = n2 ^ this.f(n1 + this.key[4])
    n1 = n1 ^ this.f(n2 + this.key[5])
    n2 = n2 ^ this.f(n1 + this.key[6])
    n1 = n1 ^ this.f(n2 + this.key[7])

    n2 = n2 ^ this.f(n1 + this.key[7])
    n1 = n1 ^ this.f(n2 + this.key[6])
    n2 = n2 ^ this.f(n1 + this.key[5])
    n1 = n1 ^ this.f(n2 + this.key[4])
    n2 = n2 ^ this.f(n1 + this.key[3])
    n1 = n1 ^ this.f(n2 + this.key[2])
    n2 = n2 ^ this.f(n1 + this.key[1])
    n1 = n1 ^ this.f(n2 + this.key[0])

    n2 = n2 ^ this.f(n1 + this.key[7])
    n1 = n1 ^ this.f(n2 + this.key[6])
    n2 = n2 ^ this.f(n1 + this.key[5])
    n1 = n1 ^ this.f(n2 + this.key[4])
    n2 = n2 ^ this.f(n1 + this.key[3])
    n1 = n1 ^ this.f(n2 + this.key[2])
    n2 = n2 ^ this.f(n1 + this.key[1])
    n1 = n1 ^ this.f(n2 + this.key[0])

    n2 = n2 ^ this.f(n1 + this.key[7])
    n1 = n1 ^ this.f(n2 + this.key[6])
    n2 = n2 ^ this.f(n1 + this.key[5])
    n1 = n1 ^ this.f(n2 + this.key[4])
    n2 = n2 ^ this.f(n1 + this.key[3])
    n1 = n1 ^ this.f(n2 + this.key[2])
    n2 = n2 ^ this.f(n1 + this.key[1])
    n1 = n1 ^ this.f(n2 + this.key[0])

    dst[0], dst[1] = n2, n1
}
