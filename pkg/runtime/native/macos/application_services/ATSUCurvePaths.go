//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUCurvePaths
*/

type ATSUCurvePaths struct {
  Contours uint `v8:"contours"`
  Contour [1]ATSUCurvePath `v8:"contour"`
}
