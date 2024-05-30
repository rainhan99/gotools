package convert

import (
	"time"
)

// project gotools
// file convert/convert
// Create by RainHan on 2024/05/30 14:48:50

func Bool(b bool) *bool {
	return &b
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Int8(i int8) *int8 {
	return &i
}

func Int16(i int16) *int16 {
	return &i
}

func Int32(i int32) *int32 {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func Float32(f float32) *float32 {
	return &f
}

func Float64(f float64) *float64 {
	return &f
}

func Uint(i uint) *uint {
	return &i
}

func Uint8(i uint8) *uint8 {
	return &i
}

func Uint16(i uint16) *uint16 {
	return &i
}

func Uint32(i uint32) *uint32 {
	return &i
}

func Uint64(i uint64) *uint64 {
	return &i
}

func Time(t time.Time) *time.Time {
	return &t
}

func Duration(d time.Duration) *time.Duration {
	return &d
}

func Byte(b byte) *byte {
	return &b
}

func Rune(r rune) *rune {
	return &r
}

func ToBool(b *bool) (c bool) {
	if b == nil {
		return c
	}
	return *b
}

func ToString(s *string) (t string) {
	if s == nil {
		return t
	}
	return *s
}

func ToInt(i *int) (t int) {
	if i == nil {
		return t
	}
	return *i
}

func ToInt8(i *int8) (t int8) {
	if i == nil {
		return t
	}
	return *i
}

func ToInt16(i *int16) (t int16) {
	if i == nil {
		return t
	}
	return *i
}

func ToInt32(i *int32) (t int32) {
	if i == nil {
		return t
	}
	return *i
}

func ToInt64(i *int64) (t int64) {
	if i == nil {
		return t
	}
	return *i
}

func ToFloat32(f *float32) (t float32) {
	if f == nil {
		return t
	}
	return *f
}

func ToFloat64(f *float64) (t float64) {
	if f == nil {
		return t
	}
	return *f
}

func ToUint(u *uint) (t uint) {
	if u == nil {
		return t
	}
	return *u
}

func ToUint8(u *uint8) (t uint8) {
	if u == nil {
		return t
	}
	return *u
}

func ToUint16(u *uint16) (t uint16) {
	if u == nil {
		return t
	}
	return *u
}

func ToUint32(u *uint32) (t uint32) {
	if u == nil {
		return t
	}
	return *u
}

func ToUint64(u *uint64) (t uint64) {
	if u == nil {
		return t
	}
	return *u
}

func ToTime(t *time.Time) (v time.Time) {
	if t == nil {
		return v
	}
	return *t
}

func ToDuration(d *time.Duration) (t time.Duration) {
	if d == nil {
		return t
	}
	return *d
}

func ToByte(b *byte) (t byte) {
	if b == nil {
		return t
	}
	return *b
}

func ToRune(r *rune) (t rune) {
	if r == nil {
		return t
	}
	return *r
}
