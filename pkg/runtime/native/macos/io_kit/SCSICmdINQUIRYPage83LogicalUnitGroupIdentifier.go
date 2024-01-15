//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier
*/

type SCSICmdINQUIRYPage83LogicalUnitGroupIdentifier struct {
  RESERVED uint16 `v8:"rESERVED"`
  LOGICALUNITGROUP uint16 `v8:"lOGICALUNITGROUP"`
}
