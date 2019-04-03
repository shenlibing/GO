package main
import (
	"go_code/account/utils"
	"fmt"
)
func main() {
	fmt.Printlln("面向对象的方式来完成.....")
	utils.NewMyFamilyAccount().MainMenu()
}