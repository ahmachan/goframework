package utils

import (
	"strconv"
	"time"
)

func ConvertInt(val string, d int) int {
	if val == "" {
		return d
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(err.Error())
	}

	return i
}

func ConvertFloat64(val string, d float64) float64 {
	if val == "" {
		return d
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}

	return f
}

func ConvertString(val string) string {
	return string(val)
}


func ConvertTime(val string, layout string, d time.Time) time.Time {
	if val == "" {
		return d
	}
	t, err := time.Parse(layout, val)
	if err != nil {
		panic(err)
	}

	return t
}

