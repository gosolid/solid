//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MPQueueInfo
*/

type MPQueueInfo struct {
  Version uint `v8:"version"`
  ProcessID *OpaqueMPProcessID `v8:"processID"`
  QueueName uint `v8:"queueName"`
  NWaiting uint64 `v8:"nWaiting"`
  WaitingTaskID *OpaqueMPTaskID `v8:"waitingTaskID"`
  NMessages uint64 `v8:"nMessages"`
  NReserved uint64 `v8:"nReserved"`
  P1 void `v8:"p1"`
  P2 void `v8:"p2"`
  P3 void `v8:"p3"`
}
