-- =========================
-- NeuroCloud 数据库初始化脚本
-- 模块：用户模块
-- 数据库：PostgreSQL
-- =========================

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY, -- 用户ID，自增主键
    username VARCHAR(50) UNIQUE NOT NULL, -- 用户名，唯一
    email VARCHAR(255) UNIQUE NOT NULL, -- 邮箱，唯一
    password VARCHAR(255) NOT NULL, -- 密码哈希
    role VARCHAR(50) DEFAULT 'user', -- 角色：user/admin
    created_at TIMESTAMP DEFAULT NOW(), -- 创建时间
    updated_at TIMESTAMP DEFAULT NOW() -- 更新时间
);

-- 为常用字段加索引
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);

CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);