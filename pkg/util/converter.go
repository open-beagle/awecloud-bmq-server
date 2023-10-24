package util

import "strconv"

func Stoi(snum string) int32 {
	inum, _ := strconv.Atoi(snum)
	return int32(inum)
}
