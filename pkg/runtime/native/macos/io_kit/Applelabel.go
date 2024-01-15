//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.applelabel
*/

type Applelabel struct {
  AlBoot0 [416]byte `v8:"alBoot0"`
  AlMagic uint16 `v8:"alMagic"`
  AlType uint16 `v8:"alType"`
  AlFlags uint `v8:"alFlags"`
  AlOffset uint64 `v8:"alOffset"`
  AlSize uint `v8:"alSize"`
  AlChecksum uint `v8:"alChecksum"`
  AlBoot1 [72]byte `v8:"alBoot1"`
}
