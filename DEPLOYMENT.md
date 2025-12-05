# GoFilm 项目部署文档

## 1. 项目概述

GoFilm是一个基于Go语言和Vue.js开发的在线影视网站，支持多播放源自动采集和在线观看。本部署文档将指导您如何使用PostgreSQL数据库部署GoFilm项目。

## 2. 技术栈

- **后端**: Go 1.18+, Gin框架, GORM ORM
- **前端**: Vue 3, Vite
- **数据库**: PostgreSQL 13+
- **缓存**: Redis 6+
- **容器化**: Docker

## 3. 部署前准备

### 3.1 环境要求

- Linux/macOS/Windows
- Go 1.18+
- Node.js 16+
- PostgreSQL 13+
- Redis 6+
- Docker (可选，用于容器化部署)

### 3.2 系统依赖安装

#### Ubuntu/Debian

```bash
sudo apt update
sudo apt install -y golang-go nodejs npm postgresql redis-server
sudo systemctl enable postgresql redis-server
sudo systemctl start postgresql redis-server
```

#### CentOS/RHEL

```bash
sudo yum install -y golang nodejs npm postgresql-server redis
sudo postgresql-setup initdb
sudo systemctl enable postgresql redis
sudo systemctl start postgresql redis
```

## 4. 数据库配置

### 4.1 创建数据库和用户

1. 登录PostgreSQL:

```bash
sudo -u postgres psql
```

2. 创建数据库和用户:

```sql
CREATE DATABASE gofilm;
CREATE USER gofilm WITH ENCRYPTED PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE gofilm TO gofilm;
\q
```

### 4.2 配置PostgreSQL

编辑PostgreSQL配置文件 (`postgresql.conf`):

```bash
# Ubuntu/Debian
sudo nano /etc/postgresql/13/main/postgresql.conf

# CentOS/RHEL
sudo nano /var/lib/pgsql/data/postgresql.conf
```

添加或修改以下配置:

```
# 允许远程连接
listen_addresses = '*'

# 优化PostgreSQL性能
shared_buffers = 256MB
work_mem = 4MB
maintenance_work_mem = 64MB
effective_cache_size = 768MB
wal_buffers = 16MB
default_statistics_target = 100
random_page_cost = 1.1
```

编辑pg_hba.conf文件，允许远程连接:

```bash
# Ubuntu/Debian
sudo nano /etc/postgresql/13/main/pg_hba.conf

# CentOS/RHEL
sudo nano /var/lib/pgsql/data/pg_hba.conf
```

添加以下行:

```
host    all             all             0.0.0.0/0               md5
host    all             all             ::/0                    md5
```

重启PostgreSQL服务:

```bash
# Ubuntu/Debian
sudo systemctl restart postgresql

# CentOS/RHEL
sudo systemctl restart postgresql
```

## 5. 项目部署

### 5.1 克隆项目代码

```bash
git clone https://github.com/hangyubin/GoFilm.git
cd GoFilm
```

### 5.2 配置文件修改

编辑后端配置文件 `server/config/config.yaml`:

```yaml
# 数据库配置
database:
  type: postgres
  host: localhost
  port: 5432
  user: gofilm
  password: your_password
  dbname: gofilm

# Redis配置
redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

# 服务器配置
server:
  port: 3600
  mode: release
```

### 5.3 后端部署

1. 安装依赖:

```bash
cd server
go mod tidy
```

2. 编译后端:

```bash
go build -o gofilm main.go
```

3. 启动后端服务:

```bash
./gofilm
```

或者使用systemd管理服务，创建`/etc/systemd/system/gofilm.service`:

```ini
[Unit]
Description=GoFilm Backend Service
After=network.target postgresql.service redis.service

[Service]
Type=simple
WorkingDirectory=/path/to/GoFilm/server
ExecStart=/path/to/GoFilm/server/gofilm
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable gofilm
sudo systemctl start gofilm
```

### 5.4 前端部署

1. 安装依赖:

```bash
cd ../client
npm install
```

2. 构建前端:

```bash
npm run build
```

3. 部署前端代码:

将`dist`目录下的文件部署到Nginx或其他Web服务器上。

#### Nginx配置示例

创建`/etc/nginx/conf.d/gofilm.conf`:

