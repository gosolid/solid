//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSLinguisticTagger : objc.NSObject
*/

type NSLinguisticTagger struct {
  *objc.NSObject

}

func (r *NSLinguisticTagger) DominantLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) AvailableTagSchemesForLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TokenRangeAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) SentenceRangeForRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagsInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) DominantLanguageForString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) AvailableTagSchemesForUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagForString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) EnumerateTagsForString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) SetString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) InitWithTagSchemes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) OrthographyAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) EnumerateTagsInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagsForString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagSchemes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) SetOrthography() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) StringEditedInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) PossibleTagsAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) String() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

