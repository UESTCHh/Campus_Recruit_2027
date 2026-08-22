# 2026-08-03 学习笔记：FastAPI AI 后端工程化与 Docker 生产化

> 建议保存路径：`~/workspace/study-notes/2026-08-03.md`  
> 项目目录：`~/workspace/ai-service-demo`  
> 阶段版本：`v0.1.0-production-ready`

---

## 一、今日学习目标

今天的核心目标，是把一个可以运行的 FastAPI Demo，升级为具备基本生产工程结构的 AI 后端服务。

最终完成了：

- FastAPI API 分层与版本化
- Application Factory 应用工厂
- 路由模块化
- 自动化测试迁移
- Ruff 代码检查
- MyPy 类型检查
- GitHub Actions CI
- Docker 镜像优化
- 非 root 容器运行
- Docker Compose 生产化配置
- 健康检查、资源限制和日志轮转
- Git 阶段版本标签

---

## 二、最终项目结构

```text
ai-service-demo/
├── .github/
│   └── workflows/
│       └── ci.yml
├── app/
│   ├── api/
│   │   ├── __init__.py
│   │   └── v1/
│   │       ├── __init__.py
│   │       ├── router.py
│   │       ├── health.py
│   │       ├── chat.py
│   │       ├── utility.py
│   │       └── debug.py
│   ├── llm/
│   │   ├── __init__.py
│   │   ├── base.py
│   │   └── mock.py
│   ├── middleware/
│   │   ├── __init__.py
│   │   └── request_id.py
│   ├── services/
│   │   ├── __init__.py
│   │   └── chat_service.py
│   ├── __init__.py
│   ├── config.py
│   ├── error_handlers.py
│   ├── exceptions.py
│   ├── factory.py
│   ├── logger.py
│   ├── main.py
│   └── schemas.py
├── tests/
│   ├── __init__.py
│   └── test_main.py
├── .dockerignore
├── .env
├── .env.example
├── .gitignore
├── .python-version
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── pyproject.toml
├── README.md
└── uv.lock
```

---

# 三、FastAPI API 版本化

## 3.1 为什么需要 API 版本

原来的接口没有统一版本前缀：

```text
/healthz
/chat/mock
/echo
/error-test
```

完成版本化之后：

```text
/api/v1/healthz
/api/v1/chat/mock
/api/v1/utility/echo
/api/v1/debug/error-test
```

这样未来可以同时存在：

```text
/api/v1/...
/api/v2/...
```

旧客户端仍然使用 `v1`，新客户端逐步迁移到 `v2`。

---

## 3.2 路由前缀是逐层拼接的

FastAPI 路由最终地址由多层前缀组合：

```text
应用前缀 + Router 前缀 + Endpoint 路径
```

例如：

```python
app.include_router(api_router, prefix="/api/v1")
```

```python
router = APIRouter(prefix="/chat")
```

```python
@router.post("/mock")
```

最终路径是：

```text
/api/v1/chat/mock
```

今天曾经出现过：

```python
router = APIRouter(prefix="/chat")

@router.post("/chat/mock")
```

这会生成：

```text
/api/v1/chat/chat/mock
```

因此测试请求 `/api/v1/chat/mock` 返回了 `404`。

正确原则是：

> 每一层只负责自己的路径片段，不要重复写上层前缀。

---

## 3.3 v1 总路由

`app/api/v1/router.py` 负责集中注册各模块：

```python
from fastapi import APIRouter

from app.api.v1.chat import router as chat_router
from app.api.v1.debug import router as debug_router
from app.api.v1.health import router as health_router
from app.api.v1.utility import router as utility_router

router = APIRouter()

router.include_router(health_router)
router.include_router(chat_router)
router.include_router(utility_router)
router.include_router(debug_router)
```

这样 `main.py` 不需要分别了解每个业务路由。

---

# 四、路由模块化

## 4.1 健康检查

接口：

