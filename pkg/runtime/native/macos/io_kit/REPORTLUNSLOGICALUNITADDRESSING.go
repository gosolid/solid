//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.REPORT_LUNS_LOGICAL_UNIT_ADDRESSING
*/

type REPORTLUNSLOGICALUNITADDRESSING struct {
  LUN uint16 `v8:"lUN"`
  BUSNUMBER uint16 `v8:"bUSNUMBER"`
  TARGET uint16 `v8:"tARGET"`
  Reserved2 uint16 `v8:"reserved2"`
  Reserved uint16 `v8:"reserved"`
}
