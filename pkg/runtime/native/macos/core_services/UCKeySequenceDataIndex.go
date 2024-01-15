//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeySequenceDataIndex
*/

type UCKeySequenceDataIndex struct {
  KeySequenceDataIndexFormat uint16 `v8:"keySequenceDataIndexFormat"`
  CharSequenceCount uint16 `v8:"charSequenceCount"`
  CharSequenceOffsets [1]uint16 `v8:"charSequenceOffsets"`
}
