//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISymbolEffectCompletionContext : objc.NSObject
*/

type UISymbolEffectCompletionContext struct {
  *objc.NSObject

}

func (r *UISymbolEffectCompletionContext) ContentTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISymbolEffectCompletionContext) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISymbolEffectCompletionContext) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISymbolEffectCompletionContext) IsFinished() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISymbolEffectCompletionContext) Sender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISymbolEffectCompletionContext) Effect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

