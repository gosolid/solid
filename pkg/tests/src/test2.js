exports.hello = {
  hello: "hello world",
};

console.info(module.exports);

let delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));
let immediate = () => new Promise((resolve) => setImmediate(resolve));

const main = async () => {
  try {
    const exports = await import("./test2.js");
    console.info(exports);
    let j = 1;
    console.info(Date.now());
    for (let i = 0; i < 1000; i++) {
      await delay(1);
      j = j + i;
    }
    console.info("✅ 1 / 3", Date.now());
    for (let i = 0; i < 100000; i++) {
      await immediate(1);
      j = j + i;
    }
    console.info("✅ 2 / 3", Date.now());
    for (let i = 0; i < 10000000; i++) {
      j = j + i;
    }
    console.info("✅ 3 / 3", Date.now());
  } catch (err) {
    console.error(err);
  }
};

main();
console.info("hello world");
