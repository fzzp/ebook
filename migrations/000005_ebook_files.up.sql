CREATE TABLE IF NOT EXISTS `ebook_files` (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    book_id BIGINT UNSIGNED NOT NULL COMMENT "图书主键",
    file_type char(6) NOT NULL COMMENT "文件类型",
    download_url varchar(255) NOT NULL COMMENT "下载地址"
) COMMENT="电子书文件表";
