package util

import "strconv"

func StringToInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int(i), err
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
