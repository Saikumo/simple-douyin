package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	. "saikumo.org/simple-douyin/src/common"
)

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
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

func GetFileUrl(fileName string) string {
	url := fmt.Sprintf(`http://%s:%d/static/%s`, Config.GetString("server.ip"),
		Config.GetInt("server.port"), fileName)
	return url
}
