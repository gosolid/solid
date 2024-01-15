//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DCLCallProc
*/

type DCLCallProc struct {
  PNextDCLCommand DCLCommand `v8:"pNextDCLCommand"`
  CompilerData uint `v8:"compilerData"`
  Opcode uint `v8:"opcode"`
  Proc func(...any) (any, error) `v8:"proc"`
  ProcData *void `v8:"procData"`
}
