//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_SemanticsInformation
*/

type CESemanticsInformation struct {
  SemanticsIdentifier SecAsn1Oid `v8:"semanticsIdentifier"`
  NameRegistrationAuthorities CEGeneralNames `v8:"nameRegistrationAuthorities"`
}
