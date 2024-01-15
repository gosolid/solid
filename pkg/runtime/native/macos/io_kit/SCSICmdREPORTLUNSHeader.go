//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_REPORT_LUNS_Header
*/

type SCSICmdREPORTLUNSHeader struct {
  LUNLISTLENGTH uint `v8:"lUNLISTLENGTH"`
  RESERVED uint `v8:"rESERVED"`
  LUN [1]SCSICmdREPORTLUNSLUNENTRY `v8:"lUN"`
}
