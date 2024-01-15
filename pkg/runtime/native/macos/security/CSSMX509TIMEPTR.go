//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_TIME_PTR
*/

type CSSMX509TIMEPTR struct {
  TimeType byte `v8:"timeType"`
  Time SecAsn1Oid `v8:"time"`
}
