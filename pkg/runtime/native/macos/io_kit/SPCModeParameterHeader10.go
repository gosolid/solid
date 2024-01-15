//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SPCModeParameterHeader10
*/

type SPCModeParameterHeader10 struct {
  MODEDATALENGTH uint16 `v8:"mODEDATALENGTH"`
  MEDIUMTYPE byte `v8:"mEDIUMTYPE"`
  DEVICESPECIFICPARAMETER byte `v8:"dEVICESPECIFICPARAMETER"`
  LONGLBA byte `v8:"lONGLBA"`
  RESERVED byte `v8:"rESERVED"`
  BLOCKDESCRIPTORLENGTH uint16 `v8:"bLOCKDESCRIPTORLENGTH"`
}
