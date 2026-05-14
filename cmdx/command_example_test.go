package cmdx

import (
	"context"
	"fmt"
	"time"
)

func ExampleCMDContext() {
	// 创建一个带有 2 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 使用 CMDContext 创建命令，如果超过 2 秒会自动终止
	err := CMDContext(ctx, "sleep", "5").Run()
	if err != nil {
		fmt.Println("Command failed:", err)
		// Output: Command failed: signal: killed
	}
}

func ExampleCMDContextWithOutput() {
	// 创建一个带有 5 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 执行一个正常的命令
	output, err := CMDContext(ctx, "echo", "Hello, World!").Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
	// Output: Hello, World!
}
