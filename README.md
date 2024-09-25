# ebook

开发一个在线电子书web服务。

- 安装

执行 `npm install && npm run dev`  --安装前端依赖和监听tailwindcss

构建 `docker compose build`

运行 `docker compose up`

或者 查看 `Makefile` 文件执行对应命令


## 技术栈


## 好记性不如烂笔头

发现 `templ` 和 `HTMX` 很有意思，第一次尝试，记录一下，方便复盘。

### 使用docker构建开发依赖

在容器里安装 `air` 和 `templ`, 本机不需要安装。

.air.toml 配置,需要加上`templ generate`,修改模板都需要重新生成,
include_ext 需要加上 `templ`，默认例子是没有的

```toml
cmd = "templ generate &&  go build -o ./tmp/main ."
include_ext = ["go", "tpl", "tmpl", "templ", "html"]
```

### 本地环境配置

参考文档配置：
[点我直达](https://templ.guide/commands-and-tools/ide-support/#visual-studio-code)

- 插件

`Tailwind CSS IntelliSense` 和 `templ-vscode`

- 配置，settings.json 添加以下配置

```json
"tailwindCSS.includeLanguages": {
    "templ": "html"
},
"[templ]": {
    "editor.defaultFormatter": "a-h.templ"
},
"emmet.includeLanguages": {
    "templ": "html"
}
```

