package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type AesEncrypt struct {
}

//加密字符串 输入类型：string 返回类型：[]byte
func (this *AesEncrypt) Encrypt(key []byte, strMesg string) ([]byte, error) {

	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil
}

//解密字符串 输入类型：[]byte  返回类型：string
func (this *AesEncrypt) Decrypt(key []byte, src []byte) (strDesc string, err error) {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}

func main() {
	aesEnc := AesEncrypt{}
	key := "1234567890123456"

	arrEncrypt, err := aesEnc.Encrypt([]byte(key), "0123456789abcdef123")
	if err != nil {
		fmt.Println(arrEncrypt)
		return
	}
	fmt.Println(arrEncrypt)

	strMsg, err := aesEnc.Decrypt([]byte(key), arrEncrypt)
	if err != nil {
		fmt.Println(arrEncrypt)
		return
	}
	fmt.Println(strMsg)
}
