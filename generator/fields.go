package main

import (
	"go.uber.org/zap"
)

func CountInt64(val int64) zap.Field {
	return zap.Int64("count", val)
}
func Count(val int) zap.Field {
	return CountInt64(int64(val))
}
func CountInt8(val int8) zap.Field {
	return CountInt64(int64(val))
}
func CountInt16(val int16) zap.Field {
	return CountInt64(int64(val))
}
func CountInt32(val int32) zap.Field {
	return CountInt64(int64(val))
}
func Email(val string) zap.Field {
	return zap.String("email", val)
}
func Error(err error) zap.Field {
	return zap.Error(err)
}
