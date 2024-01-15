//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor
*/

type SCSICmdINQUIRYPageB2ProvisioningGroupDescriptor struct {
  DESIGNATIONDESCRIPTOR [20]byte `v8:"dESIGNATIONDESCRIPTOR"`
  RESERVED [18]byte `v8:"rESERVED"`
}
