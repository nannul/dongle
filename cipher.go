package dongle

import (
	"bytes"
	"crypto/cipher"
)

// mode constants
// 模式常量
const (
	CBC = "cbc"
	ECB = "ecb"
	CFB = "cfb"
	OFB = "ofb"
	CTR = "ctr"
)

// padding constants
// 填充常量
const (
	No    = "no"
	Zero  = "zero"
	PKCS5 = "pkcs5"
	PKCS7 = "pkcs7"
)

// pem version constants
// 证书版本
const (
	PKCS1 = "pkcs1"
	PKCS8 = "pkcs8"
)

// Cipher defines a Cipher struct.
// 定义 Cipher 结构体
type Cipher struct {
	mode    string // 分组模式
	padding string // 填充模式
	key     []byte // 密钥
	iv      []byte // 偏移向量
}

// NewCipher returns a new Cipher instance.
// 初始化 Cipher 结构体
func NewCipher() *Cipher {
	return &Cipher{
		mode:    CBC,
		padding: PKCS7,
	}
}

// SetMode sets mode.
// 设置分组模式
func (c *Cipher) SetMode(mode string) {
	c.mode = mode
}

// SetPadding sets padding.
// 设置填充模式
func (c *Cipher) SetPadding(padding string) {
	c.padding = padding
}

// SetKey sets key.
// 设置密钥
func (c *Cipher) SetKey(key interface{}) {
	switch v := key.(type) {
	case string:
		c.key = string2bytes(v)
	case []byte:
		c.key = v
	}
}

// SetIV sets iv.
// 设置偏移向量
func (c *Cipher) SetIV(iv interface{}) {
	switch v := iv.(type) {
	case string:
		c.iv = string2bytes(v)
	case []byte:
		c.iv = v
	}
}

// ZeroPadding padding with Zero mode.
// 进行零填充
func (c *Cipher) ZeroPadding(src []byte, size int) []byte {
	padding := bytes.Repeat([]byte{byte(0)}, size-len(src)%size)
	return append(src, padding...)
}

// ZeroUnPadding removes padding with Zero mode.
// 移除零填充
func (c *Cipher) ZeroUnPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
}

// PKCS5Padding padding with PKCS5 mode.
// 进行 PKCS5 填充
func (c *Cipher) PKCS5Padding(src []byte) []byte {
	return c.PKCS7Padding(src, 8)
}

// PKCS5UnPadding removes padding with PKCS5 mode.
// 移除 PKCS5 填充
func (c *Cipher) PKCS5UnPadding(src []byte) []byte {
	return c.PKCS7UnPadding(src)
}

// PKCS7Padding padding with PKCS7 mode.
// 进行 PKCS7 填充
func (c *Cipher) PKCS7Padding(src []byte, size int) []byte {
	paddingCount := size - len(src)%size
	paddingText := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	return append(src, paddingText...)
}

// PKCS7UnPadding removes padding with PKCS7 mode.
// 移除 PKCS7 填充
func (c *Cipher) PKCS7UnPadding(src []byte) []byte {
	trim := len(src) - int(src[len(src)-1])
	return src[:trim]
}

// CBCEncrypt encrypts with CBC mode.
// CBC 加密
func (c *Cipher) CBCEncrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCEncrypter(block, c.iv[:block.BlockSize()]).CryptBlocks(dst, src)
	return
}

// CBCDecrypt decrypts with CBC mode.
// CBC 解密
func (c *Cipher) CBCDecrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCDecrypter(block, c.iv[:block.BlockSize()]).CryptBlocks(dst, src)
	return
}

// CFBEncrypt encrypts with CFB mode.
// CFB 加密
func (c *Cipher) CFBEncrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBEncrypter(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// CFBDecrypt decrypts with CFB mode.
// CFB 解密
func (c *Cipher) CFBDecrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBDecrypter(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// CTREncrypt encrypts with CTR mode.
// CTR 加密
func (c *Cipher) CTREncrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// CTRDecrypt decrypts with CTR mode.
// CTR 解密
func (c *Cipher) CTRDecrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// ECBEncrypt encrypts with ECB mode.
// ECB 加密
func (c *Cipher) ECBEncrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	encrypted, size := dst, block.BlockSize()
	for len(src) > 0 {
		block.Encrypt(encrypted, src[:size])
		src = src[size:]
		encrypted = encrypted[size:]
	}
	return
}

// ECBDecrypt decrypts with ECB mode.
// ECB 解密
func (c *Cipher) ECBDecrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	decrypted, size := dst, block.BlockSize()
	for len(src) > 0 {
		block.Decrypt(decrypted, src[:size])
		src = src[size:]
		decrypted = decrypted[size:]
	}
	return
}

// OFBEncrypt encrypts with OFB mode.
// OFB 加密
func (c *Cipher) OFBEncrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewOFB(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// OFBDecrypt decrypts with OFB mode.
// OFB 解密
func (c *Cipher) OFBDecrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewOFB(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}
