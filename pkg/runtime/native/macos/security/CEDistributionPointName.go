//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_DistributionPointName
*/

type CEDistributionPointName struct {
  NameType int `v8:"nameType"`
  Dpn void `v8:"dpn"`
}
