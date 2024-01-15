//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UITextChecker : objc.NSObject
*/

type UITextChecker struct {
  *objc.NSObject

}

func (r *UITextChecker) LearnWord() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) HasLearnedWord() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) AvailableLanguages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) IgnoredWords() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) SetIgnoredWords() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) RangeOfMisspelledWordInString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) GuessesForWordRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) CompletionsForPartialWordRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) IgnoreWord() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextChecker) UnlearnWord() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

