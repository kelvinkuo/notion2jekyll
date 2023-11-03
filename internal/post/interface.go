package post

import (
    "os"
    "time"
)

type Meta struct {
    Author      string
    Title       string
    CreatedTime time.Time
    Tags        []string
    Categories  []string
}

type Post interface {
    GetContent() string
    GetTitle() string
    GetMeta() Meta
    GetImages() map[string]*os.File
    SetContent(content string)
    SetMeta(meta *Meta)
    SetImages(images map[string]*os.File)
    CopyTo(dir string) error
}

// type Post interface {
//     // api for other post
//     GetContent() string
//     Meta() *Meta
//     GetImages() map[string]*os.File
//
//     // api for create post
//     SetContent(content string)
//     // SetMeta(meta *Meta)
//     CreateFiles() error
// }
