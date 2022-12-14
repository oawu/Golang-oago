/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2022
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package main

import (
	"oago/iosAppIcon"
	"oago/lib"
	"oago/thumbnail"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 { // 沒有指令
		lib.ShowHelp()
		return
	}

	// 取出指令與餐數
	command := args[0]
	arguments := args[1:]

	switch command {
	case "thumbnail":
		thumbnail.Run(arguments)
	case "ios-app-icon":
		iosAppIcon.Run(arguments)
	default:
		lib.ShowHelp()
	}
}
