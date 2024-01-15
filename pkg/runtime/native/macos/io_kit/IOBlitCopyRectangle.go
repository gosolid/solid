//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitCopyRectangle
*/

type IOBlitCopyRectangle struct {
  SourceX int `v8:"sourceX"`
  SourceY int `v8:"sourceY"`
  X int `v8:"x"`
  Y int `v8:"y"`
  Width int `v8:"width"`
  Height int `v8:"height"`
}
