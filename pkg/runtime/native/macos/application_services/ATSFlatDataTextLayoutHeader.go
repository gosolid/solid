//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataTextLayoutHeader
*/

type ATSFlatDataTextLayoutHeader struct {
  NumFlattenedTextLayouts uint `v8:"numFlattenedTextLayouts"`
  FlattenedTextLayouts [1]ATSFlatDataTextLayoutDataHeader `v8:"flattenedTextLayouts"`
}
