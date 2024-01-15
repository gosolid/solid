//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDMSF
*/

type CDMSF struct {
  Minute byte `v8:"minute"`
  Second byte `v8:"second"`
  Frame byte `v8:"frame"`
}
