package post

import (
    "os"
)

type CommonPost struct {
    Meta
    Content string
    Images  map[string]*os.File
}

func (p *CommonPost) GetContent() string {
    return p.Content
}

func (p *CommonPost) GetTitle() string {
    return p.Title
}

func (p *CommonPost) GetMeta() Meta {
    return p.Meta
}

func (p *CommonPost) GetImages() map[string]*os.File {
    return p.Images
}

// func (p *CommonPost) SetContent(content string) {
//     p.Content = content
// }

func (p *CommonPost) SetMeta(meta *Meta) {
    p.Meta = *meta
}

func (p *CommonPost) SetImages(images map[string]*os.File) {
    p.Images = images
}
