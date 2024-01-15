//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MPNotificationInfo
*/

type MPNotificationInfo struct {
  Version uint `v8:"version"`
  ProcessID *OpaqueMPProcessID `v8:"processID"`
  NotificationName uint `v8:"notificationName"`
  QueueID *OpaqueMPQueueID `v8:"queueID"`
  P1 void `v8:"p1"`
  P2 void `v8:"p2"`
  P3 void `v8:"p3"`
  EventID *OpaqueMPEventID `v8:"eventID"`
  Events uint `v8:"events"`
  SemaphoreID *OpaqueMPSemaphoreID `v8:"semaphoreID"`
}
