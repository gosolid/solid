//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOATASMARTInterface
*/

type IOATASMARTInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  SMARTEnableDisableOperations func(...any) (any, error) `v8:"sMARTEnableDisableOperations"`
  SMARTEnableDisableAutosave func(...any) (any, error) `v8:"sMARTEnableDisableAutosave"`
  SMARTReturnStatus func(...any) (any, error) `v8:"sMARTReturnStatus"`
  SMARTExecuteOffLineImmediate func(...any) (any, error) `v8:"sMARTExecuteOffLineImmediate"`
  SMARTReadData func(...any) (any, error) `v8:"sMARTReadData"`
  SMARTValidateReadData func(...any) (any, error) `v8:"sMARTValidateReadData"`
  SMARTReadDataThresholds func(...any) (any, error) `v8:"sMARTReadDataThresholds"`
  SMARTReadLogDirectory func(...any) (any, error) `v8:"sMARTReadLogDirectory"`
  SMARTReadLogAtAddress func(...any) (any, error) `v8:"sMARTReadLogAtAddress"`
  SMARTWriteLogAtAddress func(...any) (any, error) `v8:"sMARTWriteLogAtAddress"`
  GetATAIdentifyData func(...any) (any, error) `v8:"getATAIdentifyData"`
}
