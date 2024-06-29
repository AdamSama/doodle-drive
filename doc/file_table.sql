CREATE TABLE `tbl_file` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `file_sha1` char(40) NOT NULL DEFAULT '' COMMENT 'File hash',
    `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT 'File name',
    `file_size` bigint(20) DEFAULT '0' COMMENT 'File size',
    `file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT 'File storage address',
    `create_at` datetime DEFAULT NOW() COMMENT 'Creation date',
    `update_at` datetime DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update date',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT 'Status (active/disabled/deleted, etc.)',
    `ext1` int(11) DEFAULT '0' COMMENT 'Extension field 1',
    `ext2` text COMMENT 'Extension field 2',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_hash` (`file_sha1`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
