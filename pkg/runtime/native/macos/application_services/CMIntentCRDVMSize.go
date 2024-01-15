//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMIntentCRDVMSize
*/

type CMIntentCRDVMSize struct {
  RenderingIntent uint `v8:"renderingIntent"`
  VMSize uint `v8:"vMSize"`
}