```http
GET /api/v1/healthz
```

返回：

```json
{
  "status": "ok"
}
```

用途：

- 判断应用是否成功启动
- 为 Docker Healthcheck 提供检测地址
- 为负载均衡器和监控系统提供存活检测

---

## 4.2 Mock Chat

接口：

```http
POST /api/v1/chat/mock
```

请求示例：

```json
{
  "message": "测试 LLM",
  "session_id": "test-session"
}
```

响应示例：

```json
{
  "answer": "模拟回答: 测试 LLM",
  "session_id": "test-session",
  "model": "mock-model-v1"
}
```

调用链：

```text
HTTP Request
    ↓
chat.py
    ↓
chat_service.py
    ↓
MockLLM.generate()
    ↓
ChatResponse
```

这体现了 API 层、Service 层和 LLM Provider 层之间的职责分离。

---

## 4.3 Utility Echo

接口：

```http
POST /api/v1/utility/echo
```

用于验证：

- JSON 请求解析
- Pydantic 输入校验
- Pydantic 响应模型
- 字符串长度计算
- `422` Validation Error

---

## 4.4 Debug Error Test

接口：

```http
GET /api/v1/debug/error-test
```

该接口主动抛出：

```python
AIServiceException(
    message="test error",
    code=50001,
    status_code=500,
)
```

用于验证：

- 自定义异常
- 统一错误响应
- Request ID
- 服务端错误日志

---

# 五、Application Factory 应用工厂

## 5.1 改造前

之前 `main.py` 同时负责：

- 读取配置
- 初始化日志
- 创建 FastAPI
- 注册中间件
- 注册异常处理器
- 注册路由

这些操作都在模块导入时直接执行。

---

## 5.2 改造后

创建：

```text
app/factory.py
```

核心结构：

```python
def create_app() -> FastAPI:
    settings = get_settings()

    setup_logging()

    app = FastAPI(
        title=settings.app_name,
        version=settings.app_version,
        description="Production-style AI backend API",
    )

    app.add_middleware(RequestIDMiddleware)
    register_exception_handlers(app)
    app.include_router(api_router, prefix="/api/v1")

    return app
```

`app/main.py` 最终只保留：

```python
from app.factory import create_app

app = create_app()
```

---

## 5.3 Application Factory 的意义

应用工厂使应用的创建过程集中、明确且可复用。

以后可以更方便地支持：

```text
开发环境
测试环境
预发布环境
生产环境
```

测试时也可以单独创建一个新的应用实例，减少全局状态相互影响。

---

# 六、测试迁移与错误排查

## 6.1 API 版本化后的 404

路由改成 `/api/v1/...` 后，测试仍然请求旧地址：

```python
client.get("/healthz")
client.post("/echo")
client.post("/chat/mock")
client.get("/error-test")
```

所以出现：

```text
404 Not Found
```

测试必须同步改为：

```python
client.get("/api/v1/healthz")
client.post("/api/v1/utility/echo")
client.post("/api/v1/chat/mock")
client.get("/api/v1/debug/error-test")
```

经验：

> 接口重构后，应立即同步测试、README、Swagger 示例和 Docker Healthcheck。

---

## 6.2 循环导入

`debug.py` 曾经错误地包含：

```python
from app.api.v1.debug import router as debug_router
```

也就是模块导入自己：

```text
router.py
  → debug.py
      → debug.py
          → debug.py
```

最终报错：

```text
ImportError: cannot import name 'router'
from partially initialized module
```

修复方法：

`debug.py` 只定义自己的路由，不负责注册其他 Router。

---

## 6.3 路由路径重复

错误写法：

```python
router = APIRouter(prefix="/chat")

@router.post("/chat/mock")
```

实际地址：

```text
/api/v1/chat/chat/mock
```

正确写法：

```python
router = APIRouter(prefix="/chat")

@router.post("/mock")
```

---

## 6.4 最终测试结果

