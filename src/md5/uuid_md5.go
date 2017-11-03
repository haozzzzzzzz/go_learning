package main

import (
	"fmt"
	"crypto/md5"
	"github.com/satori/go.uuid"
)

func main() {
	strUuid := fmt.Sprintf("%s", uuid.NewV4())
	advanceID := fmt.Sprintf("%x", md5.Sum([]byte(strUuid)))
	fmt.Println(advanceID)
}
