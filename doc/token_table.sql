CREATE TABLE `tbl_user_token` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'Username',
    `user_token` char(40) NOT NULL DEFAULT '' COMMENT 'User login token',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