```bash
make test
```

结果：

```text
....... [100%]
7 passed in 0.48s
```

说明当前全部测试通过。

---

# 七、Ruff 代码检查

## 7.1 Ruff 的作用

Ruff 用于：

- 检查代码错误
- 检查未使用的 import
- 检查 import 排序
- 统一代码格式
- 替代部分 Black、Flake8 和 isort 的功能

执行：

```bash
uv run ruff check .
```

最初发现：

```text
I001 Import block is un-sorted or un-formatted
```

自动修复：

```bash
uv run ruff check . --fix
```

最终结果：

```text
All checks passed!
```

---

## 7.2 Ruff 配置

`pyproject.toml` 中加入：

```toml
[tool.ruff]
line-length = 88
target-version = "py312"

[tool.ruff.lint]
select = [
    "E",
    "F",
    "I",
]

[tool.ruff.format]
quote-style = "double"
```

含义：

```text
E：代码风格错误
F：Pyflakes 代码错误
I：import 排序
```

---

# 八、MyPy 类型检查

执行：

```bash
uv run mypy app
```

最终结果：

```text
Success: no issues found in 22 source files
```

MyPy 检查的是静态类型，例如：

```python
def generate_reply(message: str) -> str:
    ...
```

它可以提前发现：

- 返回值类型错误
- 参数类型错误
- 漏写类型标注
- 异步函数和普通函数使用错误
- Optional 值未处理

今天之前出现过协程未等待的问题：

```text
RuntimeWarning: coroutine 'chat' was never awaited
```

并导致 Pydantic 收到协程对象，而不是字符串。

这类问题说明：

> `async def` 返回的是 coroutine，调用时通常需要 `await`；普通同步函数则直接调用。

---

# 九、Makefile 开发命令

Makefile 把较长的命令封装为统一入口。

主要命令：

```bash
make help
make build
make up
make down
make restart
make test
make lint
make format
make type
make ci
make logs
make shell
```

典型使用方式：

```bash
make test
```

实际执行：

```bash
uv run pytest -q
```

提交前可以执行：

```bash
make ci
```

统一完成：

```text
Ruff
MyPy
Pytest
Docker Build
```

---

# 十、GitHub Actions CI

创建文件：

```text
.github/workflows/ci.yml
```

CI 在以下情况运行：

```text
push 到 main
pull request 合并到 main
```

自动执行：

```text
Checkout 代码
安装 uv
安装 Python 3.12
同步依赖
Ruff 检查
MyPy 类型检查
Pytest 测试
Docker Build
```

CI 的作用不是代替本地测试，而是保证：

> 任何人提交的代码，都必须在统一、干净的 Linux 环境中重新验证。

---

# 十一、Docker Desktop 与 WSL

## 11.1 Docker Desktop 架构

当前环境：

```text
Windows
  └── Docker Desktop
        └── WSL 2 Backend
              └── Linux Containers
```

在 Ubuntu WSL 终端中执行：

```bash
docker version
docker compose version
docker ps
docker build .
```

实际连接的是 Docker Desktop 提供的 Docker Engine。

---

## 11.2 Docker Socket 权限

曾经出现：

```text
permission denied while trying to connect
to the Docker API at unix:///var/run/docker.sock
```

解决方式是把当前用户加入 `docker` 组：

```bash
sudo usermod -aG docker $USER
```

然后必须退出并重新进入 WSL，使新的用户组生效。

检查：

```bash
groups
```

应包含：

```text
docker
```

---

## 11.3 Docker Hub 网络问题

拉取镜像时曾出现：

```text
dial tcp ... timeout
```

以及代理地址无法解析：

```text
lookup host.docker.internal: no such host
```

最终通过：

- 启动 VPN
- 正确配置 Docker Desktop Proxy
- 重启 Docker Desktop
- 重新执行 `docker pull`

成功拉取：

```bash
docker pull python:3.12-slim
```

