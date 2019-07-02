package models

import "time"

func PtrS(s string) *string {
	return &s
}

func PtrI64(i int64) *int64 {
	return &i
}

func PtrT(t time.Time) *time.Time {
	return &t
}
