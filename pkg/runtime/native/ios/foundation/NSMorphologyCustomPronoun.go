//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSMorphologyCustomPronoun : objc.NSObject
*/

type NSMorphologyCustomPronoun struct {
  *objc.NSObject

}

func (r *NSMorphologyCustomPronoun) SetObjectForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SetPossessiveForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) ReflexiveForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SetReflexiveForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SubjectForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SetSubjectForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) ObjectForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) PossessiveForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) PossessiveAdjectiveForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) SetPossessiveAdjectiveForm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) IsSupportedForLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyCustomPronoun) RequiredKeysForLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

