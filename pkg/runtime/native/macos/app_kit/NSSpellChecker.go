//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSpellChecker : objc.NSObject
*/

type NSSpellChecker struct {
  *objc.NSObject

}

func (r *NSSpellChecker) UniqueSpellDocumentTag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) CheckString() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) ShowInlinePredictionForCandidates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SubstitutionsPanelAccessoryViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) RequestCandidatesForSelectedRange() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) PreventsAutocorrectionBeforeString() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IsAutomaticPeriodSubstitutionEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) UpdateSpellingPanelWithMisspelledWord() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) DeletesAutospaceBetweenString() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) UserReplacementsDictionary() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IsAutomaticInlinePredictionEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) CountWordsInString() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IsAutomaticDashSubstitutionEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) MenuForResult() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IsAutomaticCapitalizationEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) UpdatePanels() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) LearnWord() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SubstitutionsPanel() (*NSPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) AutomaticallyIdentifiesLanguages() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) Language() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) AvailableLanguages() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) UserPreferredLanguages() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) UserQuotesArrayForLanguage() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) CompletionsForPartialWordRange() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SetWordFieldStringValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) UnlearnWord() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) CheckGrammarOfString() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) LanguageForWordRange() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SetLanguage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) RequestCheckingOfString() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SpellingPanel() (*NSPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) AccessoryView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SetSubstitutionsPanelAccessoryViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SetAccessoryView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) UpdateSpellingPanelWithGrammarString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) CorrectionForWordRange() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) HasLearnedWord() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SharedSpellChecker() (*NSSpellChecker, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) CloseSpellDocumentWithTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IsAutomaticSpellingCorrectionEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SetIgnoredWords() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) DismissCorrectionIndicatorForView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SetAutomaticallyIdentifiesLanguages() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IsAutomaticQuoteSubstitutionEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IgnoreWord() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) GuessesForWordRange() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) ShowCorrectionIndicatorOfType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IsAutomaticTextCompletionEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) CheckSpellingOfString() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IgnoredWordsInSpellDocumentWithTag() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) RecordResponse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) SharedSpellCheckerExists() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellChecker) IsAutomaticTextReplacementEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

