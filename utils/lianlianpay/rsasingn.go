package lianlianpay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
)

func Sign(signStr string) (string, error) {

	priKeyValue := `MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDU1ZsCGLVjSU6s
xgpIya1Odg7sIMA/lDOpZuiV1xHv22X/8gn6xvQcLKaUIwPkJTON1IpgwejQUBYw
A6z3sN47KN/Hi1RmXrYkQZ0WaiDgJQkIver+wsABIDC0Czc0SsmfJgwvtUCgn07X
arQsMyYMHDFg73JqTGLYGAcUtbe+GCNN04ZrIYstf+GwzQDk9DdVd3QLofezK4KY
N9QcZlgypEY+ks9/vNvpVlfleEqwu5lVzoKdSjCsKy1L3vaz7IE5SP7RqDB7IjDa
mLkBCVvrgVWe7TrOqUS9iQuSD/hnE2LBRPINOYSnPLv0D3YNwtWfWXg4MSDS6AYh
nRgp4MYvAgMBAAECggEARJR9yaunixIckUyRFMozl2QwQ8L736DEEtJvoJS4GWdD
UzktIQsJrqhulq4/uzyHg6yYWDwzZihgNeGBUyDBQyPVE8ngDMp4+27PDOFurXsR
kakMt0GOt+4k82bnx2hP1oHejGLlq7wZng0qeFGU8XMGMDuHzRvG/+WiJni1d8Fd
WN4W0SISQFIb2VUjgSU+o8z9bl/y+Fez58PcH4NinyodJxoKp9yYDT8NvWPXoD6Y
sCKCNLg6HDHlNGiBF3yEYW8kEWSaSeuCfc6V0Fh0oXlSc0GYjz8oj3N3T/0K0MaI
68+ymoFomYvPaL+TRgEpwFYohG+NAQjoRsyUx1zApQKBgQD/ZkZgKer1l7jo9z9L
SEgqN7Aht9SV4Y+ZkP7qO+jA08wKJ+FKAazX41YAU4h9yLfYL7acyUH3AkH5Ylwo
2tPcmn5tHal4ur9AmDgYOdxN36wxHQrcHvj395pmj6s2avRVB/55JKFk6FO/pfET
Tp2HvAPIApffTL5hR7iH452TOwKBgQDVVbXr9dYBLMQ5R1P7vPITAeONa7vcVr5w
SXtDHovbppPb8Po/hxpgJTdDjmbnvR2xE9J3nMiy1iojkr2WOS9jAc2AuAeeohOE
n+DldCN/FzwI1JWa/gdUqeeB15QCFebkn67UxYLagf6xE82KHrvKrCjavyEXSBug
W3F20trBnQKBgQC3fNcax4LSpwpl9Rw4Ddoq2o9j6hqFA+STQ6SgzCHQR1nUAgrN
jJ6NC/sBiQIKvGW4n67mvYDy4WBcttnXUD/dwfEKm8Dhx+kXk0TZwtgP6p5fOqpO
ssHUBMOc46LyGa0oWfHtPu/k+zUP/zxPzq9GkUF2NMa/2vb8FPbySgkveQKBgAG1
fxCV1pIJ1sSSsrSIjfPjBvyiCU45j6r8m7Us8Hfg/Lw9UPPjVC3C/o9W+7INjMhX
XU9B/UWZe5rvWP90E9IaOJg4YS8/IyLj2Pxdks84021KrqwgLtTZy5X5sSDXJhNc
yBzCPEsd0Xl/rBzYJC76dBVUhKJYHIiLUvxdsgztAoGBAKVDlv8neQk/2QTfyLEP
tGl9/TaYKW7L8QcKW2SCl//bDvKNGliV4TAjbXALADuAFD4sMjoF1Ur4e9OndY6c
c7bv1tc+HbCAYE8whXGZpafwCflpVN9grZAnS2kNeaiM+bqFG8wDbW/zeqW5Oahh
2DelLhP4FraMXtgeO7hfMKNQ`

	// 解码 Base64 编码的私钥
	privBytes, err := base64.StdEncoding.DecodeString(priKeyValue)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %w", err)
	}

	// 解析私钥
	privKey, err := x509.ParsePKCS8PrivateKey(privBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	// 确保私钥是 RSA 类型
	rsaPrivKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("not an RSA private key")
	}

	// 使用 SHA-1 生成数据的哈希值
	hashed := sha1.Sum([]byte(signStr))

	// 使用私钥生成 PKCS1v15 签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivKey, crypto.SHA1, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %w", err)
	}

	// 将签名结果编码为 Base64 格式
	return base64.StdEncoding.EncodeToString(signature), nil
}
