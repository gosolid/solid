//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIReferenceLibraryViewController : UIKit.UIViewController
*/

type UIReferenceLibraryViewController struct {
  *UIViewController

}

func (r *UIReferenceLibraryViewController) DictionaryHasDefinitionForTerm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIReferenceLibraryViewController) InitWithTerm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIReferenceLibraryViewController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIReferenceLibraryViewController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIReferenceLibraryViewController) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

