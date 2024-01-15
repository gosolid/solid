//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.GetVolParmsInfoBuffer
*/

type GetVolParmsInfoBuffer struct {
  VMVersion int16 `v8:"vMVersion"`
  VMAttrib int `v8:"vMAttrib"`
  VMLocalHand **byte `v8:"vMLocalHand"`
  VMServerAdr int `v8:"vMServerAdr"`
  VMVolumeGrade int `v8:"vMVolumeGrade"`
  VMForeignPrivID int16 `v8:"vMForeignPrivID"`
  VMExtendedAttributes int `v8:"vMExtendedAttributes"`
  VMDeviceID void `v8:"vMDeviceID"`
  VMMaxNameLength uint64 `v8:"vMMaxNameLength"`
}
