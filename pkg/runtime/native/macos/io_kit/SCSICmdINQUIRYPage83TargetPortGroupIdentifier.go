//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier
*/

type SCSICmdINQUIRYPage83TargetPortGroupIdentifier struct {
  RESERVED uint16 `v8:"rESERVED"`
  TARGETPORTGROUP uint16 `v8:"tARGETPORTGROUP"`
}
