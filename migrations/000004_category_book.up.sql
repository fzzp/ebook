CREATE TABLE IF NOT EXISTS `category_book` (
    id BIGINT UNSIGNED NOT NULL COMMENT "分类主键",
    book_id BIGINT UNSIGNED NOT NULL COMMENT "图书主键"
) COMMENT="分类图书关系表";
