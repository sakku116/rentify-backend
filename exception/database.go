package exception

import "fmt"

var (
	DbObjNotFound = fmt.Errorf("object not found")
)
