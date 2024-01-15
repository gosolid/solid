//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDCommunicationRec
*/

type VDCommunicationRec struct {
  CsBusID int `v8:"csBusID"`
  CsCommFlags uint `v8:"csCommFlags"`
  CsMinReplyDelay uint `v8:"csMinReplyDelay"`
  CsReserved2 uint `v8:"csReserved2"`
  CsSendAddress uint `v8:"csSendAddress"`
  CsSendType uint `v8:"csSendType"`
  CsSendBuffer *void `v8:"csSendBuffer"`
  CsSendSize uint64 `v8:"csSendSize"`
  CsReplyAddress uint `v8:"csReplyAddress"`
  CsReplyType uint `v8:"csReplyType"`
  CsReplyBuffer *void `v8:"csReplyBuffer"`
  CsReplySize uint64 `v8:"csReplySize"`
  CsReserved3 uint `v8:"csReserved3"`
  CsReserved4 uint `v8:"csReserved4"`
  CsReserved5 uint `v8:"csReserved5"`
  CsReserved6 uint `v8:"csReserved6"`
}
