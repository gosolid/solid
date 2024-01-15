//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSMorphologyCustomPronoun : objc.NSObject
*/

type NSMorphologyCustomPronoun struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSMorphologyCustomPronoun) SetPossessiveForm() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SetPossessiveAdjectiveForm() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) IsSupportedForLanguage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) ObjectForm() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SetObjectForm() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) PossessiveForm() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) PossessiveAdjectiveForm() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) ReflexiveForm() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SetReflexiveForm() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) RequiredKeysForLanguage() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SubjectForm() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SetSubjectForm() error {
  return fmt.Errorf("unimplemented")
}

