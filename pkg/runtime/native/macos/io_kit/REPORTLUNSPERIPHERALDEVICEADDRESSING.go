//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING
*/

type REPORTLUNSPERIPHERALDEVICEADDRESSING struct {
  TARGETLUN uint16 `v8:"tARGETLUN"`
  BUSIDENTIFIER uint16 `v8:"bUSIDENTIFIER"`
  Reserved2 uint16 `v8:"reserved2"`
  Reserved uint16 `v8:"reserved"`
}
