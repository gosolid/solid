//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPresentationIntent : objc.NSObject
*/

type NSPresentationIntent struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSPresentationIntent) Column() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ListItemIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) BlockQuoteIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) TableIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) TableCellIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Ordinal() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) IndentationLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) HeaderLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) IsEquivalentToPresentationIntent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ColumnAlignments() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) UnorderedListIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) IntentKind() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ParentIntent() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Row() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) HeaderIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) OrderedListIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) TableRowIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Identity() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ColumnCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ParagraphIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) ThematicBreakIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) TableHeaderRowIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) CodeBlockIntentWithIdentity() (*NSPresentationIntent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPresentationIntent) LanguageHint() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