经验：

> 浏览器能上网，不代表 Docker Engine 一定能连接 Docker Hub。浏览器、Windows、WSL 和 Docker Engine 可能使用不同代理路径。

---

# 十二、Dockerfile 生产化优化

## 12.1 优化前

原 Dockerfile：

```dockerfile
FROM python:3.12-slim

WORKDIR /app

COPY pyproject.toml uv.lock ./

RUN pip install uv \
    && uv sync --frozen

COPY app ./app

RUN useradd --create-home appuser \
    && chown -R appuser:appuser /app

USER appuser

EXPOSE 8000

CMD ["uv", "run", "fastapi", "run", "app/main.py", "--host", "0.0.0.0"]
```

主要问题：

- 使用 `pip install uv`
- 把开发依赖装进生产镜像
- 使用 FastAPI CLI 启动
- 镜像体积较大

---

## 12.2 优化后的核心思路

使用 uv 官方镜像复制二进制：

```dockerfile
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/
```

只安装生产依赖：

```dockerfile
RUN uv sync \
    --frozen \
    --no-dev \
    --no-install-project
```

使用 Uvicorn 生产启动：

```dockerfile
CMD ["uv", "run", "--no-dev", "uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
```

保留非 root 用户：

```dockerfile
USER appuser
```

---

## 12.3 Dockerfile JSON 格式注意事项

错误写法：

```dockerfile
CMD [
    "uv",
    "run"
]
```

Dockerfile 会把下一行 `"uv",` 当作新的指令，出现：

```text
unknown instruction: "uv",
```

正确写法是放在一行：

```dockerfile
CMD ["uv", "run", "uvicorn", "app.main:app"]
```

---

## 12.4 镜像体积变化

旧镜像：

```text
ai-service-demo-api:latest
DISK USAGE: 679MB
CONTENT SIZE: 174MB
```

优化镜像：

```text
ai-service-demo:latest
DISK USAGE: 378MB
CONTENT SIZE: 89.7MB
```

优化效果：

```text
磁盘显示体积减少约 301MB
内容体积减少约 84MB
```

---

## 12.5 非 root 用户验证

容器运行后：

```bash
docker exec -it 581668d75a0e whoami
```

结果：

```text
appuser
```

进一步检查：

```bash
docker inspect 581668d75a0e --format '{{.Config.User}}'
```

结果：

```text
appuser
```

说明生产容器不是以 root 身份运行。

---

# 十三、Docker 端口映射

运行：

```bash
docker run --rm -p 8000:8000 ai-service-demo:optimized
```

含义：

```text
宿主机 8000
    ↓
容器 8000
```

浏览器访问：

```text
http://127.0.0.1:8000/docs
```

如果报错：

```text
Bind for 0.0.0.0:8000 failed:
port is already allocated
```

说明已有容器或程序占用宿主机 8000。

检查：

```bash
docker ps
```

停掉旧容器：

```bash
docker stop <容器 ID>
```

注意：

```text
<容器 ID>
```

只是占位符，实际输入时不能保留尖括号。

正确：

```bash
docker stop 581668d75a0e
```

错误：

```bash
docker stop <581668d75a0e>
```

---

# 十四、Docker Compose 生产化配置

最终 Compose 配置包含：

```yaml
services:
  api:
    image: ai-service-demo:latest

    build:
      context: .
      dockerfile: Dockerfile

    container_name: ai-service-demo

    ports:
      - "8000:8000"

    env_file:
      - .env

    restart: unless-stopped

    stop_grace_period: 30s

    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"

    mem_limit: 512m
    cpus: 1.0

    healthcheck:
      test:
        [
          "CMD",
          "python",
          "-c",
          "import urllib.request; urllib.request.urlopen('http://localhost:8000/api/v1/healthz')"
        ]
      interval: 30s
      timeout: 5s
      retries: 3
      start_period: 10s
```

---

