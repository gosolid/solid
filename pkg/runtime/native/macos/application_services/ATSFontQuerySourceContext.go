//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFontQuerySourceContext
*/

type ATSFontQuerySourceContext struct {
  Version uint `v8:"version"`
  RefCon void `v8:"refCon"`
  Retain *func(...any) (any, error) `v8:"retain"`
  Release *func(...any) (any, error) `v8:"release"`
}
