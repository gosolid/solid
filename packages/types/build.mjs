import glob from "glob";
import fs from "fs";
import path from "path";

const __filename = new URL(import.meta.url).pathname;
const __dirname = path.dirname(__filename);

const goRoot = path.resolve(__dirname, "lib");

const files = glob.sync("**/*.d.ts", {
  cwd: goRoot,
  ignore: path.resolve(__dirname, "lib", "index.d.ts"),
});
fs.mkdirSync(path.resolve(__dirname, "lib"), { recursive: true });
fs.writeFileSync(
  path.resolve(__dirname, "lib", "index.d.ts"),
  files
    .map((file) => {
      const out = `export type * from './${file}';`;
      return out;
    })
    .join("\n")
);
files.forEach((file) => {
  fs.mkdirSync(path.dirname(path.resolve(__dirname, "lib", file)), {
    recursive: true,
  });
  fs.copyFileSync(
    path.resolve(goRoot, file),
    path.resolve(__dirname, "lib", file)
  );
});

const packages = [...new Set(files.map((file) => path.dirname(file)))];
packages.forEach((pkgName) => {
  fs.writeFileSync(
    path.resolve(__dirname, "lib", pkgName, "__DOCS__.d.ts"),
    files
      .filter((x) => x.startsWith(pkgName + "/"))
      .map((file) =>
        fs.readFileSync(path.resolve(__dirname, "lib", file)).toString()
      )
      .join("\n")
  );
});
