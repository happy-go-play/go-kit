package gokit

import (
	"fmt"
	"os"
)

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return fmt.Errorf("could not create directory: %w", err)
			}
		} else {
			return fmt.Errorf("could not stat directory: %w", err)
		}
	}
	return nil
}

func AppendToFile(filename string, data []byte) error {
	// Open file, create if not exists
	// 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	// Write log data, append newline
	// 写入日志数据，并追加换行符
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}
