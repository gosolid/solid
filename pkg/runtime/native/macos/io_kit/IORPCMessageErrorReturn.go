//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IORPCMessageErrorReturn
*/

type IORPCMessageErrorReturn struct {
  Mach IORPCMessageMach `v8:"mach"`
  Content IORPCMessageErrorReturnContent `v8:"content"`
}
