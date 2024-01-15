//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitVertices
*/

type IOBlitVertices struct {
  Operation IOBlitOperation `v8:"operation"`
  Count uint `v8:"count"`
  Vertices [2]IOBlitVertex `v8:"vertices"`
}
