//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSRuleEditorDelegate
*/

type NSRuleEditorDelegate struct {

}

func (r *NSRuleEditorDelegate) RuleEditor() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditorDelegate) RuleEditorRowsDidChange() error {
  return fmt.Errorf("unimplemented")
}

