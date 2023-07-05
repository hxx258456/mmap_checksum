package mmap_checksum

import (
	"golang.org/x/exp/mmap"
	"os"
)

func getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func readFileMMAP(filePath string) ([]byte, error) {
	r, err := mmap.Open(filePath)
	if err != nil {
		return nil, err
	}
	result := make([]byte, int64(r.Len()/10))
	if _, err := r.ReadAt(result, int64(r.Len()/10)); err != nil {
		return nil, err
	}
	return result, nil
}
