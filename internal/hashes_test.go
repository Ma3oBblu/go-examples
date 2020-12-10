package internal

import (
	"fmt"
	"testing"
)

func TestHashes(t *testing.T) {
	ids := []string{"test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test"}
	for j := 0; j < 10; j++ {
		md5Key := GenerateMD5Hash(ids)
		fmt.Println(md5Key)

		fnvKey1 := GenerateFnvHashWithWriter(ids)
		fmt.Println(fnvKey1)

		fnvKey2 := GenerateFnvHashWithSumAndConvert(ids)
		fmt.Println(fnvKey2)
	}
}
