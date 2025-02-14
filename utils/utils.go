package kss

import (
	"log"
	"os"
	"path/filepath"
)

// CreatDir 创建目录（父目录不存在则新建）
func CreatDir(dirPath string) {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatalf("创建目录失败: %v", err)
	}
	log.Printf("目录 %v 创建成功", dirPath)
}

// CreatFile 新建文件（父目录不存在则新建）
func CreatFile(filePath string) {
	CreatDir(filepath.Dir(filePath))
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()
	log.Printf("文件 %v 创建成功", filePath)
}
