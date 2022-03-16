package common

import "strconv"

func Str2Int64(data string) error {
	_, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return err
	}
	return nil
}
