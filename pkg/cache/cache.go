package cache

import (
    "sync"
    "time"
)

type File struct {
    ID         int
    FileName   string
    UploadDate time.Time
    Size       int64
    URL        string
}

var (
    cache = make(map[string][]File)
    mu    sync.RWMutex
)

func GetFiles(userID string) ([]File, bool) {
    mu.RLock()
    defer mu.RUnlock()
    files, ok := cache[userID]
    return files, ok
}

func SetFiles(userID string, files []File) {
    mu.Lock()
    defer mu.Unlock()
    cache[userID] = files
}

func Invalidate(userID string) {
    mu.Lock()
    defer mu.Unlock()
    delete(cache, userID)
}