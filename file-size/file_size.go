package fileSize

import (
	"io"
	"io/ioutil"
	"os"
)

func fileSize1(path string) (int, error) {
	file, err := os.Open(path)
	if err == nil {
		sum := 0
		buf := make([]byte, 2014)
		for {
			n, err := file.Read(buf)
			sum += n
			if err == io.EOF {
				break
			}
		}
		return sum, nil
	}
	return 0, err
}

func fileSize2(path string) (int, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return len(contents), nil
}

func fileSize3(path string) (int64, error) {
	file, err := os.Open(path)
	if err == nil {
		fi, err := file.Stat()
		if err == nil {
			return fi.Size(), nil
		}
	}
	return 0, err
}

func fileSize4(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err == nil {
		return fi.Size(), nil
	}
	return 0, err
}
