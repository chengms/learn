/*
*  构建自己简单的shell
*  author: chengms
*  create：20200619
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main()  {
	fmt.Println("running...")

	// 创建一个输入设备
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
		}
		//fmt.Println(input)

		// handle
		if err = ExecInput(input);err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func ExecInput(input string) error {
	// 移除换行
	input = strings.TrimSuffix(input, "\n")
	// windows 下移除结尾符
	input = strings.TrimSuffix(input, "\r")

	// 拆分参数
	args := strings.Split(input, " ")

	switch args[0] {
	case "":
		// 若是命令为空，则不执行
		return nil
	case "cd":
		// 如果‘cd’后面没有跟目录，不支持
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// 准备解析执行
	cmd := exec.Command(args[0], args[1:]...)

	// 设置输出
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// 执行
	return cmd.Run()
}
