/*
 * @Author: zhangniannian
 * @Date: 2022-04-08 10:45:06
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-04-08 11:10:30
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	rpcdemo "ige/rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div",
		rpcdemo.Args{10, 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div",
		rpcdemo.Args{10, 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
