// Briefly4U web builder: schema-driven forms + the Go generator compiled to WebAssembly.
// No framework, no network calls except fetching this site's own static files. Nothing is uploaded.
"use strict";
(async function () {
  const root = document.querySelector("main.builder");
  const schema = await (await fetch(root.dataset.schema)).json();

  // 1. Forms: render basics first, advanced collapsed, "add your own" last (three tiers).
  // TODO: tiny renderer for object/enum/boolean/string/array from `schema` ($defs/section, $defs/customEntry).
  document.getElementById("b4u-wizard").textContent = "TODO: wizard steps rendered from " + (schema.title || "schema");

  // 2. Generator: lazy-load the WASM only when the user reaches Review/Export (~9 MB raw, ~2.3 MB gzip).
  async function loadGenerator() {
    const go = new Go(); // from wasm_exec.js
    const { instance } = await WebAssembly.instantiateStreaming(fetch("briefly4u.wasm"), go.importObject);
    go.run(instance); // registers briefly4uValidate / briefly4uBuild on globalThis
  }

  // 3. Export: validate, build, zip in memory, download. Drafts may be kept in localStorage (with a Clear button).
  document.getElementById("b4u-export").addEventListener("click", async () => {
    await loadGenerator();
    // TODO: const yaml = serialise(formState); const v = briefly4uValidate(yaml); const b = briefly4uBuild(yaml, target);
    // TODO: zip b.files (small stored-only zip writer, or archive/zip on the Go side) and trigger a download.
  });
})();
