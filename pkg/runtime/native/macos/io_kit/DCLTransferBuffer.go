//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DCLTransferBuffer
*/

type DCLTransferBuffer struct {
  PNextDCLCommand DCLCommand `v8:"pNextDCLCommand"`
  CompilerData uint `v8:"compilerData"`
  Opcode uint `v8:"opcode"`
  Buffer void `v8:"buffer"`
  Size uint `v8:"size"`
  PacketSize uint16 `v8:"packetSize"`
  Reserved uint16 `v8:"reserved"`
  BufferOffset uint `v8:"bufferOffset"`
}
