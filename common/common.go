package common

import (
	"io/ioutil"
	"bytes"
)

func WriteToFile(filename string, data []byte) {
	ioutil.WriteFile(filename, data, 0644)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ByteSliceToString(data []byte) string {
	n := bytes.IndexByte(data, 0)
	return string(data[:n])
}