## 14.1 健康检查路径同步

API 版本化前：

```text
/healthz
```

版本化后：

```text
/api/v1/healthz
```

Compose 最初仍然检测旧路径，容器因此显示：

```text
unhealthy
```

修改 Healthcheck 后：

```text
Up 15 seconds (healthy)
```

经验：

> API 地址发生变化时，要同时检查测试、README、监控和容器健康检查。

---

## 14.2 资源限制

```yaml
mem_limit: 512m
cpus: 1.0
```

验证：

```bash
docker stats
```

可以看到：

```text
MEM USAGE / LIMIT
62.53MiB / 512MiB
```

---

## 14.3 日志轮转

```yaml
logging:
  driver: json-file
  options:
    max-size: "10m"
    max-file: "3"
```

作用：

- 单个日志文件最大约 10MB
- 最多保留 3 个文件
- 防止容器日志无限增长

---

## 14.4 优雅关闭

```yaml
stop_grace_period: 30s
```

Docker 停止容器时会给应用一段时间：

- 完成正在处理的请求
- 执行 shutdown 生命周期
- 关闭连接
- 释放资源

之后才会强制终止。

---

# 十五、今日最终验证

## Ruff

```bash
make lint
```

结果：

```text
All checks passed!
```

## MyPy

```bash
uv run mypy app
```

结果：

```text
Success: no issues found in 22 source files
```

## Pytest

```bash
make test
```

结果：

```text
7 passed in 0.48s
```

## Git

```bash
git status
```

结果：

```text
nothing to commit, working tree clean
```

## Docker

```bash
docker ps
```

结果：

```text
ai-service-demo:latest
Up ... (healthy)
```

---

# 十六、今日主要 Git 提交

```text
0f99629 chore: improve production compose configuration
3f6e279 chore: optimize production docker image
f5d0530 chore: improve developer commands
1eeeb99 ci: add github actions pipeline
c08ea40 style: format imports with ruff
d507a27 refactor: introduce application factory
d8be513 refactor: modularize api v1 routers
b0d96c6 docs: add project documentation
eddb630 chore: add makefile developer commands
a92f3e8 chore: add production docker resource limits and logging
```

创建阶段标签：

```bash
git tag v0.1.0-production-ready
```

验证：

```bash
git tag
```

结果：

```text
v0.1.0-production-ready
```

---

# 十七、今日最重要的工程经验

1. **重构路由后，必须同步修改测试和 Healthcheck。**
2. **Router prefix 与 endpoint path 会逐层拼接。**
3. **模块不能导入自己，否则会造成循环导入。**
4. **测试失败不是坏事，它能及时暴露重构遗漏。**
5. **Docker 镜像、容器和端口是三个不同概念。**
6. **浏览器能访问网络，不代表 Docker Engine 能访问网络。**
7. **容器应尽可能使用非 root 用户运行。**
8. **生产镜像不应安装 Ruff、MyPy、Pytest 等开发依赖。**
9. **本地检查和 GitHub CI 应保持一致。**
10. **每个稳定阶段都应该提交 Git，并创建可回退的版本节点。**

---

# 十八、下一阶段计划

下一阶段进入真实 AI 能力开发：

```text
Phase 2：Real LLM Capability
```

建议顺序：

1. 设计 LLM Provider 接口
2. 增加 Provider Factory
3. 接入真实模型 API
4. 配置 API Key
5. 实现超时和重试
6. 实现流式响应
7. 增加 Token Usage
8. 增加对话历史
9. 接入 Redis
10. 接入 PostgreSQL
11. 实现 RAG 知识库

---

# 为什么今天 C 盘少了大约 10GB？

首先，这里减少的是 **C 盘磁盘空间**，不是运行内存 RAM。

## 最可能的原因

今天同时安装和使用了：

