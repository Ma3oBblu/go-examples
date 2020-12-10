package internal

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"hash/fnv"
	"strconv"
	"strings"
)

// GenerateMD5Hash генерация md5 ключа
func GenerateMD5Hash(ids []string) string {
	md5Key := md5.Sum([]byte(strings.Join(ids, "|")))
	return hex.EncodeToString(md5Key[:])
}

// GenerateFnvHashWithWriter генерация fnv ключа через writer
func GenerateFnvHashWithWriter(ids []string) string {
	h := fnv.New64a()
	_, _ = h.Write([]byte(strings.Join(ids, "|")))
	return strconv.FormatUint(h.Sum64(), 10)
}

// GenerateFnvHashWithSumAndConvert генерация fnv ключа через конвертацию в int->string
func GenerateFnvHashWithSumAndConvert(ids []string) string {
	h := fnv.New64a().Sum([]byte(strings.Join(ids, "|")))
	fnvKey2 := binary.BigEndian.Uint64(h)
	return strconv.Itoa(int(fnvKey2))
}
