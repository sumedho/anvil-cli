package api

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// func Login() {
// 	fmt.Println("Login")
// }

func ObjectAttributes() {
	// GET
	fmt.Println("Anvil object attributes")
}

func ObjectAttributesChange() {
	//POST
	fmt.Println("Anvil object attr change")
}

func ObjectInfo() {
	// GET
	fmt.Println("Anvil object info")
}

func ObjectFiles() {
	// GET
	fmt.Println("Get object files")
}

// func CatalogueSummary() {
// 	// GET
// 	fmt.Println("Catalogue summary")
// }

func CatalogueQueryAttributes() {
	// GET
	fmt.Println("Get catalogue attributes")
}

func CataloguePrefixesGet() {
	// GET
	fmt.Println("Catalogue Prefix list")
}

func CataloguePrefixPost() {
	// POST
	fmt.Println("Create new Catalogue prefix")
}

func MagmaActive() {
	// GET
	fmt.Println("Current active Magma sessions")
}

func ListWorkflowSchedules(cCtx cli.Context) {
	catalogue_id := cCtx.String("id")
	fmt.Println("Workflow schedules", catalogue_id)
}
