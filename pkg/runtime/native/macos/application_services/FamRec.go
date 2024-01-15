//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.FamRec
*/

type FamRec struct {
  FfFlags int16 `v8:"ffFlags"`
  FfFamID int16 `v8:"ffFamID"`
  FfFirstChar int16 `v8:"ffFirstChar"`
  FfLastChar int16 `v8:"ffLastChar"`
  FfAscent int16 `v8:"ffAscent"`
  FfDescent int16 `v8:"ffDescent"`
  FfLeading int16 `v8:"ffLeading"`
  FfWidMax int16 `v8:"ffWidMax"`
  FfWTabOff int `v8:"ffWTabOff"`
  FfKernOff int `v8:"ffKernOff"`
  FfStylOff int `v8:"ffStylOff"`
  FfProperty [9]int16 `v8:"ffProperty"`
  FfIntl [2]int16 `v8:"ffIntl"`
  FfVersion int16 `v8:"ffVersion"`
}
