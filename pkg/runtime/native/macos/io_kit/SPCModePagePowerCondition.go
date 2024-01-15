//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SPCModePagePowerCondition
*/

type SPCModePagePowerCondition struct {
  Header ModePageFormatHeader `v8:"header"`
  RESERVED byte `v8:"rESERVED"`
  IDLESTANDBY byte `v8:"iDLESTANDBY"`
  IDLECONDITIONTIMER uint `v8:"iDLECONDITIONTIMER"`
  STANDBYCONDITIONTIMER uint `v8:"sTANDBYCONDITIONTIMER"`
}
