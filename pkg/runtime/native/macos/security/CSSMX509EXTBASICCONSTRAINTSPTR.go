//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509EXT_BASICCONSTRAINTS_PTR
*/

type CSSMX509EXTBASICCONSTRAINTSPTR struct {
  CA int `v8:"cA"`
  PathLenConstraintPresent int `v8:"pathLenConstraintPresent"`
  PathLenConstraint uint `v8:"pathLenConstraint"`
}
