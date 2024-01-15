//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyboardTypeHeader
*/

type UCKeyboardTypeHeader struct {
  KeyboardTypeFirst uint `v8:"keyboardTypeFirst"`
  KeyboardTypeLast uint `v8:"keyboardTypeLast"`
  KeyModifiersToTableNumOffset uint `v8:"keyModifiersToTableNumOffset"`
  KeyToCharTableIndexOffset uint `v8:"keyToCharTableIndexOffset"`
  KeyStateRecordsIndexOffset uint `v8:"keyStateRecordsIndexOffset"`
  KeyStateTerminatorsOffset uint `v8:"keyStateTerminatorsOffset"`
  KeySequenceDataIndexOffset uint `v8:"keySequenceDataIndexOffset"`
}
