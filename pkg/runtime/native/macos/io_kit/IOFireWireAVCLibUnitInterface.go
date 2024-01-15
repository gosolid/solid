//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFireWireAVCLibUnitInterface
*/

type IOFireWireAVCLibUnitInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  Open func(...any) (any, error) `v8:"open"`
  OpenWithSessionRef func(...any) (any, error) `v8:"openWithSessionRef"`
  GetSessionRef func(...any) (any, error) `v8:"getSessionRef"`
  Close func(...any) (any, error) `v8:"close"`
  AddCallbackDispatcherToRunLoop func(...any) (any, error) `v8:"addCallbackDispatcherToRunLoop"`
  RemoveCallbackDispatcherFromRunLoop func(...any) (any, error) `v8:"removeCallbackDispatcherFromRunLoop"`
  SetMessageCallback func(...any) (any, error) `v8:"setMessageCallback"`
  AVCCommand func(...any) (any, error) `v8:"aVCCommand"`
  AVCCommandInGeneration func(...any) (any, error) `v8:"aVCCommandInGeneration"`
  GetAncestorInterface func(...any) (any, error) `v8:"getAncestorInterface"`
  GetProtocolInterface func(...any) (any, error) `v8:"getProtocolInterface"`
  GetAsyncConnectionPlugCounts func(...any) (any, error) `v8:"getAsyncConnectionPlugCounts"`
  CreateConsumerPlug func(...any) (any, error) `v8:"createConsumerPlug"`
  UpdateAVCCommandTimeout func(...any) (any, error) `v8:"updateAVCCommandTimeout"`
  MakeP2PInputConnection func(...any) (any, error) `v8:"makeP2PInputConnection"`
  BreakP2PInputConnection func(...any) (any, error) `v8:"breakP2PInputConnection"`
  MakeP2POutputConnection func(...any) (any, error) `v8:"makeP2POutputConnection"`
  BreakP2POutputConnection func(...any) (any, error) `v8:"breakP2POutputConnection"`
  CreateAVCAsynchronousCommand func(...any) (any, error) `v8:"createAVCAsynchronousCommand"`
  AVCAsynchronousCommandSubmit func(...any) (any, error) `v8:"aVCAsynchronousCommandSubmit"`
  AVCAsynchronousCommandReinit func(...any) (any, error) `v8:"aVCAsynchronousCommandReinit"`
  AVCAsynchronousCommandCancel func(...any) (any, error) `v8:"aVCAsynchronousCommandCancel"`
  AVCAsynchronousCommandRelease func(...any) (any, error) `v8:"aVCAsynchronousCommandRelease"`
  AVCAsynchronousCommandReinitWithCommandBytes func(...any) (any, error) `v8:"aVCAsynchronousCommandReinitWithCommandBytes"`
}
