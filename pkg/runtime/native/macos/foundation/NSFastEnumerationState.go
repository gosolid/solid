//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSFastEnumerationState
*/

type NSFastEnumerationState struct {
  State uint64 `v8:"state"`
  ItemsPtr *any `v8:"itemsPtr"`
  MutationsPtr uint64 `v8:"mutationsPtr"`
  Extra [5]uint64 `v8:"extra"`
}
