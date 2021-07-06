package main

import (
	"fmt"
	"log"

	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

var sentData string
var readData string

// Component for decorator pattern
type Component interface {
	Operator(string)
}

// SendData origin
type SendData struct{}

// Operator origin
func (s *SendData) Operator(data string) {
	sentData = data
}

// ZipComponent Decorator
type ZipComponent struct {
	com Component
}

// Operator Decorator
func (z *ZipComponent) Operator(data string) {
	zipdate, err := lzw.Write([]byte(data))
	if err != nil {
		log.Fatal("could not zip data", err)
		return
	}

	z.com.Operator(string(zipdate))
}

// EncryptComponent Decorator
type EncryptComponent struct {
	key string
	com Component
}

// Operator Decorator
func (e *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), e.key)
	if err != nil {
		log.Fatal("could not encrypt data", err)
		return
	}
	e.com.Operator(string(encryptData))
}

// ReadData origin
type ReadData struct{}

// Operator origin
func (r *ReadData) Operator(data string) {
	readData = data
}

// UnzipComponent Decorator
type UnzipComponent struct {
	com Component
}

// Operator Decorator
func (u *UnzipComponent) Operator(data string) {
	unZipData, err := lzw.Read([]byte(data))
	if err != nil {
		log.Fatal("could not unzip data", err)
		return
	}

	u.com.Operator(string(unZipData))
}

// DecryptComponent Decorator
type DecryptComponent struct {
	key string
	com Component
}

// Operator Decorator
func (d *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), d.key)
	if err != nil {
		log.Fatal("could not decrypt data", err)
		return
	}

	d.com.Operator(string(decryptData))
}

func main() {
	sender := &EncryptComponent{
		key: "abcd",
		com: &ZipComponent{
			com: &SendData{},
		},
	}

	sender.Operator("Hello Decorator")
	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcd",
			com: &ReadData{},
		},
	}

	receiver.Operator(sentData)
	fmt.Println(readData)
}