- Docker Desktop
- WSL 2 Docker 后端
- 多个 Docker 镜像
- Docker BuildKit 构建缓存
- Ubuntu WSL 虚拟磁盘
- Python `.venv`
- uv 下载缓存
- VS Code Server
- Ruff、MyPy、Pytest 缓存

所以减少约 10GB 是完全可能的。

---

## 1. Docker Desktop 默认可能把数据放在 C 盘

Docker Desktop 使用 WSL 2 后端时，默认数据位置通常位于：

```text
C:\Users\<用户名>\AppData\Local\Docker\wsl
```

其中保存 Docker 镜像、容器、层和构建数据。Docker Desktop 可以在 **Settings → Resources → Advanced** 中查看或调整磁盘映像位置。

即使 Docker Desktop 程序本体安装在其他位置，也不代表 Docker 数据一定不在 C 盘。

你应当在 Docker Desktop 中检查：

```text
Settings
→ Resources
→ Advanced
→ Disk image location
```

如果这里显示的是 C 盘，那么 Docker 是主要原因之一。

---

## 2. 今天产生了多个 Docker 镜像和构建缓存

你现在至少有：

```text
ai-service-demo-api:latest
ai-service-demo:latest
ai-service-demo:optimized
python:3.12-slim
hello-world:latest
```

虽然 `docker images` 中的数字不能简单相加——镜像之间会共享层——但它也没有完整展示所有 BuildKit 构建缓存。查看 Docker 实际占用应使用：

```bash
docker system df -v
docker buildx du
```

Docker 官方说明，未使用的镜像、停止的容器、网络和构建缓存通常不会自动全部删除，因此开发过程中反复构建会持续占用空间。

---

## 3. Ubuntu WSL 自己也有一个虚拟磁盘

你的代码位于：

```text
/home/hui/workspace/ai-service-demo
```

它不是普通的 Windows 文件夹，而是存储在 Ubuntu WSL 的 Linux 虚拟磁盘中。

WSL 2 会为每个发行版建立一个 `ext4.vhdx` 虚拟硬盘文件，并随着 Linux 文件增加而扩展。

这个虚拟盘中包含：

```text
~/workspace/ai-service-demo
~/workspace/ai-service-demo/.venv
~/.cache/uv
~/.vscode-server
Python 包
Ubuntu 系统文件
```

如果 Ubuntu 发行版仍按默认位置安装，`ext4.vhdx` 很可能也位于 C 盘。

---

## 4. `.venv` 和 uv 缓存也会占空间

你项目中的：

```text
.venv
```

安装了：

- FastAPI
- Uvicorn
- Pydantic
- HTTPX
- Pytest
- Ruff
- MyPy
- 其他间接依赖

同时 uv 会保留下载缓存，避免下次重新下载。

可在 **VS Code 的 WSL 终端**执行：

```bash
du -sh .venv
du -sh ~/.cache/uv 2>/dev/null
du -sh ~/.vscode-server 2>/dev/null
du -sh .pytest_cache .ruff_cache .mypy_cache 2>/dev/null
```

这些目录单独可能不够 10GB，但会和 Docker、WSL 虚拟磁盘一起累积。

---

# 如何确认空间到底被谁占了

## 第一步：检查 Docker 占用

**操作位置：VS Code 的 WSL 终端**

```bash
docker system df -v
```

检查构建缓存：

```bash
docker buildx du
```

Docker 官方将 `docker system df` 定义为查看 Docker daemon 磁盘占用的命令。

---

## 第二步：检查 Linux 开发文件

**操作位置：VS Code 的 WSL 终端**

```bash
du -sh \
  ~/.cache \
  ~/.vscode-server \
  ~/workspace/ai-service-demo/.venv \
  ~/workspace/ai-service-demo \
  2>/dev/null
```

---

## 第三步：查看 Docker 磁盘映像位置

**操作位置：Docker Desktop**

```text
Settings
→ Resources
→ Advanced
→ Disk image location
```

