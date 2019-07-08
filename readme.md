# iid
Small globally unique IDs implemented in Go.

[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/robojones/iid)
[![CircleCI](https://circleci.com/gh/robojones/iid/tree/master.svg?style=svg)](https://circleci.com/gh/robojones/iid/tree/master)

## Features

- Can be sorted by creation time
- Globally unique
- Small size (can be stored in a 64 bit unsigned integer or an 11 byte string)
- base64url encoded string format: the ids can be used in URLs

## ID Format

| 4 Byte               | 4 Byte                                |
| -------------------- | ------------------------------------- |
| Timestamp in seconds | Cryptographically secure random bytes |

## Usage Example

```go
package main

import (
	"github.com/robojones/iid"
	"log"
	"reflect"
)

func main() {
	// Generate new iid
	id := iid.New()
	log.Printf("buffer format: %#v", id)
	
	// Export as base64url string
	str := id.String()
	log.Printf("base64 string: %s", str)
	
	// Import the id from the string
	parsed, err := iid.FromString(str)
	if err != nil {
		panic(err)
	}
	log.Printf("parsed iid from string is identical to the original: %t", reflect.DeepEqual(id, parsed))
	
	// Export as uint64
	i := id.Uint64()
	log.Printf("integer: %d", i)
	
	// Import the id from the integer
	parsed = iid.FromUint64(i)
	log.Printf("parsed iid from uint64 is identical to original: %t", reflect.DeepEqual(id, parsed))
}
```
