//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXKeyMapping
*/

type NXKeyMapping struct {
  Size int `v8:"size"`
  Mapping byte `v8:"mapping"`
}
