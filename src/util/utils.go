package util

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	LogPath    = "D:\\practice/go/src/log/" 						//基准路径
	FileFormat = "20060102" 					//时间格式化
	LogFormat  = "2006-01-02 15:04:05 MST Mon"  //时间格式化
	LineFeed   = "\r\n"                         //换行符
)

var(
	year  = strconv.Itoa(time.Now().Year())
	month = strconv.Itoa(int(time.Now().Month()))
	path  = LogPath + year + "/"  + month + "/"
)

//记录日志
func Log(logType string, msg string, err error) error {
	filename := time.Now().Format(FileFormat) + ".log"
	msg = time.Now().Format(LogFormat) + " [" + logType + "] " + "[" + msg + "] " + err.Error()
	res := WriteFile(filename, msg)
	return res
}

//打印错误信息并记录日志
func LogAndPrintln(logType string, msg string, err error) {
	fmt.Println(msg, err)
	_ = Log(logType, msg, err)
	return
}

//写入文件
func WriteFile(fileName, msg string) error {
	if !IsExist(path) {
		_ = CreateDir(path)
	}

	f, err := os.OpenFile(path + fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, msg + LineFeed)

	defer f.Close()
	return err
}

//创建文件
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	_ = os.Chmod(path, os.ModePerm)
	return nil
}

//判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}