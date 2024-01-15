//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXSize
*/

type NXSize struct {
  Width float32 `v8:"width"`
  Height float32 `v8:"height"`
}
