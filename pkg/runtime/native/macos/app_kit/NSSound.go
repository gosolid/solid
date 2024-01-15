//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSSound : objc.NSObject
*/

type NSSound struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
  *NSPasteboardReading
  *NSPasteboardWriting
}

func (r *NSSound) Pause() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSound) Loops() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSound) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) SetVolume() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSound) SoundNamed() (*NSSound, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) InitWithContentsOfFile() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) InitWithPasteboard() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) Play() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSound) Stop() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSound) PlaybackDeviceIdentifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) CanInitWithPasteboard() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSound) Resume() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSound) SoundUnfilteredTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) Volume() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSound) SetChannelMapping() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSound) Name() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) Duration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSound) CurrentTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSound) SetName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSound) ChannelMapping() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) SetLoops() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSound) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSound) WriteToPasteboard() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSound) SetCurrentTime() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSound) SetPlaybackDeviceIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSound) IsPlaying() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSound) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

