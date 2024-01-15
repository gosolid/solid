//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSITaskInterface
*/

type SCSITaskInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  IsTaskActive func(...any) (any, error) `v8:"isTaskActive"`
  SetTaskAttribute func(...any) (any, error) `v8:"setTaskAttribute"`
  GetTaskAttribute func(...any) (any, error) `v8:"getTaskAttribute"`
  SetCommandDescriptorBlock func(...any) (any, error) `v8:"setCommandDescriptorBlock"`
  GetCommandDescriptorBlockSize func(...any) (any, error) `v8:"getCommandDescriptorBlockSize"`
  GetCommandDescriptorBlock func(...any) (any, error) `v8:"getCommandDescriptorBlock"`
  SetScatterGatherEntries func(...any) (any, error) `v8:"setScatterGatherEntries"`
  SetTimeoutDuration func(...any) (any, error) `v8:"setTimeoutDuration"`
  GetTimeoutDuration func(...any) (any, error) `v8:"getTimeoutDuration"`
  SetTaskCompletionCallback func(...any) (any, error) `v8:"setTaskCompletionCallback"`
  ExecuteTaskAsync func(...any) (any, error) `v8:"executeTaskAsync"`
  ExecuteTaskSync func(...any) (any, error) `v8:"executeTaskSync"`
  AbortTask func(...any) (any, error) `v8:"abortTask"`
  GetSCSIServiceResponse func(...any) (any, error) `v8:"getSCSIServiceResponse"`
  GetTaskState func(...any) (any, error) `v8:"getTaskState"`
  GetTaskStatus func(...any) (any, error) `v8:"getTaskStatus"`
  GetRealizedDataTransferCount func(...any) (any, error) `v8:"getRealizedDataTransferCount"`
  GetAutoSenseData func(...any) (any, error) `v8:"getAutoSenseData"`
  SetAutoSenseDataBuffer func(...any) (any, error) `v8:"setAutoSenseDataBuffer"`
  ResetForNewTask func(...any) (any, error) `v8:"resetForNewTask"`
}
