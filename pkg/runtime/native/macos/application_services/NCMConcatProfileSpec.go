//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.NCMConcatProfileSpec
*/

type NCMConcatProfileSpec struct {
  RenderingIntent uint `v8:"renderingIntent"`
  TransformTag uint `v8:"transformTag"`
  Profile *OpaqueCMProfileRef `v8:"profile"`
}
