package notion

import (
    "errors"
    "net/url"
    "os"
    "path"
    "path/filepath"
    "regexp"
    "strings"
    
    "github.com/kelvinkuo/notion2jekyll/internal/post"
)

type Post struct {
    post.CommonPost
}

func NewPost() *Post {
    return &Post{
        CommonPost: post.CommonPost{
            Meta: post.Meta{
                Categories: make([]string, 0),
                Tags:       make([]string, 0),
            },
            Images: make(map[string]*os.File),
        },
    }
}

func NewPostWithDir(root string) (*Post, error) {
    p := NewPost()
    
    // fill content
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if !info.IsDir() {
            if strings.HasSuffix(info.Name(), ".md") {
                content, err := os.ReadFile(path)
                if err != nil {
                    return err
                }
                p.SetContent(string(content))
                return nil
            }
        }
        
        return nil
    })
    if err != nil {
        return nil, err
    }
    
    // fill images
    imageReg := regexp.MustCompile(`!\[.+\]\(.+\)`)
    // imageTitleReg := regexp.MustCompile(`\[.+\]`)
    imagePathReg := regexp.MustCompile(`\(.+\)`)
    images := imageReg.FindAllString(p.Content, -1)
    for _, image := range images {
        // imageTitle := imageTitleReg.FindString(image)
        imagePath := imagePathReg.FindString(image)
        imagePath = imagePath[1 : len(imagePath)-1]
        imagePathDecoded, err := url.QueryUnescape(imagePath)
        if err != nil {
            return nil, err
        }
        imagePathDecoded = path.Join(root, imagePathDecoded)
        imageFile, err := os.Open(imagePathDecoded)
        if err != nil {
            return nil, err
        }
        p.CommonPost.Images[imagePath] = imageFile
    }
    
    // fill title
    titleReg := regexp.MustCompile(`# (.+)\n`)
    titles := titleReg.FindAllString(p.Content, -1)
    if len(titles) > 0 {
        p.CommonPost.Title = titles[0][2 : len(titles[0])-1]
    } else {
        return nil, errors.New("can not find title")
    }
    
    return p, nil
}

func (p *Post) SetContent(content string) {
    p.Content = content
}

func (p *Post) CopyTo(dir string) error {
    // TODO
    return nil
}
