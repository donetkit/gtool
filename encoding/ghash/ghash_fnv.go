package ghash

import (
	"hash/fnv"
	"strconv"
)

func Fnv(input []byte) string {
	hash := fnv.New64()
	// We purposely ignore the error because the implementation always returns nil.
	_, _ = hash.Write(input)

	return strconv.FormatUint(hash.Sum64(), 16)
}
