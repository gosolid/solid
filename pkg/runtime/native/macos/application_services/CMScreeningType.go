//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMScreeningType
*/

type CMScreeningType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  ScreeningFlag uint `v8:"screeningFlag"`
  ChannelCount uint `v8:"channelCount"`
  ChannelInfo [1]CMScreeningChannelRec `v8:"channelInfo"`
}
