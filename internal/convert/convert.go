package convert

import (
    "log"
    "time"
    
    "github.com/kelvinkuo/notion2jekyll/internal/jkeyll"
    "github.com/kelvinkuo/notion2jekyll/internal/notion"
    "github.com/kelvinkuo/notion2jekyll/internal/post"
)

func Notion2Jkeyll(author, createdTime, notionDir, jkeyllDir string, categories, tags []string) {
    t, err := time.Parse(time.DateTime, createdTime)
    if err != nil {
        log.Fatalf("create time format err = %s", err)
    }
    
    nPost, err := notion.NewPostWithDir(notionDir)
    if err != nil {
        log.Fatalf("notion dir %s is error", notionDir)
    }
    
    jPost := jkeyll.NewPost()
    jPost.SetMeta(&post.Meta{
        Author:      author,
        Title:       nPost.GetTitle(),
        CreatedTime: t,
        Tags:        tags,
        Categories:  categories,
    })
    jPost.SetContent(nPost.Content)
    jPost.SetImages(nPost.GetImages())
    err = jPost.CopyTo(jkeyllDir)
    if err != nil {
        log.Fatal(err)
    }
}
