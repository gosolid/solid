export const toDelimited = (value: number) => {
  let s = Array.from(value.toString()).reverse();

  for (let i = 0; i + 3 < s.length; i += 4) {
    s = [...s.slice(0, i + 3), ',', ...s.slice(i + 3)];
  }
  return s.reverse().join('');
};
