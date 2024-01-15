//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataTextLayoutDataHeader
*/

type ATSFlatDataTextLayoutDataHeader struct {
  SizeOfLayoutData uint `v8:"sizeOfLayoutData"`
  TextLayoutLength uint `v8:"textLayoutLength"`
  OffsetToLayoutControls uint `v8:"offsetToLayoutControls"`
  OffsetToLineInfo uint `v8:"offsetToLineInfo"`
}
