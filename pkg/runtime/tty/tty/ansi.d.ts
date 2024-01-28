/**
 * @moduleName tty
@packageDescription
@module tty*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 3 @column 0 */
declare module "tty" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
  class ANSI {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 37 @column 0 */
  enable(  ): ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 44 @column 0 */
  disable(  ): ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 51 @column 0 */
  home(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 62 @column 0 */
  strip(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 67 @column 0 */
  move(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 76 @column 0 */
  moveUp(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 85 @column 0 */
  moveDown(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 94 @column 0 */
  moveRight(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 103 @column 0 */
  moveLeft(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 112 @column 0 */
  moveToStartDown(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 121 @column 0 */
  moveToStartUp(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 130 @column 0 */
  moveToStart(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 139 @column 0 */
  requestCursorPosition(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 148 @column 0 */
  moveUpOne(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 157 @column 0 */
  saveCursorPosition(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 166 @column 0 */
  restoreCursorPosition(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 175 @column 0 */
  saveCursorPositionSCO(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 184 @column 0 */
  requestCursorPositionSCO(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 193 @column 0 */
  eraseInDisplay(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 202 @column 0 */
  eraseFromCursorToEndOfScreen(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 211 @column 0 */
  eraseFromCursorToStartOfScreen(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 220 @column 0 */
  eraseEntireScreen(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 229 @column 0 */
  eraseSavedLines(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 238 @column 0 */
  eraseInLine(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 247 @column 0 */
  eraseFromCursorToEndOfLine(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 256 @column 0 */
  eraseFromCursorToStartOfLine(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 265 @column 0 */
  eraseLine(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 589 @column 0 */
  indexedColor(  ): ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 595 @column 0 */
  backgroundIndexedColor(  ): ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 601 @column 0 */
  color(  ): ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 607 @column 0 */
  backgroundColor(  ): ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly brightGreen: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBrightGreen: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly brightYellow: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly reverseReset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBlack: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBrightBlack: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly hiddenReset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly brightBlack: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBrightBlue: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly brightMagenta: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly brightWhite: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly reset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly boldReset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly dimReset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundGreen: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly brightBlue: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly magenta: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly white: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly default: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly italicReset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly strikethrough: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundRed: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly blinkReset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly brightCyan: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly black: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly brightRed: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly cyan: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBrightCyan: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundWhite: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly underlineReset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly red: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly blue: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBrightMagenta: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundDefault: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBrightRed: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly strikethroughReset: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBrightWhite: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly yellow: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundYellow: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly blink: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundMagenta: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundCyan: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly enabled: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly dim: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly underline: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBlue: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly bold: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly hidden: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly backgroundBrightYellow: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly italic: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly reverse: ANSI;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/ansi.go @line 24 @column 0 */
    readonly green: ANSI;
  }
}
