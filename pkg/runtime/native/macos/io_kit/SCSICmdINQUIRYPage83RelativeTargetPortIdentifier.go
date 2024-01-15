//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier
*/

type SCSICmdINQUIRYPage83RelativeTargetPortIdentifier struct {
  OBSOLETE uint16 `v8:"oBSOLETE"`
  RELATIVETARGETPORTIDENTIFIER uint16 `v8:"rELATIVETARGETPORTIDENTIFIER"`
}
