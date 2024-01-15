//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.disk_blk0
*/

type DiskBlk0 struct {
  Bootcode [446]byte `v8:"bootcode"`
  Parts [4]FdiskPart `v8:"parts"`
  Signature uint16 `v8:"signature"`
}
