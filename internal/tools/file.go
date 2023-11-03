package tools

import (
    "io"
    "os"
)

func Copy(src, dest string) error {
    s, err := os.Open(src)
    if err != nil {
        return err
    }
    defer s.Close()
    
    d, err := os.Create(dest)
    if err != nil {
        panic(err)
    }
    defer d.Close()
    
    _, err = io.Copy(d, s)
    if err != nil {
        return err
    }
    
    return nil
}
