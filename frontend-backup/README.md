# Grove Studio 前端项目

- [Grove Studio 前端项目](#grove-studio-前端项目)
  - [⚙️ 环境要求](#️-环境要求)
  - [📂 项目目录结构](#-项目目录结构)
  - [📦 主要依赖包介绍](#-主要依赖包介绍)
    - [核心依赖](#核心依赖)
    - [UI相关](#ui相关)
    - [网络请求](#网络请求)
    - [开发工具](#开发工具)
  - [🚀 开发命令](#-开发命令)
  - [📚 相关文档](#-相关文档)

## ⚙️ 环境要求

- **Node.js版本**: ^18.18.0 || ^20.9.0 || >=21.1.0
- **包管理器**: pnpm >= 7.33.7

## 📂 项目目录结构

```
├── public/              # 静态资源目录
├── src/                 # 源代码目录
│   ├── api/             # API接口定义和请求封装
│   ├── assets/          # 项目资源文件(图片、字体等)
│   ├── components/      # 公共组件
│   │   └── ui/          # UI组件库
│   ├── lib/             # 工具库
│   ├── router/          # 路由配置
│   ├── store/           # 状态管理
│   ├── types/           # TypeScript类型定义
│   ├── utils/           # 工具函数
│   ├── views/           # 页面视图组件
│   ├── App.vue          # 根组件
│   ├── main.ts          # 应用入口文件
│   ├── style.css        # 全局样式
│   └── index.scss       # SCSS样式
├── .gitignore           # Git忽略文件
├── components.json      # 组件配置
├── eslint.config.mjs    # ESLint配置
├── index.html           # HTML模板
├── package.json         # 项目依赖配置
├── tsconfig.json        # TypeScript配置
├── vite.config.ts       # Vite配置
```

## 📦 主要依赖包介绍

### 核心依赖

- **vue**: - 渐进式JavaScript框架，本项目的核心框架
- **vue-router**: - Vue官方路由管理器
- **pinia**: - Vue官方推荐的状态管理库，替代Vuex
- **vite**: - 现代前端构建工具，提供极速的开发体验

### UI相关

- **shadcn/vue**: - [UI框架](https://npm.d1m.cn/-/web/detail/@d1m-mp/wsc-tracker)
  - **添加组件**:
    ```bash
      npx shadcn-vue@latest add button
    ```
- **tailwindcss**: - CSS元子化框架
- **@tailwindcss/vite**: - Tailwind CSS的Vite插件
- **tailwind-merge**: - 智能合并Tailwind CSS类
- **tw-animate-css**: - Tailwind CSS动画库

### 网络请求

- **axios**: - 基于Promise的HTTP客户端

### 开发工具

- **typescript**: - JavaScript的超集，添加静态类型
- **eslint**: - 代码质量和风格检查工具
- **@antfu/eslint-config**: - ESLint配置
- **vue-tsc**: - Vue文件的TypeScript类型检查

## 🚀 开发命令

```bash
# 安装依赖
npm install

# 启动开发服务器
npm dev

# 构建生产版本
npm build

# 代码检查
npm lint

# 自动修复代码问题
npm lint:fix
```

## 📚 相关文档

- [Vue 3文档](https://vuejs.org/)
- [Vue Router文档](https://router.vuejs.org/)
- [Pinia文档](https://pinia.vuejs.org/)
- [Vite文档](https://vitejs.dev/)
- [TypeScript文档](https://www.typescriptlang.org/)
- [Tailwind CSS文档](https://tailwindcss.com/)
- [shadcn/vue](https://npm.d1m.cn/-/web/detail/@d1m-mp/wsc-tracker/)
