package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type LocalStorage struct {
	BasePath string
}
type Storage interface {
	SaveAvatar(file *multipart.FileHeader) (string, error)
	DeleteAvatar(path string) error
}

func NewLocalStorage() Storage {
	return &LocalStorage{}
}

func (l *LocalStorage) SaveAvatar(file *multipart.FileHeader) (string, error) {

	filename := fmt.Sprintf(
		"%d_%s",
		time.Now().Unix(),
		file.Filename,
	)

	dst := filepath.Join(l.BasePath, filename)

	if err := os.MkdirAll(l.BasePath, 0755); err != nil {
		return "", err
	}

	if err := saveMultipartFile(file, dst); err != nil {
		return "", err
	}

	return filename, nil
}

func (l *LocalStorage) DeleteAvatar(path string) error {
	return os.Remove(filepath.Join(l.BasePath, path))
}

func saveMultipartFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
