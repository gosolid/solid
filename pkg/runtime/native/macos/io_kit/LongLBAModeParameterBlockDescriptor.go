//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.LongLBAModeParameterBlockDescriptor
*/

type LongLBAModeParameterBlockDescriptor struct {
  NUMBEROFBLOCKS uint64 `v8:"nUMBEROFBLOCKS"`
  DENSITYCODE byte `v8:"dENSITYCODE"`
  RESERVED [3]byte `v8:"rESERVED"`
  BLOCKLENGTH uint `v8:"bLOCKLENGTH"`
}
