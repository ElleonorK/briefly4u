//go:build js && wasm

// Command briefly4u-wasm exposes the generator to the static web builder.
// Build: GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o site/static/builder/briefly4u.wasm ./cmd/briefly4u-wasm
// It registers three JS functions on globalThis: briefly4uValidate, briefly4uBuild, briefly4uVersion.
package main

import (
	"syscall/js"

	"github.com/ElleonorK/briefly4u/internal/generator"
	"github.com/ElleonorK/briefly4u/internal/library"
	"github.com/ElleonorK/briefly4u/internal/profile"
)

func main() {
	js.Global().Set("briefly4uVersion", js.FuncOf(func(this js.Value, args []js.Value) any { return "dev" }))

	// briefly4uValidate(profileYAML string) -> {ok: bool, errors: string[]}
	js.Global().Set("briefly4uValidate", js.FuncOf(func(this js.Value, args []js.Value) any {
		_, errs := profile.Parse([]byte(args[0].String()))
		out := make([]any, 0, len(errs))
		for _, e := range errs {
			out = append(out, e.Error())
		}
		return map[string]any{"ok": len(errs) == 0, "errors": out}
	}))

	// briefly4uBuild(profileYAML string, target string) -> {files: {path: content}, error?: string}
	// The page zips the files in JS (TODO: or zip in Go with archive/zip and return a Uint8Array).
	js.Global().Set("briefly4uBuild", js.FuncOf(func(this js.Value, args []js.Value) any {
		p, errs := profile.Parse([]byte(args[0].String()))
		if len(errs) > 0 {
			return map[string]any{"error": errs[0].Error()}
		}
		b, err := generator.Build(p, generator.Options{Target: args[1].String(), Library: library.Embedded()})
		if err != nil {
			return map[string]any{"error": err.Error()}
		}
		files := map[string]any{}
		for path, f := range b.Files {
			files[path] = string(f.Content)
		}
		return map[string]any{"files": files}
	}))

	select {} // keep the Go runtime alive for JS callbacks
}
