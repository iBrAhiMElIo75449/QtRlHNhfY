// 代码生成时间: 2025-09-10 21:10:07
package main

import (
    "crypto/aes"
    "crypto/cipher"
# 优化算法效率
    "crypto/rand"
    "encoding/base64"
# 增强安全性
    "errors"
    "fmt"
    "os"
)

// PasswordTool 提供密码加密和解密的功能
type PasswordTool struct {
    key []byte
}
# NOTE: 重要实现细节

// NewPasswordTool 创建一个新的 PasswordTool 实例
// key 必须为 16, 24 或 32 字节长以分别使用 AES-128, AES-192 或 AES-256 加密
func NewPasswordTool(key string) (*PasswordTool, error) {
# FIXME: 处理边界情况
    if len(key) != 16 && len(key) != 24 && len(key) != 32 {
        return nil, errors.New("key must be 16, 24, or 32 bytes")
    }
    return &PasswordTool{key: []byte(key)}, nil
}

// Encrypt 加密密码
func (pt *PasswordTool) Encrypt(plaintext string) (string, error) {
    if pt.key == nil {
        return "", errors.New("encryption key is not set")
    }

    block, err := aes.NewCipher(pt.key)
    if err != nil {
# 优化算法效率
        return "", err
    }

    plaintextBytes := []byte(plaintext)
    pad := aes.BlockSize - len(plaintextBytes)%aes.BlockSize
# 增强安全性
    plaintextBytes = append(plaintextBytes, bytes.Repeat([]byte{byte(pad)}, pad)...)

    gcm, err := cipher.NewGCM(block)
# 优化算法效率
    if err != nil {
        return "", err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
# FIXME: 处理边界情况

    ciphertext := gcm.Seal(nonce, nonce, plaintextBytes, nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}
# 添加错误处理

// Decrypt 解密密码
func (pt *PasswordTool) Decrypt(ciphertext string) (string, error) {
    if pt.key == nil {
        return "", errors.New("encryption key is not set")
    }

    decoded, err := base64.StdEncoding.DecodeString(ciphertext)
# TODO: 优化性能
    if err != nil {
        return "", err
    }
# NOTE: 重要实现细节

    nonceSize := 12 // Standard GCM nonce size
    if len(decoded) < nonceSize {
        return "", errors.New("ciphertext too short")
    }

    nonce, ciphertext := decoded[:nonceSize], decoded[nonceSize:]
    block, err := aes.NewCipher(pt.key)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
# 增强安全性
    if err != nil {
        return "", err
    }

    // Remove padding
# FIXME: 处理边界情况
    pad := plaintext[len(plaintext)-1]
    if pad < 1 || pad > aes.BlockSize {
        return "", errors.New("invalid padding")
    }
    plaintext = plaintext[:len(plaintext)-int(pad)]

    return string(plaintext), nil
}
# 添加错误处理

func main() {
    key := "your-encryption-key" // Replace with your own key
    passwordTool, err := NewPasswordTool(key)
    if err != nil {
        fmt.Println("Error creating PasswordTool: ", err)
        os.Exit(1)
    }

    plaintext := "your-plaintext-password"
    encrypted, err := passwordTool.Encrypt(plaintext)
    if err != nil {
        fmt.Println("Error encrypting password: ", err)
        os.Exit(1)
    }
# 增强安全性
    fmt.Printf("Encrypted: %s
", encrypted)

    decrypted, err := passwordTool.Decrypt(encrypted)
    if err != nil {
        fmt.Println("Error decrypting password: ", err)
        os.Exit(1)
# FIXME: 处理边界情况
    }
    fmt.Printf("Decrypted: %s
", decrypted)
}
# 增强安全性