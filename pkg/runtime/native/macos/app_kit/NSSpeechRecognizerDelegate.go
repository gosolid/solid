//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSSpeechRecognizerDelegate
*/

type NSSpeechRecognizerDelegate struct {

}

func (r *NSSpeechRecognizerDelegate) SpeechRecognizer() error {
  return fmt.Errorf("unimplemented")
}

