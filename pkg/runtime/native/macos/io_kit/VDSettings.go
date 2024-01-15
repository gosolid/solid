//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDSettings
*/

type VDSettings struct {
  CsParamCnt int16 `v8:"csParamCnt"`
  CsBrightMax int16 `v8:"csBrightMax"`
  CsBrightDef int16 `v8:"csBrightDef"`
  CsBrightVal int16 `v8:"csBrightVal"`
  CsCntrstMax int16 `v8:"csCntrstMax"`
  CsCntrstDef int16 `v8:"csCntrstDef"`
  CsCntrstVal int16 `v8:"csCntrstVal"`
  CsTintMax int16 `v8:"csTintMax"`
  CsTintDef int16 `v8:"csTintDef"`
  CsTintVal int16 `v8:"csTintVal"`
  CsHueMax int16 `v8:"csHueMax"`
  CsHueDef int16 `v8:"csHueDef"`
  CsHueVal int16 `v8:"csHueVal"`
  CsHorizDef int16 `v8:"csHorizDef"`
  CsHorizVal int16 `v8:"csHorizVal"`
  CsHorizMax int16 `v8:"csHorizMax"`
  CsVertDef int16 `v8:"csVertDef"`
  CsVertVal int16 `v8:"csVertVal"`
  CsVertMax int16 `v8:"csVertMax"`
}
