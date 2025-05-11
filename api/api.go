package api

import "fmt"

func Login() {
	fmt.Println("Login")
}

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

func CatalogueSummary() {
	// GET
	fmt.Println("Catalogue summary")
}

func CatalogueQuery() {
	// POST
	fmt.Println("Catalogue Query")
}

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
