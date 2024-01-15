//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIInputViewController : UIKit.UIViewController
*/

type UIInputViewController struct {
  *UIViewController

}

func (r *UIInputViewController) SetPrimaryLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) DismissKeyboard() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) AdvanceToNextInputMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) HandleInputModeListFromView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) SetInputView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) NeedsInputModeSwitchKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) InputView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) PrimaryLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) SetHasDictationKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) HasFullAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) RequestSupplementaryLexiconWithCompletion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) TextDocumentProxy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputViewController) HasDictationKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

