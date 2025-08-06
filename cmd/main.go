package main

import (
	editor_api "dido/internal/editor"
)

func main() {
	editor := editor_api.NewEditor()
	defer editor.Close()

	editor.Run()
}
