package ghash

import (
	"hash/fnv"
	"strconv"
)

func Fnv(input string) string {
	hash := fnv.New64()
	// We purposely ignore the error because the implementation always returns nil.
	_, _ = hash.Write([]byte(input))

	return strconv.FormatUint(hash.Sum64(), 16)
}
