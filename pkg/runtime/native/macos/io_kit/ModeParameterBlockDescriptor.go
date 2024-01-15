//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.ModeParameterBlockDescriptor
*/

type ModeParameterBlockDescriptor struct {
  DENSITYCODE byte `v8:"dENSITYCODE"`
  NUMBEROFBLOCKS [3]byte `v8:"nUMBEROFBLOCKS"`
  RESERVED byte `v8:"rESERVED"`
  BLOCKLENGTH [3]byte `v8:"bLOCKLENGTH"`
}
