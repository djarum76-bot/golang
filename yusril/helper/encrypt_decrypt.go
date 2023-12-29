package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(text, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(os.Getenv("SECRET_BYTES")))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decrypt(text, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, []byte(os.Getenv("SECRET_BYTES")))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func EncryptList(datas []string) ([]string, error) {
	arrString := []string{}

	if datas[0] == "" {
		return datas, nil
	} else {
		for _, data := range datas {
			dataEncrypt, err := Encrypt(data, os.Getenv("SECRET_STRING"))
			if err != nil {
				return arrString, err
			}

			arrString = append(arrString, dataEncrypt)
		}

		return arrString, nil
	}
}

func DecryptList(datas []string) ([]string, error) {
	arrString := []string{}

	if datas[0] == "" {
		return datas, nil
	} else {
		for _, data := range datas {
			dataDecrypt, err := Decrypt(data, os.Getenv("SECRET_STRING"))
			if err != nil {
				return arrString, err
			}

			arrString = append(arrString, dataDecrypt)
		}

		return arrString, nil
	}
}
