package main

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

	//参数设置
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		schema.UserMessage("请帮帮我，史瓦罗先生，{task}"),
	)

	param := map[string]interface{}{
		"role": "机器人史瓦罗先生",
		"task": "帮我生成一首五言绝句",
	}
	message, err := template.Format(ctx, param)
	if err != nil {
		return
	}
	//调用模型生成响应
	response, err := model.Generate(ctx, message)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Content)
}*/

// 文本转为向量
/*func main() {
	EmedText()
}
func EmedText() {
	//加载配置文件
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	//创建model
	//使用的模型是多模的，因此这里必须为MultiModal
	apitype := emark.APITypeMultiModal
	model, err := emark.NewEmbedder(ctx, &emark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		APIType: &apitype,
	})
	if err != nil {
		panic(err)
	}
	text := []string{
		"这",
		"这",
	}
	//向量化
	embedStrings, err := model.EmbedStrings(ctx, text)
	if err != nil {
		panic(err)
	}
	for i, embedString := range embedStrings {
		fmt.Println("文本", i, ":", "维度:", len(embedString))
	}
}*/
