# 全栈计算器项目

这是一个全栈计算器项目，展示了使用 Go + ConnectRPC 实现后端服务，以及使用 Next.js 与 ConnectRPC Web 实现前端调用的完整流程。项目支持基本的加减乘除运算（支持两个数计算），同时附带前后端单元测试.

## 目录结构

```plaintext
calculator/
├── client/                # 前端 Next.js 项目
│   ├── app/               # 使用 App Router 的前端页面（例如 app/page.tsx）
│   ├── src/
│   │   ├── gen/           # 由 protoc 生成的 Connect RPC 前端 stub 文件
│   │   └── types/         # 手动定义的类型文件（CalculatorRequest, CalculatorResponse）
│   ├── __tests__/         # 前端单元测试
│   ├── package.json
│   └── jest.config.js     # Jest 配置文件
├── proto/                 # 存放 .proto 文件
└── server/                # 后端 Go 项目
    ├── main.go            # 后端服务入口（使用 ConnectRPC 和 rs/cors）
    └── calculatorv1connect/  # 生成的 Connect 服务处理器代码
```

---

## 功能介绍

### 后端功能
- 使用 Go 与 ConnectRPC 实现基本加减乘除运算
- 支持两个数的四则运算并返回结果
- 使用 rs/cors 中间件解决跨域问题
- 监听端口：8080

### 前端功能
- 使用 Next.js 构建用户界面（App Router）
- 包含输入框、运算符选择器和结果展示区
- 通过 ConnectRPC Web 客户端调用后端服务
- 支持 JSON 格式数据交互
- 开发服务器运行端口：3000

### 单元测试
- **后端**：使用 Go 的 testing 包编写，运行 `go test -v`
- **前端**：使用 Jest 和 React Testing Library 编写

---

## 安装与运行

### 前端部分

1. **安装依赖**
```bash
cd client/
npm install
```
 如果出现 peer 依赖冲突，请尝试添加 --legacy-peer-deps 参数：
```
npm install --legacy-peer-deps
```
2. **运行前端开发服务器**
 在 client/ 目录下运行：
```
npm run dev
```
 前端服务会默认启动在 http://localhost:3000。

 ### 后端部分

 1. **安装Go依赖**
    在 server/ 目录下，确保 go.mod 文件中包含所有依赖（如 ConnectRPC、rs/cors 等）。如果未安装 rs/cors，请运行：
```
    go get github.com/rs/cors
```
  2. **运行后端服务**
     在 server/ 目录下运行：
```
     go run .
```
服务会监听在端口 8080，并通过 CORS 中间件允许来自 http://localhost:3000 的请求。

---
## 单元测试

### 前端单元测试
前端使用 Jest 和 React Testing Library 编写了测试用例。测试文件位于 client/__tests__/page.test.tsx。运行测试方法如下：

1. **运行前端测试**
   在 client/ 目录下运行：
```
   npm run test
```
   或者
```
   npx jest
```
   测试应当覆盖 UI 渲染和计算功能（模拟输入、点击计算按钮、检查计算结果）。

### 后端单元测试
后端使用 Go 的 testing 包编写了单元测试（例如在 server/ 目录下的 _test.go 文件）。运行测试方法如下：

1. **运行后端测试**
   在 server/ 目录下运行：
```
   go test -v
```
   确保所有测试用例都通过。

---

## 数据交互说明
- 前端与后端通过 ConnectRPC 进行通信。前端使用 Connect Web 客户端调用后端服务接口，后端返回 JSON 格式的响应（例如 {"result":10}）。
- 为了解决跨域问题，后端使用了 rs/cors 中间件，允许来自 http://localhost:3000 的请求。

---

## 注意事项
- 请确保 Go 后端服务运行在 http://localhost:8080，前端服务运行在 http://localhost:3000。
- 如果出现跨域问题，请检查后端 CORS 配置。

---

## 总结
这个项目展示了如何利用 Go 与 ConnectRPC 实现后端服务，以及如何使用 Next.js 和 ConnectRPC Web 实现前端调用，
并配合单元测试来确保代码质量。请按照上述说明运行项目，并通过单元测试验证各项功能
