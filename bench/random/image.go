package random

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
)

var files []fs.FileInfo

func init() {
	var err error
	// 画像ファイル群の読み込み
	files, err = ioutil.ReadDir(imageFolderPath)
	if err != nil {
		log.Fatalf("%+v", fmt.Errorf("%w", err))
	}
}

func Image() ([]byte, error) {
	fileInfo := files[rand.Intn(len(files))]
	return ioutil.ReadFile(filepath.Join(imageFolderPath, fileInfo.Name()))
}
