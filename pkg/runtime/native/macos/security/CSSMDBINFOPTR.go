//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DBINFO_PTR
*/

type CSSMDBINFOPTR struct {
  NumberOfRecordTypes uint `v8:"numberOfRecordTypes"`
  DefaultParsingModules *CSSMDBPARSINGMODULEINFOPTR `v8:"defaultParsingModules"`
  RecordAttributeNames *CSSMDBRECORDATTRIBUTEINFOPTR `v8:"recordAttributeNames"`
  RecordIndexes *CSSMDBRECORDINDEXINFOPTR `v8:"recordIndexes"`
  IsLocal int `v8:"isLocal"`
  AccessPath byte `v8:"accessPath"`
  Reserved void `v8:"reserved"`
}
