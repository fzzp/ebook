CREATE TABLE IF NOT EXISTS `category` (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    category_name varchar(60) NOT NULL COMMENT "分类名",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "更新时间",
    INDEX idx_category_name(category_name),
    UNIQUE unq_category_name(category_name)
) COMMENT="分类表";
