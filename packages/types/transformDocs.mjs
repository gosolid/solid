import glob from "glob";
import fs from "fs";
import path from "path";

const __filename = new URL(import.meta.url).pathname;
const __dirname = path.dirname(__filename);

const json = JSON.parse(
  fs.readFileSync(path.resolve(__dirname, "..", "docs", "js.json"))
);

const transformNode = (node, file = node) => {
  node.name = node.name?.replace(/\/__DOCS__$/, "");
  node.children?.forEach((node) => transformNode(node, file));
  node.signatures?.forEach((node) => transformNode(node, file));
  if (node.sources) {
    const fileName = node.comment?.blockTags.find(
      ({ tag }) => tag == "@filename"
    );
    const line = node.comment?.blockTags.find(({ tag }) => tag == "@line");
    const character = node.comment?.blockTags.find(
      ({ tag }) => tag == "@column"
    );

    if (fileName) {
      node.comment.blockTags.splice(
        node.comment.blockTags.indexOf(fileName),
        1
      );
    }
    if (line) {
      node.comment.blockTags.splice(node.comment.blockTags.indexOf(line), 1);
    }
    if (character) {
      node.comment.blockTags.splice(
        node.comment.blockTags.indexOf(character),
        1
      );
    }

    if (fileName) {
      file.symbolIdMap[node.id].sourceFileName = fileName?.content[0]?.text;
      node.sources.unshift({
        fileName: fileName?.content[0]?.text,
        line: line?.content[0]?.text,
        character: character?.content[0]?.text,
      });
      node.sources.splice(1, node.sources.length);
    } else if (
      node.signatures?.length > 0 &&
      node.signatures[0]?.sources?.length > 0
    ) {
      node.sources = node.signatures[0].sources;
      file.symbolIdMap[node.id].sourceFileName = node.sources[0].fileName;
    } else {
      const childSources = node.children?.reduce(
        (a, child) => [...a, ...child.sources],
        []
      );
      if (childSources) {
        node.sources = childSources;
      } else {
        node.sources = [];
      }
    }
  }

  if (node.symbolIdMap) {
    Object.keys(node.symbolIdMap)
      .filter((key) =>
        node.symbolIdMap[key].sourceFileName.endsWith("__DOCS__.d.ts")
      )
      .forEach((key) => {
        delete node.symbolIdMap[key].sourceFileName;
      });
  }
};

transformNode(json);

fs.writeFileSync(
  path.resolve(__dirname, "..", "docs", "js.json"),
  JSON.stringify(json, null, 2)
);
