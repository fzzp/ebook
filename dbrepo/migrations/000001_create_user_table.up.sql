CREATE TABLE IF NOT EXISTS `users` (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    email varchar(60) NOT NULL COMMENT "用户邮箱",
    password_hash varchar(60) NOT NULL COMMENT "哈希密码",
    username varchar(60) NOT NULL COMMENT "用户名称",
    avatar varchar(255) NOT NULL COMMENT "用户头像",
    role tinyint(1) NOT NULL DEFAULT 0 COMMENT "角色;0:客户,1:管理员",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "更新时间",
    INDEX idx_email(email),
    UNIQUE unq_email(email)
) COMMENT="用户表";
