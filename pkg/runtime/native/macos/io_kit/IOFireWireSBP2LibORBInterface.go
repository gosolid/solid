//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFireWireSBP2LibORBInterface
*/

type IOFireWireSBP2LibORBInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  SetRefCon func(...any) (any, error) `v8:"setRefCon"`
  GetRefCon func(...any) (any, error) `v8:"getRefCon"`
  SetCommandFlags func(...any) (any, error) `v8:"setCommandFlags"`
  SetMaxORBPayloadSize func(...any) (any, error) `v8:"setMaxORBPayloadSize"`
  SetCommandTimeout func(...any) (any, error) `v8:"setCommandTimeout"`
  SetCommandGeneration func(...any) (any, error) `v8:"setCommandGeneration"`
  SetCommandBuffersAsRanges func(...any) (any, error) `v8:"setCommandBuffersAsRanges"`
  ReleaseCommandBuffers func(...any) (any, error) `v8:"releaseCommandBuffers"`
  SetCommandBlock func(...any) (any, error) `v8:"setCommandBlock"`
  LSIWorkaroundSetCommandBuffersAsRanges func(...any) (any, error) `v8:"lSIWorkaroundSetCommandBuffersAsRanges"`
  LSIWorkaroundSyncBuffersForOutput func(...any) (any, error) `v8:"lSIWorkaroundSyncBuffersForOutput"`
  LSIWorkaroundSyncBuffersForInput func(...any) (any, error) `v8:"lSIWorkaroundSyncBuffersForInput"`
}
