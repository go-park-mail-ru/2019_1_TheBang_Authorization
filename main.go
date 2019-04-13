package main

import (
	"BangAuth/config"
	"fmt"
)

func main() {
	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("PORT: %v", config.PORT))

}
