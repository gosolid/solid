//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DCLUpdateDCLList
*/

type DCLUpdateDCLList struct {
  PNextDCLCommand DCLCommand `v8:"pNextDCLCommand"`
  CompilerData uint `v8:"compilerData"`
  Opcode uint `v8:"opcode"`
  DclCommandList DCLCommand `v8:"dclCommandList"`
  NumDCLCommands uint `v8:"numDCLCommands"`
}
