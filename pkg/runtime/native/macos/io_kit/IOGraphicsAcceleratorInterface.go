//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOGraphicsAcceleratorInterface
*/

type IOGraphicsAcceleratorInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  Probe func(...any) (any, error) `v8:"probe"`
  Start func(...any) (any, error) `v8:"start"`
  Stop func(...any) (any, error) `v8:"stop"`
  Reset func(...any) (any, error) `v8:"reset"`
  CopyCapabilities func(...any) (any, error) `v8:"copyCapabilities"`
  GetBlitProc func(...any) (any, error) `v8:"getBlitProc"`
  Flush func(...any) (any, error) `v8:"flush"`
  WaitForCompletion func(...any) (any, error) `v8:"waitForCompletion"`
  Synchronize func(...any) (any, error) `v8:"synchronize"`
  GetBeamPosition func(...any) (any, error) `v8:"getBeamPosition"`
  AllocateSurface func(...any) (any, error) `v8:"allocateSurface"`
  FreeSurface func(...any) (any, error) `v8:"freeSurface"`
  LockSurface func(...any) (any, error) `v8:"lockSurface"`
  UnlockSurface func(...any) (any, error) `v8:"unlockSurface"`
  SwapSurface func(...any) (any, error) `v8:"swapSurface"`
  SetDestination func(...any) (any, error) `v8:"setDestination"`
  GetBlitter func(...any) (any, error) `v8:"getBlitter"`
  WaitComplete func(...any) (any, error) `v8:"waitComplete"`
  GaInterfaceReserved [24]*void `v8:"gaInterfaceReserved"`
}
