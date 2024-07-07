
# MC 服务器配置文件

## 部署方式
在 `deploy` 目录中有一个 `docker-compose.yaml` 配置文件，可以在服务器上使用以下命令进行部署：
```sh
docker compose up -d
```
MC 服务器的数据将会保存在 `./mc-data` 目录下。

## 资源下载服务

### 1. main 目录
`main` 目录下包含了一个 HTTP 文件下载服务的源代码，用于下载 mod 和 resourcepacks 等资源文件。

### 2. files 目录
`files` 目录下包含所有的资源文件。

### 3. 编译好的二进制文件
编译好的二进制文件保存在 `build/bin` 目录下。
- `files-download-server` 是一个在 Linux 系统上运行的二进制文件，适用于 `amd64` 架构。它会将 `./files` 目录下的文件部署为 HTTP 服务，供用户下载。

## consts 目录
`consts` 目录下定义了全局变量。