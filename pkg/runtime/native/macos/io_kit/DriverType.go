//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.DriverType
*/

type DriverType struct {
  NameInfoStr [32]byte `v8:"nameInfoStr"`
  Version objc.NumVersion `v8:"version"`
}
