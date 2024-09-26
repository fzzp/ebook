package main

// NOTE: 废弃

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"os/signal"
// 	"path/filepath"
// 	"syscall"
// )

// // watch 执行命tailwindcss命令css放到指定目录
// func watch(css string) {
// 	// npx tailwindcss -i ./tailwind/home.css -o ./static/css/home.css --minify --watch
// 	dst := fmt.Sprintf("./static/css/%s", filepath.Base(css))
// // 这里如果是用 --watch 会生成失败
// 	cmd := exec.Command("npx", "tailwindcss", "-i", css, "-o", dst, "--minify")
// 	// fmt.Println(cmd.String())
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Printf("watch %s err: %v\n", css, err)
// 		return
// 	}
// }

// func main() {
// 	quitChan := make(chan os.Signal, 1)
// 	cssFiles, err := filepath.Glob("./tailwind/*.css")
// 	if err != nil {
// 		log.Fatal("匹配文件失败: ", err)
// 	}

// 	for _, css := range cssFiles {
// 		go watch(css)
// 	}

// 	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

// 	<-quitChan
// }
