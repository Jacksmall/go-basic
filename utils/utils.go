package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

type fnv64a struct{}

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

func (f *fnv64a) Sum64(s string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(s); i++ {
		hash ^= uint64(s[i])
		hash *= prime64
	}
	return hash
}

func main() {
	// fmt.Println(isPowerOfTwo(1024))
	// cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	// cache.Set("MyKey1", []byte("abc"))

	// res, err := cache.Get("MyKey1")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(res))

	f := &fnv64a{}
	fmt.Printf("MyKey1 hashKey is:%d\n", f.Sum64("MyKey1"))
	fmt.Println(f.Sum64("MyKey1") & uint64(1023))

	fmt.Println(wrapEntry(uint64(time.Now().Unix()), f.Sum64("MyKey1"), "MyKey1", []byte("1234"), &[]byte{}))
}

// func isPowerOfTwo(number int) bool {
// 	return (number != 0) && (number&(number-1) == 0)
// }

const (
	timestampInBytesLength   = 8
	hashKeyInBytesLength     = 8
	keyInBytesLength         = 2
	headersSizeInBytesLength = timestampInBytesLength + hashKeyInBytesLength + keyInBytesLength
)

func wrapEntry(timestamp uint64, hash uint64, key string, entry []byte, buffer *[]byte) []byte {
	keyLength := len(key)
	blobLength := len(entry) + headersSizeInBytesLength + keyLength

	if blobLength > len(*buffer) {
		*buffer = make([]byte, blobLength)
	}
	blob := *buffer
	// 将所有东西(包括整型的数字)encoding放进blob字节切片中
	binary.LittleEndian.PutUint64(blob, timestamp)
	binary.LittleEndian.PutUint64(blob[:timestampInBytesLength], hash)
	binary.LittleEndian.PutUint16(blob[:timestampInBytesLength+hashKeyInBytesLength], uint16(keyLength))

	copy(blob[headersSizeInBytesLength:], key)
	copy(blob[headersSizeInBytesLength+keyLength:], entry)

	return blob[:blobLength]
}
