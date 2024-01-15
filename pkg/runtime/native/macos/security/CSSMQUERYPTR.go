//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_QUERY_PTR
*/

type CSSMQUERYPTR struct {
  RecordType uint `v8:"recordType"`
  Conjunctive uint `v8:"conjunctive"`
  NumSelectionPredicates uint `v8:"numSelectionPredicates"`
  SelectionPredicate *CSSMSELECTIONPREDICATEPTR `v8:"selectionPredicate"`
  QueryLimits CSSMQUERYLIMITSPTR `v8:"queryLimits"`
  QueryFlags uint `v8:"queryFlags"`
}
