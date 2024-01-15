//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSSpeechSynthesizerDelegate
*/

type NSSpeechSynthesizerDelegate struct {

}

func (r *NSSpeechSynthesizerDelegate) SpeechSynthesizer() error {
  return fmt.Errorf("unimplemented")
}

