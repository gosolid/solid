//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextInputTraits
*/

type NSTextInputTraits struct {

}

func (r *NSTextInputTraits) TextReplacementType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) TextCompletionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetTextCompletionType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetSpellCheckingType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) GrammarCheckingType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SmartQuotesType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SmartDashesType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SmartInsertDeleteType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) LinkDetectionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetLinkDetectionType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) AutocorrectionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SpellCheckingType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetGrammarCheckingType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetSmartInsertDeleteType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetInlinePredictionType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetAutocorrectionType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetSmartQuotesType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetSmartDashesType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetTextReplacementType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) DataDetectionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) SetDataDetectionType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputTraits) InlinePredictionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

