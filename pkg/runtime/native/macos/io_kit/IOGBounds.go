//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOGBounds
*/

type IOGBounds struct {
  Minx int16 `v8:"minx"`
  Maxx int16 `v8:"maxx"`
  Miny int16 `v8:"miny"`
  Maxy int16 `v8:"maxy"`
}
