package test_utils

import "os"

func ReadFileOrErr(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err.Error()
	}
	return string(content)
}
