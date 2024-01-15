//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFireWireAVCLibConsumerInterface
*/

type IOFireWireAVCLibConsumerInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  SetSubunit func(...any) (any, error) `v8:"setSubunit"`
  SetRemotePlug func(...any) (any, error) `v8:"setRemotePlug"`
  ConnectToRemotePlug func(...any) (any, error) `v8:"connectToRemotePlug"`
  DisconnectFromRemotePlug func(...any) (any, error) `v8:"disconnectFromRemotePlug"`
  SetFrameStatusHandler func(...any) (any, error) `v8:"setFrameStatusHandler"`
  FrameProcessed func(...any) (any, error) `v8:"frameProcessed"`
  SetMaxPayloadSize func(...any) (any, error) `v8:"setMaxPayloadSize"`
  SetSegmentSize func(...any) (any, error) `v8:"setSegmentSize"`
  GetSegmentSize func(...any) (any, error) `v8:"getSegmentSize"`
  GetSegmentBuffer func(...any) (any, error) `v8:"getSegmentBuffer"`
  SetPortStateHandler func(...any) (any, error) `v8:"setPortStateHandler"`
  SetPortFlags func(...any) (any, error) `v8:"setPortFlags"`
  ClearPortFlags func(...any) (any, error) `v8:"clearPortFlags"`
  GetPortFlags func(...any) (any, error) `v8:"getPortFlags"`
}
