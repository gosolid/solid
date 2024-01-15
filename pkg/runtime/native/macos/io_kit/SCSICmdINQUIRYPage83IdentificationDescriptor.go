//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSICmd_INQUIRY_Page83_Identification_Descriptor
*/

type SCSICmdINQUIRYPage83IdentificationDescriptor struct {
  CODESET byte `v8:"cODESET"`
  IDENTIFIERTYPE byte `v8:"iDENTIFIERTYPE"`
  RESERVED byte `v8:"rESERVED"`
  IDENTIFIERLENGTH byte `v8:"iDENTIFIERLENGTH"`
  IDENTIFIER byte `v8:"iDENTIFIER"`
}
