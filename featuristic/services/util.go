package services

import (
	"fmt"
	"hash/fnv"
)

func getCacheKey(namespace, key string) string {
	return fmt.Sprintf("%s_%s", namespace, key)
}

func hashUserID(userID string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(userID))
	if err != nil {
		return 0
	}
	return int(h.Sum32())
}