Docker 官方建议从该页面查看和移动磁盘映像，不要直接在文件资源管理器中手动移动 Docker 的 VHDX 文件。

---

## 第四步：在 Windows 查找大型 VHDX

**操作位置：Windows PowerShell**

```powershell
Get-ChildItem "$env:LOCALAPPDATA\Docker\wsl" `
    -Recurse `
    -Filter *.vhdx `
    -ErrorAction SilentlyContinue |
Select-Object FullName,
    @{Name="SizeGB"; Expression={[math]::Round($_.Length / 1GB, 2)}}
```

查找 Ubuntu 的虚拟磁盘：

```powershell
Get-ChildItem "$env:LOCALAPPDATA\Packages" `
    -Recurse `
    -Filter ext4.vhdx `
    -ErrorAction SilentlyContinue |
Select-Object FullName,
    @{Name="SizeGB"; Expression={[math]::Round($_.Length / 1GB, 2)}}
```

这两个结果基本就能确认 10GB 去了哪里。

---

# 可以安全清理哪些内容

当前仍在使用：

```text
ai-service-demo:latest
```

建议保留。

较旧镜像可以考虑删除：

```bash
docker image rm ai-service-demo-api:latest
docker image rm ai-service-demo:optimized
```

如果提示镜像被容器使用，先检查：

```bash
docker ps -a
```

不要强制删除正在使用的镜像。

---

## 清理构建缓存

先查看：

```bash
docker buildx du
```

再执行：

```bash
docker buildx prune
```

该命令只清理当前 Builder 的可回收构建缓存，并会先要求确认。

---

## Docker 综合清理

```bash
docker system prune
```

默认会清理：

- 已停止容器
- 未使用网络
- dangling 镜像
- 未使用构建缓存

它默认不会删除 Volume。

目前不要直接使用：

```bash
docker system prune -a --volumes
```

因为它可能删除：

- 所有未被容器引用的镜像
- 未使用 Volume
- 将来数据库中的持久数据

---

## 清理小型开发缓存

**操作位置：VS Code 的 WSL 终端**

```bash
rm -rf .pytest_cache .ruff_cache .mypy_cache
```

可选清理 uv 缓存：

```bash
rm -rf ~/.cache/uv
```

清理后下次 `uv sync` 可能需要重新下载依赖。

---

# 删除以后为什么 C 盘可能没有立刻恢复？

Docker 和 WSL 的 Linux 文件系统存储在虚拟硬盘文件中。WSL 的 VHD 会随着数据增加而扩展，而 Windows 层看到的是这个 VHDX 文件。

新版 Docker Desktop 在 Windows WSL 2 托管虚拟磁盘上支持自动回收部分空间，但回收不一定与文件删除同时发生；Ubuntu 发行版自己的 `ext4.vhdx` 又是另一块虚拟磁盘。

因此常见现象是：

```text
Linux 内部已经删除文件
但 Windows C 盘可用空间没有立即等量增加
```

不要直接删除、重命名或手动移动任何 `docker_data.vhdx` 或 `ext4.vhdx`。手动压缩虚拟磁盘需要先关闭 WSL、确认准确路径并做好备份；虚拟磁盘压缩本身要求磁盘处于分离或只读状态。

---

## 当前最合理的判断

仅凭目前的 `docker images` 输出，还不能精确断定 10GB 全部来自哪里。

最可能的组合是：

```text
Docker Desktop 安装和数据
+ Docker 镜像
+ BuildKit 构建缓存
+ Ubuntu ext4.vhdx 增长
+ Python .venv
+ uv 与 VS Code 缓存
≈ 10GB
```

优先执行下面三条即可定位：

```bash
docker system df -v
docker buildx du
du -sh ~/.cache ~/.vscode-server ~/workspace/ai-service-demo/.venv 2>/dev/null
```

然后再检查 Docker Desktop 的 `Disk image location`。不要在未查看占用明细前进行大范围删除。