package config

import "fmt"

func name(x string) string {
	return fmt.Sprintf("kubernetes-the-hard-way:config:%s", x)
}
