package main

import (
 "archive/zip"
 "bytes"
 "io"
 "os"
)

func main() {
 data, err := io.ReadAll(os.Stdin); if err != nil { panic(err) }
 zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data))); if err != nil { panic(err) }
 if len(zr.File) != 2 { panic("entry count") }
 for _, f := range zr.File {
  r, err := f.Open(); if err != nil { panic(err) }
  b, err := io.ReadAll(r); if err != nil { panic(err) }; r.Close()
  if f.Name == "sample" { for i, c := range b { if c != byte(i%251) { panic("payload") } }; if len(b)!=100003 { panic("size") } } else if f.Name!="empty" || len(b)!=0 { panic("empty") }
 }
}
