package remote

import (
	"fmt"
)

func name(x string) string {
	return fmt.Sprintf("kubernetes-the-hard-way:remote:%s", x)
}
