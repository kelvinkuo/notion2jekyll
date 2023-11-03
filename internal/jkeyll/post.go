package jkeyll

import (
    "fmt"
    "os"
    "path"
    "regexp"
    "strings"
    "time"
    
    "github.com/kelvinkuo/notion2jekyll/internal/post"
    "github.com/kelvinkuo/notion2jekyll/internal/tools"
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

func (p *Post) SetContent(content string) {
    hRegList := []*regexp.Regexp{
        regexp.MustCompile(`#### (.+)\n`),
        regexp.MustCompile(`### (.+)\n`),
        regexp.MustCompile(`## (.+)\n`),
        regexp.MustCompile(`# (.+)\n`),
    }
    
    for _, hReg := range hRegList {
        hList := hReg.FindAllString(content, -1)
        for _, h := range hList {
            content = strings.ReplaceAll(content, h, "#"+h)
        }
    }
    
    // remove title in content
    firstLine := strings.Index(content, "\n")
    content = content[firstLine:]
    
    // add meta
    content = fmt.Sprintf(`---
layout: post
title:  "%s"
date: %s +0800
categories: [%s]
tags: [%s]     # TAG names should always be lowercase1
---`, p.Title, p.CreatedTime.Format(time.DateTime), strings.Join(p.Categories, ", "), strings.Join(p.Tags, ", ")) + content
    
    p.Content = content
}

func (p *Post) CopyTo(dir string) error {
    fileName := fmt.Sprintf("%s-%s.md", p.Meta.CreatedTime.Format(time.DateOnly), strings.ReplaceAll(p.Meta.Title, " ", "-"))
    
    // output images
    imgDir := path.Join(dir, "assets/img", fileName[:len(fileName)-3])
    if err := os.MkdirAll(imgDir, 0777); err != nil {
        return err
    }
    
    for oldPath, file := range p.Images {
        newPath := path.Join(imgDir, path.Base(file.Name()))
        err := tools.Copy(file.Name(), newPath)
        if err != nil {
            return err
        }
        
        // replace link
        p.Content = strings.ReplaceAll(p.Content, oldPath, path.Join("/", newPath))
    }
    
    // output MarkDown file
    if err := os.MkdirAll(path.Join(dir, "_posts"), 0777); err != nil {
        return err
    }
    
    if err := os.WriteFile(path.Join(dir, "_posts", fileName), []byte(p.Content), 0644); err != nil {
        return err
    }
    
    return nil
}
