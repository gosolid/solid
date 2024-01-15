//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.UserExportDCLNuDCLLeader
*/

type UserExportDCLNuDCLLeader struct {
  PClientDCLStruct uint64 `v8:"pClientDCLStruct"`
  PNextDCLCommand uint64 `v8:"pNextDCLCommand"`
  CompilerData uint64 `v8:"compilerData"`
  Opcode uint `v8:"opcode"`
  Program uint64 `v8:"program"`
}
