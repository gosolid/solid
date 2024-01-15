//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSPresentationIntent : objc.NSObject
*/

type NSPresentationIntent struct {
  *objc.NSObject

}

func (r *NSPresentationIntent) TableRowIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ListItemIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) TableCellIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) CodeBlockIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ColumnCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ThematicBreakIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) OrderedListIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) IsEquivalentToPresentationIntent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Ordinal() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ColumnAlignments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Column() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Row() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ParentIntent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) HeaderIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) BlockQuoteIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) IntentKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ParagraphIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Identity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) UnorderedListIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) TableIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) TableHeaderRowIntentWithIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) HeaderLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) LanguageHint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) IndentationLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