```nginx
server {
    listen 80;
    server_name your_domain.com;

    root /path/to/GoFilm/client/dist;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    # 反向代理后端API
    location /api {
        proxy_pass http://localhost:3600;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

重启Nginx服务:

```bash
sudo systemctl restart nginx
```

## 6. Docker容器化部署

### 6.1 使用Docker Compose

1. 创建`docker-compose.yml`文件:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:13
    restart: always
    environment:
      POSTGRES_DB: gofilm
      POSTGRES_USER: gofilm
      POSTGRES_PASSWORD: your_password
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  redis:
    image: redis:6
    restart: always
    ports:
      - "6379:6379"

  backend:
    build:
      context: ./server
      dockerfile: Dockerfile
    restart: always
    depends_on:
      - postgres
      - redis
    environment:
      - DATABASE_TYPE=postgres
      - DATABASE_HOST=postgres
      - DATABASE_PORT=5432
      - DATABASE_USER=gofilm
      - DATABASE_PASSWORD=your_password
      - DATABASE_NAME=gofilm
      - REDIS_HOST=redis
      - REDIS_PORT=6379
    ports:
      - "3600:3600"

  frontend:
    build:
      context: ./client
      dockerfile: Dockerfile
    restart: always
    depends_on:
      - backend
    ports:
      - "80:80"

volumes:
  postgres_data:
```

2. 在后端和前端目录创建Dockerfile:

#### 后端 Dockerfile (`server/Dockerfile`):

```dockerfile
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gofilm main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gofilm .
COPY config/config.yaml ./config/

EXPOSE 3600

CMD ["./gofilm"]
```

#### 前端 Dockerfile (`client/Dockerfile`):

```dockerfile
FROM node:16-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

3. 创建前端Nginx配置文件 (`client/nginx.conf`):

```nginx
server {
    listen 80;
    server_name localhost;

    root /usr/share/nginx/html;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    location /api {
        proxy_pass http://backend:3600;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

4. 启动容器:

```bash
docker-compose up -d
```

## 7. 项目初始化

### 7.1 访问管理后台

项目启动后，访问 `http://your_domain.com/manage`，使用默认账号密码登录:

- 用户名: admin
- 密码: admin

### 7.2 配置采集源

1. 登录管理后台
2. 点击"采集管理" -> "采集站点"
3. 添加或修改采集源配置
4. 点击"采集管理" -> "定时任务"，配置定时采集任务

### 7.3 手动触发采集

```bash
# 进入后端目录
cd /path/to/GoFilm/server

# 执行采集命令
./gofilm --collect
```

## 8. 性能优化

### 8.1 PostgreSQL优化

1. **索引优化**: 系统会自动为常用字段创建索引
2. **连接池配置**: 已配置合理的连接池参数
3. **全文搜索**: 使用PostgreSQL全文搜索提高搜索效率
4. **JSONB支持**: 支持JSONB数据类型，优化复杂数据存储

### 8.2 Redis缓存

系统使用Redis缓存以下数据:
- 影片详情信息
- 影片基本信息
- 搜索标签信息
- 播放源信息

## 9. 监控与维护

### 9.1 日志查看

- **后端日志**: `server/logs/app.log`
- **PostgreSQL日志**: `/var/log/postgresql/postgresql-13-main.log` (Ubuntu/Debian)
- **Redis日志**: `/var/log/redis/redis-server.log` (Ubuntu/Debian)

### 9.2 数据库备份

```bash
# 备份数据库
sudo -u postgres pg_dump gofilm > gofilm_backup.sql

# 恢复数据库
sudo -u postgres psql gofilm < gofilm_backup.sql
```

### 9.3 常见问题排查

1. **无法连接数据库**: 检查数据库配置和PostgreSQL服务状态
2. **采集失败**: 检查采集源配置和网络连接
3. **前端无法访问后端API**: 检查Nginx反向代理配置和后端服务状态
4. **影片无法播放**: 检查播放源配置和网络连接

## 10. 安全建议

1. **修改默认密码**: 登录管理后台后立即修改默认密码
2. **配置HTTPS**: 使用Let's Encrypt配置HTTPS
3. **限制数据库访问**: 只允许特定IP访问数据库
4. **定期备份**: 定期备份数据库和配置文件
5. **更新依赖**: 定期更新项目依赖，修复安全漏洞
6. **配置防火墙**: 只开放必要的端口

## 11. 联系方式

- 项目地址: https://github.com/hangyubin/GoFilm
- Issues: https://github.com/hangyubin/GoFilm/issues

---

**部署完成后，您可以通过 `http://your_domain.com` 访问GoFilm网站。**
