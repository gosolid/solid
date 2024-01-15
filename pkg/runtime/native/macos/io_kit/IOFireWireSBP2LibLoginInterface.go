//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFireWireSBP2LibLoginInterface
*/

type IOFireWireSBP2LibLoginInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  SubmitLogin func(...any) (any, error) `v8:"submitLogin"`
  SubmitLogout func(...any) (any, error) `v8:"submitLogout"`
  SetLoginFlags func(...any) (any, error) `v8:"setLoginFlags"`
  SetLoginCallback func(...any) (any, error) `v8:"setLoginCallback"`
  SetLogoutCallback func(...any) (any, error) `v8:"setLogoutCallback"`
  SetRefCon func(...any) (any, error) `v8:"setRefCon"`
  GetRefCon func(...any) (any, error) `v8:"getRefCon"`
  GetMaxCommandBlockSize func(...any) (any, error) `v8:"getMaxCommandBlockSize"`
  GetLoginID func(...any) (any, error) `v8:"getLoginID"`
  SetMaxPayloadSize func(...any) (any, error) `v8:"setMaxPayloadSize"`
  SetReconnectTime func(...any) (any, error) `v8:"setReconnectTime"`
  CreateORB func(...any) (any, error) `v8:"createORB"`
  SubmitORB func(...any) (any, error) `v8:"submitORB"`
  SetUnsolicitedStatusNotify func(...any) (any, error) `v8:"setUnsolicitedStatusNotify"`
  SetStatusNotify func(...any) (any, error) `v8:"setStatusNotify"`
  SetFetchAgentResetCallback func(...any) (any, error) `v8:"setFetchAgentResetCallback"`
  SubmitFetchAgentReset func(...any) (any, error) `v8:"submitFetchAgentReset"`
  SetFetchAgentWriteCallback func(...any) (any, error) `v8:"setFetchAgentWriteCallback"`
  RingDoorbell func(...any) (any, error) `v8:"ringDoorbell"`
  EnableUnsolicitedStatus func(...any) (any, error) `v8:"enableUnsolicitedStatus"`
  SetBusyTimeoutRegisterValue func(...any) (any, error) `v8:"setBusyTimeoutRegisterValue"`
  SetPassword func(...any) (any, error) `v8:"setPassword"`
}
