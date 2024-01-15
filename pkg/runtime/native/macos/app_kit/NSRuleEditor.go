//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSRuleEditor : AppKit.NSControl
*/

type NSRuleEditor struct {
  *NSControl

}

func (r *NSRuleEditor) ParentRowForRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) FormattingDictionary() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) Predicate() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetRowTypeKeyPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SubrowsKeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) ReloadCriteria() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) PredicateForRow() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) RowHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) NumberOfRows() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SelectedRowIndexes() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) RowTypeKeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetFormattingDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetRowHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) ReloadPredicate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) CriteriaForRow() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) RowForDisplayValue() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) RowTypeForRow() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) AddRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) RemoveRowsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) DisplayValuesKeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) CriteriaKeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SubrowIndexesForRow() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) DisplayValuesForRow() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetCriteria() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetFormattingStringsFilename() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetNestingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) CanRemoveAllRows() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) RowClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetCriteriaKeyPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetSubrowsKeyPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) RemoveRowAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SelectRowIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) NestingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetCanRemoveAllRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetRowClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) InsertRowAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) FormattingStringsFilename() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRuleEditor) SetDisplayValuesKeyPath() error {
  return fmt.Errorf("unimplemented")
}

