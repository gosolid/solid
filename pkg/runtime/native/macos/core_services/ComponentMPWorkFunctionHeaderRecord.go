//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ComponentMPWorkFunctionHeaderRecord
*/

type ComponentMPWorkFunctionHeaderRecord struct {
  HeaderSize uint `v8:"headerSize"`
  RecordSize uint `v8:"recordSize"`
  WorkFlags uint `v8:"workFlags"`
  ProcessorCount uint16 `v8:"processorCount"`
  Unused byte `v8:"unused"`
  IsRunning byte `v8:"isRunning"`
}
