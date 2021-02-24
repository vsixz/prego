package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

type HashMode uint

const (
	_ HashMode = iota
	Md5
	Sha1
	Sha256
	Sha512
)

func Hash(str string, mode HashMode) string {
	h := getHash(mode)
	h.Write([]byte(str))
	bytes := h.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}

func DoubleHash(str string, mode HashMode) string {
	h := getHash(mode)
	h.Write([]byte(str))
	bytes := h.Sum(nil)
	h.Reset()
	h.Write(bytes)
	bytes = h.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}

func getHash(mode HashMode) hash.Hash {
	var h hash.Hash
	switch mode {
	case Md5:
		h = md5.New()
	case Sha1:
		h = sha1.New()
	case Sha256:
		h = sha256.New()
	case Sha512:
		h = sha512.New()
	default:
		panic("unsupported hash mode.")
	}
	return h
}
