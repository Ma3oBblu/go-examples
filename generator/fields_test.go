// Code generated by Generator. DO NOT EDIT.
// Source: ./fields.go

package main

import (
	"fmt"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestCountInt64(t *testing.T) {
	tests := []struct {
		name string
		val  int64
		want zap.Field
	}{
		{"correct", int64(123), zap.Int64("count", int64(123))},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := CountInt64(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCount(t *testing.T) {
	tests := []struct {
		name string
		val  int
		want zap.Field
	}{
		{"correct", int(123), CountInt64(int64(123))},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCountInt8(t *testing.T) {
	tests := []struct {
		name string
		val  int8
		want zap.Field
	}{
		{"correct", int8(123), CountInt64(int64(123))},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := CountInt8(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCountInt16(t *testing.T) {
	tests := []struct {
		name string
		val  int16
		want zap.Field
	}{
		{"correct", int16(123), CountInt64(int64(123))},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := CountInt16(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCountInt32(t *testing.T) {
	tests := []struct {
		name string
		val  int32
		want zap.Field
	}{
		{"correct", int32(123), CountInt64(int64(123))},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := CountInt32(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestEmail(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want zap.Field
	}{
		{"correct", "test_string", zap.String("email", "test_string")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Email(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Email() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestError(t *testing.T) {
	tests := []struct {
		name string
		val  error
		want zap.Field
	}{
		{"correct", fmt.Errorf("boo"), zap.Error(fmt.Errorf("boo"))},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Error(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}