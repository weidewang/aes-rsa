package crypto

import (
	"log"
	"testing"
)

var plain = []byte(`>>>I need to validate a google id_token and one step involves to check the token signature.<<<`)

var sigData = []byte("abcdefg1234567890")

// AES 加密测试
func TestAES(t *testing.T) {
	t.SkipNow()
	key := []byte("123456789012345678901234567890AA")

	log.Println("原始明文:", string(plain))

	en, err := AESEncrypt(key, plain)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("密文:", en)

	pn, err := AESDecrypt(key, en)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("解密明文:", string(pn))

	if string(plain) != string(pn) {
		t.Fatal("解密结果与明文不符")
	}
}

// RSA 测试

func TestRSA(t *testing.T) {
	t.SkipNow()
	publicKey, privateKey, err := GenRSAKey(768)

	if err != nil {
		t.Fatal(err)
	}

	en, err := RSAEncryptPKCS1v15(publicKey, plain)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("RSA 加密密文", en)

	pn, err := RSADecryptPKCS1v15(privateKey, en)

	if err != nil {
		t.Fatal(err)
	}
	log.Println("RSA 解密明文", string(pn))

	if string(plain) != string(pn) {
		t.Fatal("解密结果与明文不符")
	}

	sd, err := RSASignPKCS1v15(privateKey, sigData)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("签名结果", sd)

	err = RSAVerifyPKCS1v15(publicKey, sigData, sd)
	if err != nil {
		log.Println("验签失败")
		t.Fatal(err)
	}

}

func TestAREncrypt(t *testing.T) {
	publicKey, privateKey, _ := GenRSAKey(1024)

	ar := NewArCrypto(publicKey, privateKey)

	en, err := ar.Encrypt(plain)

	if err != nil {
		t.Fatal(err)
	}

	pn, err := ar.Decrypt(en)
	if err != nil {
		t.Fatal(err)
	}

	if string(plain) != string(pn) {
		t.Fatal("解密结果与明文不符")
	}

	en64, err := ar.EncryptToString(plain)

	if err != nil {
		t.Fatal(err)
	}
	log.Println("加密密文:", en64)

	p64p, err := ar.DecryptString(en64)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("解密明文:", string(p64p))

}
