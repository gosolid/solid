//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_MANAGER_EVENT_NOTIFICATION_PTR
*/

type CSSMMANAGEREVENTNOTIFICATIONPTR struct {
  DestinationModuleManagerType uint `v8:"destinationModuleManagerType"`
  SourceModuleManagerType uint `v8:"sourceModuleManagerType"`
  Event uint `v8:"event"`
  EventId uint `v8:"eventId"`
  EventData SecAsn1Oid `v8:"eventData"`
}
