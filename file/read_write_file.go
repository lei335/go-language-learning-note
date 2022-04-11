package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 1.1、将文件整个读入内存（效率比较高，占用内存也最高）
func read0(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)

	return string(content), nil
}

// 1.2、将文件整个读入内存（效率比较高，占用内存也最高）
func read1(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// 2、按字节读取文件
func read2(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer fi.Close()

	r := bufio.NewReader(fi)

	chunks := make([]byte, 0)
	buf := make([]byte, 1024) // 一次读取多少个字节
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		fmt.Println(string(buf[:n]))
		if n == 0 {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
	return string(chunks), nil
}

// 3、按行读取文件
func read3(path string) (string, error) {
	fi, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}

	defer fi.Close()

	stat, err := fi.Stat()
	if err != nil {
		return "", err
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(fi)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return "", err
			}
		}

	}
	return "", nil
}

// 如果文件不存在，则创建文件；如果存在，就会覆盖写
func write0() error {
	content := []byte("test1\ntest2\n")
	err := ioutil.WriteFile("test.txt", content, 0644)
	if err != nil {
		return err
	}
	return nil
}

// 追加写，也可以指定成覆盖写
func write1() error {
	content := "test1\ntest2\n"
	filename := "./test.txt"

	var f *os.File

	_, err := os.Stat(filename)
	if os.IsNotExist(err) { // 文件不存在
		f, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) // 打开文件
		if err != nil {
			return err
		}
	}

	defer f.Close()

	n, err := io.WriteString(f, content) // 写入文件
	if err != nil {
		return err
	}

	fmt.Println("写入 ", n, " 个字节")
	return nil
}

// 追加写
func write2() error {
	content := "test1\ntest2\n"
	filename := "./test.txt"

	var f *os.File

	_, err := os.Stat(filename)
	if os.IsNotExist(err) { // 文件不存在
		f, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) // 打开文件
		if err != nil {
			return err
		}
	}

	defer f.Close()

	n, err := f.Write([]byte(content)) // 写入文件（字节数组）
	if err != nil {
		return err
	}
	fmt.Println("写入 ", n, " 个字节")

	n, err = f.WriteString(content) // 写入文件（字符串）
	if err != nil {
		return err
	}
	fmt.Println("写入 ", n, " 个字节")

	f.Sync()
	return nil
}

// 追加写
func write3() error {
	content := "test1\ntest2\n"
	filename := "./test.txt"

	var f *os.File

	_, err := os.Stat(filename)
	if os.IsNotExist(err) { // 文件不存在
		f, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) // 打开文件
		if err != nil {
			return err
		}
	}

	defer f.Close()

	w := bufio.NewWriter(f) // 创建新的 Writer 对象
	n, _ := w.WriteString(content)
	fmt.Println("写入 ", n, " 个字节")
	w.Flush()

	return nil
}
