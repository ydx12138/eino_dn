Eino 学习练习
## 记录我个人学习和练习 Eino 的代码仓库。

什么是 Eino?
Eino 是字节跳动 CloudWeGo 团队开源的 Go 语言 LLM 应用开发框架，提供：
组件抽象：ChatModel、Tool、Retriever、Embedding、Prompt 等
编排能力：Chain（链式）、Graph（图）、Workflow（工作流）
流式处理：原生支持 Stream 流式输入输出
生态集成：对接 OpenAI、Ollama、Ark 等模型服务

目录会随练习进度持续更新。

环境准备
Go 1.25+

一个可用的模型服务（OpenAI API / 火山方舟等）

安装依赖
bash
go mod tidy
配置环境变量
在项目根目录创建 .env（已加入 .gitignore）：


说明
本仓库为个人学习用途，代码以练习为主，可能包含试错内容，仅供参考。欢迎交流指正。