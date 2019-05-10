package main

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

/**
  对text加密，text必须是一个hash值，例如md5、sha1等
  使用私钥prk
  使用随机熵增强加密安全，安全依赖于此熵，randsign
  返回加密结果，结果为数字证书r、s的序列化后拼接，然后用hex转换为string
*/
func sign(text []byte,randSign string,prk *ecdsa.PrivateKey) (string, error) {
	r, s, err := ecdsa.Sign(strings.NewReader(randSign), prk, text)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

/**
  证书分解
  通过hex解码，分割成数字证书r，s
*/
func getSign( signature string) (rint, sint big.Int, err error)  {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrypt error, "+ err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err !=  nil {
		err = errors.New("decode error,"+err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		fmt.Println("decode = ",err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]),"+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, "+ err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, "+ err.Error())
		return
	}
	return

}
/**
  校验文本内容是否与签名一致
  使用公钥校验签名和文本内容
*/
func Verify(text string, signature string, key ecdsa.PublicKey) (bool, error) {
	salt := "131ilzaw"
	htext := hashtext(text,salt)
	rint, sint, err :=getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(&key,htext,&rint,&sint)
	return result,nil
}

/**
  hash加密
  使用md5加密
*/
func hashtext(text, salt string) ([]byte) {

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(text))
	result := Md5Inst.Sum([]byte(salt))

	return result
}


func ToSign(text string,key *ecdsa.PrivateKey)(string) {
	//随机熵，用于加密安全
	randSign := "20180619zafes20180619zafes20180619zafessss"//至少36位
	//hash加密使用md5用到的salt
	salt := "131ilzaw"
	//待加密的明文
	//text1 := "hlloaefaefaefaefaefaefaefhelloaefaefaefaefaefaefaefhelloaefaefaefaefaefaefaef1"

	//hash取值
	htext := hashtext(text,salt)
	//htext1 := hashtext(text1,salt)
	//hash值编码输出
	//fmt.Println(hex.EncodeToString(htext))

	//hash值进行签名
	result, err := sign(htext,randSign,key)
	if err != nil {
		fmt.Println(err)
	}
	//签名输出
	fmt.Println("数据签名：",result)

	//签名与hash值进行校验
	//tmp, err := Verify(text,result,key.PublicKey)
	//fmt.Println(tmp)
	return result
}

