<palign="center"><img src="assets/logo.png" alt="grove Logo" width="200"/></p>

<p align="center">
    <strong>让AI回归简单，像安装Office一样使用AI</strong>
</p>

<p align="center">
    <a href="#核心特性">核心特性</a> •
    <a href="#快速开始">快速开始</a> •
    <a href="#应用场景">应用场景</a> •
    <a href="#技术架构">技术架构</a>•
    <a href="#截图展示">截图展示</a>•
    <a href="#路线图">路线图</a>•
    <a href="#社区与贡献">社区与贡献</a> •
    <a href="#许可协议">许可协议</a>
</p>

## 核心特性

### 🚀 本地优先，一键上手

- 极简安装体验：下载、安装、使用，像普通软件一样简单
- 无需服务器：告别复杂部署，无需配置服务器、域名和SSL证书
- 离线可用：核心功能完全离线，不依赖网络连接

### 🛡️ 隐私至上，数据安全

- 数据不离开设备：所有处理在本地完成，数据绝不上传
- 完全控制：您决定数据存储位置和使用方式
- 开源透明：核心代码完全开源，接受社区审计

### 🧠 强大的本地AI能力

- 多种开源模型：支持Llama、Mistral、Yi、Qwen等主流模型
- 本地知识库：构建基于您文档的专属知识库
- 文本处理：总结、创作、翻译、代码生成等功能

### 🔌 丰富的扩展生态

- 应用商店：安装各类垂直场景的AI应用
- 服务集成：可选接入在线API和服务增强能力
- 跨平台支持：Windows、macOS、Linux全平台支持

## 快速开始

### 安装

1. 从发布页面下载适合您系统的安装包
2. 双击安装包，按照向导完成安装
3. 启动 Grove Studio，进入 AI 助手页面

### 初次使用

1. 启动应用后，可以在"模型管理"中选择下载适合您设备的AI模型
2. 在"服务管理"中启用您需要的本地服务
3. 开始使用内置应用，或从应用商店安装更多功能

## 应用场景

### 💬 智能对话助手

与AI进行自然对话，获取信息、解决问题或头脑风暴

### 📚 本地知识库

- 导入公司文档、研究论文或个人笔记
- 使用自然语言快速查询信息
- 数据不离开本地，确保信息安全

### ✍️内容创作

- 撰写文章、邮件或创意内容
- 润色和改进现有文本
- 多风格多场景的写作辅助

### 📊 数据分析

导入表格数据，进行分析并生成洞见，全程本地处理

## 开发

### 安装依赖

- golang
- npm
- make

### 安装 wails

```shell  
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 开发  
  
```
shell wails dev
```  
  
### 构建  
  
```
shell make all
```

## 技术架构

Grove Studio 采用现代化技术栈，确保性能、安全和用户体验：

- 主应用核心：Golang + Wails + Vue.js + Tailwind CSS
- 本地AI引擎：基于GGML/GGUF的高效推理引擎
- 知识库服务：嵌入式向量数据库支持
- 应用扩展：标准化API接口与SDK

我们的架构设计原则是：

- 极致性能优化，降低硬件要求
- 模块化设计，便于扩展
- 隐私优先，本地处理为主
- 用户体验优先，简化复杂技术

## 截图展示

<p align="center">
    <img src="assets/screenshot-dashboard.png" alt="Dashboard" width="45%"/>
    <img src="assets/screenshot-chat.png" alt="Chat Interface" width="45%"/>
</p>

<p align="center">
    <img src="assets/screenshot-knowledge.png" alt="Knowledge Base" width="45%"/>
    <img src="assets/screenshot-models.png" alt="Model Management" width="45%"/>
</p>

## 路线图

- [x] 核心框架与基础功能
- [x] 本地模型管理与推理
- [x] 基础知识库功能
- [ ] 应用商店生态
- [ ] 更多垂直场景应用
- [ ] 高级知识库功能
- [ ] 多模态支持（图像、音频）
- [ ] 企业级功能与支持

## 社区与贡献

Grove Studio 是一个开源项目，我们欢迎社区贡献：

- 🐛 报告Bug
- 💡 提出新功能
- 🔧 提交PR
- 📚 改进文档
- 🌐 帮助翻译

加入我们的讨论区，分享您的想法和使用体验！

## 支持 Grove Studio

如果您喜欢我们的项目，可以通过以下方式支持我们：

- ⭐ 在GitHub上给我们 Star
- 📢 向朋友推荐 Grove Studio
- 🤝 参与贡献代码或文档
- 💰 成为赞助者

<p align="center">
    <sub>让每个人都能轻松使用AI，无需技术门槛</sub>
</p>

<p align="center">
    <a href="https://github.com/GroveAITeam/grove-studio">GitHub</a> •
    <a href="https://github.com/GroveAITeam/grove-studio">文档</a> •
    <a href="https://github.com/GroveAITeam/grove-studio">博客</a> •
    <a href="mailto:ai.shellphy@gmail.com">联系我们</a>
</p>

