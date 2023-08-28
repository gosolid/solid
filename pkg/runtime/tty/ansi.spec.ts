import { expect } from 'chai';
import { ansi } from 'tty';
const ESC = '\u001b';

enum AnsiCodes {
  reset = 0,

  bold = 1,
  dim = 2,
  italic = 3,
  underline = 4,
  blink = 5,
  reverse = 7,
  hidden = 8,
  strikethrough = 9,

  boldReset = 22,
  dimReset = 22,
  italicReset = 23,
  underlineReset = 24,
  blinkReset = 25,
  reverseReset = 27,
  hiddenReset = 28,
  strikethroughReset = 29,

  black = 30,
  red = 31,
  green = 32,
  yellow = 33,
  blue = 34,
  magenta = 35,
  cyan = 36,
  white = 37,
  default = 39,

  backgroundBlack = 40,
  backgroundRed = 41,
  backgroundGreen = 42,
  backgroundYellow = 43,
  backgroundBlue = 44,
  backgroundMagenta = 45,
  backgroundCyan = 46,
  backgroundWhite = 47,
  backgroundDefault = 49,
}

describe('tty.ANSI', () => {
  Object.entries(AnsiCodes).forEach(([key, value]) => {
    if (typeof value === 'number' && value >= 1 && value <= 10) {
      it(`should return ${key}`, () => {
        expect(ansi[key]('test')).to.equal(
          `${ESC}[${value}mtest${ESC}[${
            AnsiCodes[(key + 'Reset') as any]
          }m${ESC}[${AnsiCodes.reset}m`
        );
      });
    }
  });

  Object.entries(AnsiCodes).forEach(([key, value]) => {
    if (typeof value === 'number' && value >= 30 && value <= 50) {
      it(`should return ${key}`, () => {
        expect(ansi[key]('test')).to.equal(
          `${ESC}[${value}mtest${ESC}[${AnsiCodes.reset}m`
        );
      });
    }
  });

  it('should handle chained formatting', () => {
    expect(ansi.bold.cyan('bold cyan')).to.equal(
      `${ESC}[${AnsiCodes.bold};${AnsiCodes.cyan}mbold cyan${ESC}[${AnsiCodes.boldReset}m${ESC}[${AnsiCodes.reset}m`
    );
  });

  it('should handle embedded formatting', () => {
    expect(
      ansi.cyan(
        'cyan ' +
          ansi.red(
            'red ' + ansi.bold('bold ' + ansi.dim('dim') + ' bold') + ' red'
          ) +
          ' cyan'
      )
    ).to.equal(
      `${ESC}[${AnsiCodes.cyan}mcyan ${ESC}[${AnsiCodes.red}mred ${ESC}[${AnsiCodes.bold}mbold ${ESC}[${AnsiCodes.dim}mdim${ESC}[${AnsiCodes.boldReset}m${ESC}[${AnsiCodes.bold}m bold${ESC}[${AnsiCodes.boldReset}m${ESC}[${AnsiCodes.red}m red${ESC}[${AnsiCodes.cyan}m cyan${ESC}[${AnsiCodes.reset}m`
    );
  });

  it('should strip formatting', () => {
    expect(
      ansi.strip(
        ansi.cyan(
          'cyan ' +
            ansi.red(
              'red ' + ansi.bold('bold ' + ansi.dim('dim') + ' bold') + ' red'
            ) +
            ' cyan'
        )
      )
    ).to.equal('cyan red bold dim bold red cyan');
  });
});
