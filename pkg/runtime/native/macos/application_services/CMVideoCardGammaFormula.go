//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMVideoCardGammaFormula
*/

type CMVideoCardGammaFormula struct {
  RedGamma int `v8:"redGamma"`
  RedMin int `v8:"redMin"`
  RedMax int `v8:"redMax"`
  GreenGamma int `v8:"greenGamma"`
  GreenMin int `v8:"greenMin"`
  GreenMax int `v8:"greenMax"`
  BlueGamma int `v8:"blueGamma"`
  BlueMin int `v8:"blueMin"`
  BlueMax int `v8:"blueMax"`
}
