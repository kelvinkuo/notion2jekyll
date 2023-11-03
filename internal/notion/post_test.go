package notion

import (
    "reflect"
    "testing"
    
    "github.com/kelvinkuo/notion2jekyll/internal/post"
)

func TestNewPostWithDir(t *testing.T) {
    type args struct {
        path string
    }
    tests := []struct {
        name string
        args args
        want post.Meta
    }{
        {
            args: args{path: "./test/testnotion/9980ca75-e32b-48ef-b4a6-f4066388bbfd_Export-7b1c344d-5502-4593-b996-c13b057e05f3"},
            want: post.Meta{
                Title:      "Nexus 7 2013 刷机",
                Categories: make([]string, 0),
                Tags:       make([]string, 0),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := NewPostWithDir(tt.args.path)
            if err != nil {
                t.Errorf("NewPostWithDir() err = %s", err)
            }
            if !reflect.DeepEqual(got.CommonPost.Meta, tt.want) {
                t.Errorf("NewPostWithDir() = %v, want %v", got.CommonPost.Meta, tt.want)
            }
        })
    }
}
