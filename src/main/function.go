package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func GetNewsId(infoId string, thirdSource string) string {
	year, month, _ := time.Now().Date()
	return fmt.Sprintf("%d%d", year, month) + GetMd5(infoId+time.Now().Format("2006-01-02 15:04:05")+thirdSource)
}

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
