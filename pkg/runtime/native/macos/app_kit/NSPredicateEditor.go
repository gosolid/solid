//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPredicateEditor : AppKit.NSRuleEditor
*/

type NSPredicateEditor struct {
  *NSRuleEditor

}

func (r *NSPredicateEditor) SetRowTemplates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditor) RowTemplates() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

