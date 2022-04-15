package exp

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	CorsHandle := []string{viper.GetString("CORS.AllowOrigins")}
	fmt.Println(CorsHandle)
}
