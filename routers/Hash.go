package routers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"net/http"
	"strconv"

	"github.com/JavDomGom/gcat/models"
)

func Hash(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")

	var s = models.Hash{
		Function: r.URL.Path[1:],
		Msg:      msg,
	}

	switch r.URL.Path {
	case "/md5":
		sum := md5.Sum([]byte(msg))
		s.Sum = hex.EncodeToString(sum[:])
	case "/sha1":
		sum := sha1.Sum([]byte(msg))
		s.Sum = hex.EncodeToString(sum[:])
	case "/sha256":
		sum := sha256.Sum256([]byte(msg))
		s.Sum = hex.EncodeToString(sum[:])
	case "/sha512":
		sum := sha512.Sum512([]byte(msg))
		s.Sum = hex.EncodeToString(sum[:])
	case "/fnv":
		sum := fnv.New32().Sum([]byte(msg))
		s.Sum = hex.EncodeToString(sum[:])
	case "/adler32":
		sum := adler32.Checksum([]byte(msg))
		s.Sum = strconv.FormatUint(uint64(sum), 16)
	case "/crc32":
		crc32Table := crc32.MakeTable(0xD5828281)
		sum := crc32.Checksum([]byte(msg), crc32Table)
		s.Sum = strconv.FormatUint(uint64(sum), 16)
	case "/crc64":
		crc64Table := crc64.MakeTable(0xC96C5795D7870F42)
		sum := crc64.Checksum([]byte(msg), crc64Table)
		s.Sum = strconv.FormatUint(sum, 16)
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}
