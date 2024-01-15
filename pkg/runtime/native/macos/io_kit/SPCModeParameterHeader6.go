//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SPCModeParameterHeader6
*/

type SPCModeParameterHeader6 struct {
  MODEDATALENGTH byte `v8:"mODEDATALENGTH"`
  MEDIUMTYPE byte `v8:"mEDIUMTYPE"`
  DEVICESPECIFICPARAMETER byte `v8:"dEVICESPECIFICPARAMETER"`
  BLOCKDESCRIPTORLENGTH byte `v8:"bLOCKDESCRIPTORLENGTH"`
}
