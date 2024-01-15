//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_QUERY_LIMITS_PTR
*/

type CSSMQUERYLIMITSPTR struct {
  TimeLimit uint `v8:"timeLimit"`
  SizeLimit uint `v8:"sizeLimit"`
}
