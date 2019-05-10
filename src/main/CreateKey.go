package main

import (
"crypto/ecdsa"
"crypto/elliptic"
"crypto/sha256"
"crypto/x509"
"encoding/hex"
"errors"
"fmt"
"golang.org/x/crypto/ripemd160"
"math/big"
"math/rand"
"strings"
"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"


type Gkey struct {
	privateKey *ecdsa.PrivateKey
	publicKey  ecdsa.PublicKey
}
//生成随机字符串
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		b[i] = letterBytes[r.Intn(n)]
	}
	return string(b)
}
//生成公私钥
func MakeNewKey(randKey string) (*ecdsa.PrivateKey, error) {
	var curve elliptic.Curve
	var prk *ecdsa.PrivateKey
	lenth := len(randKey)
	if lenth < 224/8+8 {
		err := errors.New("RandKey is too short,It must be longer than 36 bits")
		fmt.Println(err)
	} else if lenth > 521/8+8 {
		curve = elliptic.P521()
	} else if lenth > 384/8+8 {
		curve = elliptic.P384()
	} else if lenth > 256/8+8 {
		curve = elliptic.P256()
	} else if lenth > 224/8+8 {
		curve = elliptic.P224()
	}
	prk, err := ecdsa.GenerateKey(curve, strings.NewReader(randKey))
	if err != nil {
		fmt.Println(err)
	}
	return prk, err
}
//将密钥转字符串
func StrKey(key *ecdsa.PrivateKey) (string, string) {
	x509PrivateKeyByte, _ := x509.MarshalECPrivateKey(key)
	privateKeyStr := hex.EncodeToString(x509PrivateKeyByte)
	publicKey := key.PublicKey
	x509PublicKeyByte, _ := x509.MarshalPKIXPublicKey(&publicKey)
	publicKeyStr := hex.EncodeToString(x509PublicKeyByte)
	return privateKeyStr, publicKeyStr
}
//将字符串转密钥
func UnStrKey(privateKeyStr string) *ecdsa.PrivateKey {
	fx509PrivateKeyByte, _ := hex.DecodeString(privateKeyStr)
	fPrivateKey, _ := x509.ParseECPrivateKey(fx509PrivateKeyByte)
	return fPrivateKey
}
//获取地址
func GetAddress(k ecdsa.PublicKey) (address string) {
	publicKey := k
	sha256_h := sha256.New()
	sha256_h.Reset()
	publicKeyByte, _ := x509.MarshalPKIXPublicKey(&publicKey)
	sha256_h.Write(publicKeyByte)
	pub_hash_1 := sha256_h.Sum(nil) // 对公钥进行hash256运算
	ripemd160_h := ripemd160.New()
	ripemd160_h.Reset()
	ripemd160_h.Write(pub_hash_1)
	pub_hash_2 := ripemd160_h.Sum(nil) // 对公钥hash进行ripemd160运算
	address = b58checkencode(0x00, pub_hash_2)
	return address
}
func b58checkencode(ver uint8, b []byte) (s string) {
	bcpy := append([]byte{ver}, b...)
	sha256H := sha256.New()
	sha256H.Reset()
	sha256H.Write(bcpy)
	hash1 := sha256H.Sum(nil)
	sha256H.Reset()
	sha256H.Write(hash1)
	hash2 := sha256H.Sum(nil)
	bcpy = append(bcpy, hash2[0:4]...)
	s = b58encode(bcpy)
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}
	return s
}
func b58encode(b []byte) (s string) {
	/* See https://en.bitcoin.it/wiki/Base58Check_encoding */
	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	x := new(big.Int).SetBytes(b)
	// Initialize
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	/* Convert big int to string */
	for x.Cmp(zero) > 0 {
		/* x, r = (x /58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
	}
	return s
}
