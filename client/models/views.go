package Models

import (
	"fmt"
)

func (m OriginalModel) Views() string {
	welcomeMsg := "Welcome to Echo Chat!\n"
	return fmt.Sprintf("%s\n", welcomeMsg)
}
