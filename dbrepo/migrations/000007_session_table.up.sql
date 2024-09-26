CREATE TABLE IF NOT EXISTS `sessions` (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    user_id BIGINT UNSIGNED NOT NULL COMMENT "用户主键",
    token_hash varchar(255) NOT NULL COMMENT "会话id",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    INDEX idx_user_id(user_id),
    INDEX idx_token_hash(token_hash),
    UNIQUE unq_token_hash(token_hash)
) COMMENT="会话表";
