//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.EvGlobals
*/

type EvGlobals struct {
  CursorSema int `v8:"cursorSema"`
  ENum int `v8:"eNum"`
  Buttons int `v8:"buttons"`
  EventFlags int `v8:"eventFlags"`
  VertRetraceClock int `v8:"vertRetraceClock"`
  CursorLoc IOGPoint `v8:"cursorLoc"`
  Frame int `v8:"frame"`
  WorkBounds IOGBounds `v8:"workBounds"`
  MouseRect IOGBounds `v8:"mouseRect"`
  Version int `v8:"version"`
  StructSize int `v8:"structSize"`
  LastFrame int `v8:"lastFrame"`
  ScreenCursorFixed IOFixedPoint32 `v8:"screenCursorFixed"`
  DesktopCursorFixed IOFixedPoint32 `v8:"desktopCursorFixed"`
  ReservedA [27]uint `v8:"reservedA"`
  Reserved uint `v8:"reserved"`
  UpdateCursorPositionFromFixed uint `v8:"updateCursorPositionFromFixed"`
  LogCursorUpdates uint `v8:"logCursorUpdates"`
  WantPressure uint `v8:"wantPressure"`
  WantPrecision uint `v8:"wantPrecision"`
  DontWantCoalesce uint `v8:"dontWantCoalesce"`
  DontCoalesce uint `v8:"dontCoalesce"`
  MouseRectValid uint `v8:"mouseRectValid"`
  MovedMask int `v8:"movedMask"`
  WaitCursorSema int `v8:"waitCursorSema"`
  AALastEventSent int `v8:"aALastEventSent"`
  AALastEventConsumed int `v8:"aALastEventConsumed"`
  WaitCursorUp int `v8:"waitCursorUp"`
  CtxtTimedOut byte `v8:"ctxtTimedOut"`
  WaitCursorEnabled byte `v8:"waitCursorEnabled"`
  GlobalWaitCursorEnabled byte `v8:"globalWaitCursorEnabled"`
  WaitThreshold int `v8:"waitThreshold"`
  LLEHead int `v8:"lLEHead"`
  LLETail int `v8:"lLETail"`
  LLELast int `v8:"lLELast"`
  Lleq [240]NXEQElement `v8:"lleq"`
}
