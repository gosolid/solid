//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOGSize
*/

type IOGSize struct {
  Width int16 `v8:"width"`
  Height int16 `v8:"height"`
}
