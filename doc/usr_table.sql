CREATE TABLE `tbl_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'Username',
    `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT 'User encoded password',
    `email` varchar(64) DEFAULT '' COMMENT 'Email',
    `phone` varchar(128) DEFAULT '' COMMENT 'Phone number',
    `email_validated` tinyint(1) DEFAULT 0 COMMENT 'Email validated',
    `phone_validated` tinyint(1) DEFAULT 0 COMMENT 'Phone number validated',
    `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'Signup date',
    `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last active timestamp',
    `profile` text COMMENT 'User profile',
    `status` int(11) NOT NULL DEFAULT 0 COMMENT 'Account status (enabled/disabled/locked/marked for deletion, etc.)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_phone` (`phone`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
