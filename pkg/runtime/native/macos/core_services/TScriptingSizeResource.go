//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TScriptingSizeResource
*/

type TScriptingSizeResource struct {
  ScriptingSizeFlags int16 `v8:"scriptingSizeFlags"`
  MinStackSize uint `v8:"minStackSize"`
  PreferredStackSize uint `v8:"preferredStackSize"`
  MaxStackSize uint `v8:"maxStackSize"`
  MinHeapSize uint `v8:"minHeapSize"`
  PreferredHeapSize uint `v8:"preferredHeapSize"`
  MaxHeapSize uint `v8:"maxHeapSize"`
}
