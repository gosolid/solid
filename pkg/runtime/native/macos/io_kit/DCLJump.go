//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DCLJump
*/

type DCLJump struct {
  PNextDCLCommand DCLCommand `v8:"pNextDCLCommand"`
  CompilerData uint `v8:"compilerData"`
  Opcode uint `v8:"opcode"`
  PJumpDCLLabel DCLLabel `v8:"pJumpDCLLabel"`
}
