/*
 * @Author: zhangniannian
 * @Date: 2022-04-08 10:45:06
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-04-08 11:09:18
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"os"

	"bufio"

	"ige/functional/fib"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			// Uncomment panic to see
			// how it works with defer
			// panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
