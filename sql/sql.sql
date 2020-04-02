-- 用户
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(25) COLLATE utf8mb4_bin DEFAULT NULL,
  `password` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `salt` char(4) COLLATE utf8mb4_bin NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `profession` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_users_email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 文章
CREATE TABLE `articles` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `introduction` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `content_md` text COLLATE utf8mb4_bin DEFAULT NULL,
  `content_html` text COLLATE utf8mb4_bin DEFAULT NULL,
  `user_id` int(10) NOT NULL DEFAULT 0,
  `tags` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `view_num` int(11) NOT NULL DEFAULT 0,
  `directory_html` text COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 标签
CREATE TABLE `tags` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(25) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `use_num` int(10) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 归档
create table archives (
    id int(10) primary key auto_increment,
    archive_date varchar(25) not null comment '归档日期',
    article_ids varchar(255) not null default '' comment '文章ids',
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted_at timestamp
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;