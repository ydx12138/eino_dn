package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

//完整生成
/*func main() {
	//加载配置文件
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	//创建一个模型
	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	//调用模型生成响应
	input := []*schema.Message{
		schema.SystemMessage("你是一个可爱的高中美少女"),
		schema.UserMessage("你好"),
	}
	response, err := model.Generate(ctx, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.String())
}*/

// 流式生成
/*func main() {
	//加载配置文件
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	//创建一个模型
	timeout := time.Second * 100
	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("MODEL"),
		Timeout: &timeout,
	})
	//调用模型生成响应
	input := []*schema.Message{
		schema.SystemMessage("你是一个可爱的高中美少女"),
		schema.UserMessage("你好"),
	}
	reader, err := model.Stream(ctx, input)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	//打印流式响应的content
	for {
		resp, err := reader.Recv()
		if err != nil {
			break
		}
		print(resp.Content)
	}
}*/

// ChatTemplate
func main() {
	//加载配置文件
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	//创建一个模型
	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	//调用模型生成响应
	input := []*schema.Message{
		schema.SystemMessage("你是一个可爱的高中美少女"),
		schema.UserMessage("你好"),
	}
	response, err := model.Generate(ctx, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.String())
}
