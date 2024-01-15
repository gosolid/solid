//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSSpeechSynthesizer : objc.NSObject
*/

type NSSpeechSynthesizer struct {
  *objc.NSObject

}

func (r *NSSpeechSynthesizer) Rate() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) IsAnyApplicationSpeaking() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) PauseSpeakingAtBoundary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) ContinueSpeaking() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) UsesFeedbackWindow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) DefaultVoice() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) IsSpeaking() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) AvailableVoices() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) SetVoice() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) PhonemesFromText() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) StopSpeakingAtBoundary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) SetUsesFeedbackWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) SetObject() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) AddSpeechDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) ObjectForProperty() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) SetVolume() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) InitWithVoice() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) StartSpeakingString() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) SetRate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) Voice() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) AttributesForVoice() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) StopSpeaking() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechSynthesizer) Volume() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

