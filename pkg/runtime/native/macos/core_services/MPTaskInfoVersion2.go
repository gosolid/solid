//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.MPTaskInfoVersion2
*/

type MPTaskInfoVersion2 struct {
  Version uint `v8:"version"`
  Name uint `v8:"name"`
  QueueName uint `v8:"queueName"`
  RunState uint16 `v8:"runState"`
  LastCPU uint16 `v8:"lastCPU"`
  Weight uint `v8:"weight"`
  ProcessID *OpaqueMPProcessID `v8:"processID"`
  CpuTime objc.UnsignedWide `v8:"cpuTime"`
  SchedTime objc.UnsignedWide `v8:"schedTime"`
  CreationTime objc.UnsignedWide `v8:"creationTime"`
  CodePageFaults uint64 `v8:"codePageFaults"`
  DataPageFaults uint64 `v8:"dataPageFaults"`
  Preemptions uint64 `v8:"preemptions"`
  CpuID *OpaqueMPCpuID `v8:"cpuID"`
}
