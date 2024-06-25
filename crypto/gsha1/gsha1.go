// Package gsha1 provides useful API for SHA1 encryption algorithms.
package gsha1

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"github.com/xrc360/golang/errors/gerror"
	"github.com/xrc360/golang/util/gconv"
)

// Encrypt encrypts any type of variable using SHA1 algorithms.
// It uses package gconv to convert `v` to its bytes type.
func Encrypt(v interface{}) string {
	r := sha1.Sum(gconv.Bytes(v))
	return hex.EncodeToString(r[:])
}

// EncryptFile encrypts file content of `path` using SHA1 algorithms.
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, path)
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// MustEncryptFile encrypts file content of `path` using SHA1 algorithms.
// It panics if any error occurs.
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
