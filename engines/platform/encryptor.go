package platform

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"hash"
)

type Encryptor interface {
	Encode(buf []byte) ([]byte, error)
	Decode(buf []byte) ([]byte, error)
	Sum(buf []byte) []byte
	Equal(plain, code []byte) bool
}

type AesHmacEncryptor struct {
	Cipher cipher.Block
	Hash   hash.Hash
}

func (p *AesHmacEncryptor) Encode(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(p.Cipher, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

func (p *AesHmacEncryptor) Decode(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(p.Cipher, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil

}

func (p *AesHmacEncryptor) Sum(buf []byte) []byte {
	return p.Hash.Sum(buf)
}
func (p *AesHmacEncryptor) Equal(plain, code []byte) bool {
	return hmac.Equal(p.Hash.Sum(plain), code)
}

func NewAesHmacEncryptor(ck, hk []byte) (Encryptor, error) {
	cip, err := aes.NewCipher(ck)
	if err != nil {
		return nil, err
	}
	return &AesHmacEncryptor{
		Cipher: cip,
		Hash:   hmac.New(sha512.New, hk),
	}, nil
}
