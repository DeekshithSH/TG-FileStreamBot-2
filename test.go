package main

import (
	"EverythingSuckz/fsb/internal/types"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

type HashableFileStruct struct {
	FileName string
	FileSize int64
	MimeType string
	FileID   int64
}

func (f *HashableFileStruct) Pack() string {
	hasher := md5.New()
	val := reflect.ValueOf(*f)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		var fieldValue []byte
		switch field.Kind() {
		case reflect.String:
			fieldValue = []byte(field.String())
		case reflect.Int64:
			fieldValue = []byte(strconv.FormatInt(field.Int(), 10))
		}

		hasher.Write(fieldValue)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func PackFile(fileName string, fileSize int64, mimeType string, fileID int64) string {
	return (&types.HashableFileStruct{FileName: fileName, FileSize: fileSize, MimeType: mimeType, FileID: fileID}).Pack()
}

func main() {
	fullhash := PackFile(
		"photo_2023-11-28_20-31-55.jpg",
		95244,
		"image/jpeg",
		6240121537265930958,
	)
	fmt.Println(fullhash)
}
