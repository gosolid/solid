//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DASDModeParameterBlockDescriptor
*/

type DASDModeParameterBlockDescriptor struct {
  NUMBEROFBLOCKS uint `v8:"nUMBEROFBLOCKS"`
  DENSITYCODE byte `v8:"dENSITYCODE"`
  BLOCKLENGTH [3]byte `v8:"bLOCKLENGTH"`
}
