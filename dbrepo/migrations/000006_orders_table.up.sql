CREATE TABLE IF NOT EXISTS `orders` (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    user_id BIGINT UNSIGNED NOT NULL COMMENT "用户主键",
    book_id BIGINT UNSIGNED NOT NULL COMMENT "图书主键",
    order_type tinyint(1) NOT NULL COMMENT "1:电子书,2:实体书",
    order_amount int UNSIGNED NOT NULL COMMENT "总金额，单位分",
    order_status tinyint(1) NOT NULL DEFAULT 0 COMMENT "支付状态: 0未支付 1已支付",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "更新时间"
) COMMENT="订单表";
