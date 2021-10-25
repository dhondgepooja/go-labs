package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public (or exported)" // NOTE the start letter is UPERCASE
var PackageVar = "PackageVar"

func notExported() {

}

func Exported() {

}

func PrintMe(one, two string) {
	fmt.Println(one, two, PackageVar)
}
