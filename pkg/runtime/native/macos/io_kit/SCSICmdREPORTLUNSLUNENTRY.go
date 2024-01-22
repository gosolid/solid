//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_REPORT_LUNS_LUN_ENTRY
*/

type SCSICmdREPORTLUNSLUNENTRY struct {
  FIRSTLEVELADDRESSING uint16 `v8:"fIRSTLEVELADDRESSING"`
  SECONDLEVELADDRESSING uint16 `v8:"sECONDLEVELADDRESSING"`
  THIRDLEVELADDRESSING uint16 `v8:"tHIRDLEVELADDRESSING"`
  FOURTHLEVELADDRESSING uint16 `v8:"fOURTHLEVELADDRESSING"`
}