//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_PARSING_MODULE_INFO_PTR
*/

type CSSMDBPARSINGMODULEINFOPTR struct {
  RecordType uint `v8:"recordType"`
  ModuleSubserviceUid CSSMSUBSERVICEUIDPTR `v8:"moduleSubserviceUid"`
}
