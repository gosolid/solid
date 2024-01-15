import { DefaultTheme } from "typedoc";

export function load(app) {
  app.renderer.defineTheme("@solide/typedoc-theme", DefaultTheme);
}
