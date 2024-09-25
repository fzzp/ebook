CREATE TABLE IF NOT EXISTS `books` (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    book_title varchar(255) NOT NULL COMMENT "书名",
    book_isbn char(13) NOT NULL COMMENT "图片编号",
    book_cover varchar(255) NOT NULL COMMENT "封面",
    book_author varchar(60) NOT NULL COMMENT "作者",
    book_price int UNSIGNED NOT NULL COMMENT "实体书价格，单位分",
    book_inventory int UNSIGNED NOT NULL COMMENT "实体书库存",
    ebook_price int UNSIGNED NOT NULL COMMENT "电子书价格,单位分",
    publish_at TIMESTAMP NOT NULL COMMENT "发布年份",
    version int NOT NULL DEFAULT 1 COMMENT "版本号",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "更新时间",
    INDEX idx_book_title(book_title),
    INDEX idx_book_isbn(book_isbn),
    UNIQUE unq_book_isbn(book_isbn)
) COMMENT="图书表";
