package main

import (
	"crypto/md5"
	"crypto/rc4"
	"fmt"
)

func main() {
	key := "123456"
	md5sum := md5.Sum([]byte(key))

	cipher_encode, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)
	}
	data := []byte("中国人民解放军")
	fmt.Println("原始数据：", data)

	cipher_encode.XORKeyStream(data, data)
	fmt.Println("加密后的数据:", data)

	cipher_decode, err := rc4.NewCipher(md5sum[:])

	cipher_decode.XORKeyStream(data, data)
	fmt.Println("解密后的数据：", string(data))
}
