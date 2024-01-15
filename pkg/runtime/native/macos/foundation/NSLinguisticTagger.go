//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSLinguisticTagger : objc.NSObject
*/

type NSLinguisticTagger struct {
  *objc.NSObject

}

func (r *NSLinguisticTagger) AvailableTagSchemesForLanguage() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TokenRangeAtIndex() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagsInRange() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) String() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) SetString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) SetOrthography() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) SentenceRangeForRange() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagsForString() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) EnumerateTagsForString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) InitWithTagSchemes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) AvailableTagSchemesForUnit() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) EnumerateTagsInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) PossibleTagsAtIndex() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagSchemes() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) OrthographyAtIndex() (*NSOrthography, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) StringEditedInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagAtIndex() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) DominantLanguageForString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) TagForString() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLinguisticTagger) DominantLanguage() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

