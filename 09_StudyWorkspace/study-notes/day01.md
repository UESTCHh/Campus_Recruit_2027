# Day 01 — 2026-08-02

## 今日目标

搭建 Windows + WSL2 + Ubuntu + VS Code + Git + Go 开发环境，
完成从本地代码到 GitHub 的完整工作流。

## 今日完成

- [x] 启用 WSL2
- [x] 将 Ubuntu 24.04 安装到 D 盘
- [x] 创建 Linux 用户 hui
- [x] 确认 Ubuntu 使用 WSL2
- [x] 创建 `/home/hui/workspace`
- [x] 安装 Linux 基础开发工具
- [x] 配置 Git 用户名和邮箱
- [x] 创建 GitHub SSH 密钥
- [x] 成功连接 GitHub
- [x] 创建并推送 `wsl-git-test`
- [x] 将 VS Code 主程序和数据放到 D 盘
- [x] 安装 VS Code WSL 扩展
- [x] 正确进入 WSL Remote 模式
- [x] 安装 Go
- [x] 安装 gopls 和 goimports
- [x] 安装 VS Code Go 扩展
- [x] 创建并运行 `go-demo`
- [x] 将 `go-demo` 推送到 GitHub

## 当前环境

```text
Windows:
WSL: WSL2
Linux: Ubuntu 24.04 LTS
Linux 用户: hui
Ubuntu 存储位置: D:\WSL\Ubuntu-24.04
Linux 工作目录: /home/hui/workspace
GitHub 用户名: UESTCHh
Go 项目: ~/workspace/go-demo

重要目录：
/home/hui
├── workspace
│   ├── git-test
│   ├── go-demo
│   └── study-notes
└── go
    └── bin
        ├── gopls
        └── goimports

昨天遇到的问题
1. Ubuntu 默认安装位置问题
现象：
默认使用 wsl --install 会把 Ubuntu 发行版安装到 C 盘。
解决：
wsl --install -d Ubuntu-24.04 --location D:\WSL\Ubuntu-24.04
结论：
WSL Windows 组件仍在系统盘，但 Ubuntu 文件系统和 Linux 开发环境主要位于 D 盘。
2. 把路径当作命令执行
错误操作：
/home/hui/workspace
正确操作：
mkdir -p ~/workspace
cd ~/workspace
结论：
路径本身不是进入目录的命令，进入目录需要使用 cd。
3. GitHub 远程仓库地址拼写错误
现象：
ERROR: Repository not found.
原因：
GitHub 用户名拼写错误。
正确地址：
git@github.com:UESTCHh/仓库名.git
修复：
git remote set-url origin git@github.com:UESTCHh/仓库名.git
4. VS Code 没有真正进入 WSL
错误状态：
PowerShell
\\wsl.localhost\Ubuntu-24.04\...
正确状态：
WSL: Ubuntu-24.04
bash
/home/hui/workspace/...
验证命令：
pwd
whoami
uname -a
5. Go 模块下载超时
现象：
proxy.golang.org: i/o timeout
解决：
go env -w GOPROXY=https://goproxy.cn,direct
6. Ctrl+Z 导致命令暂停
现象：
Stopped
原因：
Ctrl+Z 会暂停前台任务，而不是终止任务。
正确操作：
Ctrl+C：终止当前程序
Ctrl+Z：暂停当前程序
今日 Git 工作流
修改文件
→ git add
→ git commit
→ git push
→ GitHub
常用命令：
git status
git add .
git commit -m "提交说明"
git push
git log --oneline
git remote -v
今日 Go 工作流
初始化模块：
go mod init github.com/UESTCHh/go-demo
运行代码：
go run ./cmd/main.go
查看版本：
go version
gopls version
which goimports
闭卷问题
先不要查答案，用自己的语言回答。
1. git add 和 git commit 分别做什么？
我的回答：
2. git push -u origin main 中的 -u 有什么作用？
我的回答：
3. go mod init 创建了什么？
我的回答：
4. /home/hui/workspace 实际存储在哪个盘？
我的回答：
5. 如何判断 VS Code 当前在 Windows 还是 WSL 中？
我的回答：
6. Ctrl+C 和 Ctrl+Z 有什么区别？
我的回答：
7. 为什么代码放在 ~/workspace，而不是 /mnt/d/Projects？
我的回答：
Day 1 最小验收
cd ~/workspace/go-demo

pwd
whoami
git status
git remote -v
go version
gopls version
go run ./cmd/main.go
预期结果：
当前用户是 hui
Git 工作区干净
远程地址属于 UESTCHh
Go 可以正常运行
程序输出 Hello Go Backend!
后续复习
第 1 次：2026-08-03
第 3 次：2026-08-05
第 7 次：2026-08-09
第 21 次：2026-08-23