package storage

import (
    "io"
    "os"
    "path/filepath"
)

func UploadFile(file io.Reader, filename string) (string, error) {
    uploadDir := "./uploads"
    os.MkdirAll(uploadDir, os.ModePerm)
    filePath := filepath.Join(uploadDir, filename)
    out, err := os.Create(filePath)
    if err != nil {
        return "", err
    }
    defer out.Close()
    _, err = io.Copy(out, file)
    if err != nil {
        return "", err
    }
    return "http://localhost:8080/uploads/" + filename, nil
}

func DeleteFile(filename string) error {
    filePath := filepath.Join("./uploads", filename)
    return os.Remove(filePath)
}