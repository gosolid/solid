//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyboardLayout
*/

type UCKeyboardLayout struct {
  KeyLayoutHeaderFormat uint16 `v8:"keyLayoutHeaderFormat"`
  KeyLayoutDataVersion uint16 `v8:"keyLayoutDataVersion"`
  KeyLayoutFeatureInfoOffset uint `v8:"keyLayoutFeatureInfoOffset"`
  KeyboardTypeCount uint `v8:"keyboardTypeCount"`
  KeyboardTypeList [1]UCKeyboardTypeHeader `v8:"keyboardTypeList"`
}
