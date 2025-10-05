# NeuroCloud

一个 AI 训练平台

## 架构

采用 go-zero + golang-migrate + postgresql + milvus + redis

## 开发事项

### 数据库迁移

因 `golang-migrate` 无法执行建库语句，所以，从 `./sql` 中，执行对应的建库语句，再进行数据库迁移。

使用 `golang-migrate` 进行数据库迁移，在 `./migrations/` 下编写数据库迁移语句，使用 `make migrate-new {file_name}` 即可创建对应迁移文件

直接使用 `make migrate-up` 即可将进行迁移

### 微服务开发

采用 `go-zero` 编写相应的 `api` 和 `rpc` 服务， 使用 `go-ctl` 根据对应的 `.api` 和 `.proto` 文件生成相应的代码

### services 的使用

对于 `go-zero` 没办法直接使用的 Service （如 `Milvus`、`redis` 等），本项目则将相应的集成在 `./services` 中，各个微服务自行使用

### pkg 的使用

对于无需持久化使用的服务，则在 `./pkg` 中进行编写