package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func SHA2(str string) string {
	w := sha256.New()
	io.WriteString(w, str)
	bw := w.Sum(nil)
	shastr2 := hex.EncodeToString(bw)
	return shastr2
}
