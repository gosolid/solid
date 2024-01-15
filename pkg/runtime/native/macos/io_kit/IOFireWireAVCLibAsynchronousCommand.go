//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFireWireAVCLibAsynchronousCommand
*/

type IOFireWireAVCLibAsynchronousCommand struct {
  CmdState int `v8:"cmdState"`
  PRefCon void `v8:"pRefCon"`
  PCommandBuf byte `v8:"pCommandBuf"`
  CmdLen uint `v8:"cmdLen"`
  PInterimResponseBuf byte `v8:"pInterimResponseBuf"`
  InterimResponseLen uint `v8:"interimResponseLen"`
  PFinalResponseBuf byte `v8:"pFinalResponseBuf"`
  FinalResponseLen uint `v8:"finalResponseLen"`
}
