//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSPredicateEditorRowTemplate : objc.NSObject
*/

type NSPredicateEditorRowTemplate struct {
  *objc.NSObject
  *foundation.NSCoding
  *foundation.NSCopying
}

func (r *NSPredicateEditorRowTemplate) SetPredicate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) TemplatesWithAttributeKeyPaths() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) Modifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) MatchForPredicate() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) InitWithCompoundTypes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) Options() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) CompoundTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) PredicateWithSubpredicates() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) InitWithLeftExpressions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) TemplateViews() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) RightExpressions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) RightExpressionAttributeType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) Operators() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) DisplayableSubpredicatesOfPredicate() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicateEditorRowTemplate) LeftExpressions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

