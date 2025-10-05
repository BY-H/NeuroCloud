# 数据库连接，日后改成用注入的方式，防止密码谢罗，本地 Makefile 仅用于测试！
DB_URL=postgres://postgres:123456@localhost:5432/neurocloud?sslmode=disable

.PHONY: migrate-up migrate-down migrate-reset migrate-new
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-reset:
	migrate -path migrations -database "$(DB_URL)" down -all
	migrate -path migrations -database "$(DB_URL)" up

migrate-new:
	migrate create -ext sql -dir migrations $(name)
