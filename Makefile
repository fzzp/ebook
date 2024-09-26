MIGRATE_SQL=./dbrepo/migrations

# 这里是容器直接通信
DB_SOURCE=mysql://root:ebook_scret@tcp(ebook_mysql:3306)/ebook


# ==================================================================================== #
# docker-compose
# ==================================================================================== #

## docker/up: 在后台运行 docker compose 
docker/build:
	docker-compose build

## docker/start: 直接运行
docker/start:
	docker-compose up 

## docker/up: 在后台运行
docker/up:
	docker-compose up -d

## docker/down: 停止 docker compose 
docker/down:
	docker-compose down

## docker/web: 进入web服务
docker/web: 
	docker compose run --service-ports web bash

## docker/tidy: 安装依赖
docker/tidy:
	docker exec ebook_web go mod tidy

.PHONY: docker/build docker/up docker/start docker/down docker/web


# ==================================================================================== #
# sqlc 生成CRUD代码
# ==================================================================================== #

## sqlc/gen: 生成CRUD代码
sqlc/gen:
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc generate

# ==================================================================================== #
# SQL 迁移
# ==================================================================================== #

## migrate/check 检查migrate是否安装了
migrate/check:
	docker exec ebook_web migrate -version

## migrate/new name=$1: 创建迁移SQL文件，如 make migrate/new name=create_user_table
migrate/new:
	docker exec ebook_web migrate create -seq -ext=.sql -dir=${MIGRATE_SQL} ${name}

## migrate/up: 向上迁移所有
migrate/up:
	docker exec ebook_web migrate -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose up

## migrate/down: 向下迁移所有
migrate/down:
	docker exec ebook_web migrate -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose down

## migrate/up1: 向上迁移一次
migrate/up1: 
	docker exec ebook_web migrate -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose up 1

## migrate/down1: 向下迁移一次
migrate/down1:
	docker exec ebook_web migrate -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose down 1

## migrate/force version=$1: 强制迁移到指定版本
migrate/force:
	docker exec ebook_web migrate -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose force ${version}

## migrate/goto version=$1: 迁移到指定版本
migrate/goto:
	docker exec ebook_web migrate -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose goto ${version}

## migrate/fix version=$1: 向上迁移过程中如果出错，执行此命令快速修复，先强制升迁成功，再向下迁移
# migrate/fix:
# 	make migrate/force version=${version}
# 	make migrate/down1

# migrate/version: 查看当前的迁移版本
migrate/version:
	docker exec ebook_web migrate -path=${MIGRATE_SQL} -database="${DB_SOURCE}" version

.PHONY: migrate/check migrate/new migrate/up migrate/up1 migrate/down migrate/down1 migrate/force migrate/goto migrate/version

