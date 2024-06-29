-- master/init.sql
CREATE USER 'repluser'@'%' IDENTIFIED BY 'replpassword';
GRANT REPLICATION SLAVE ON *.* TO 'repluser'@'%';
FLUSH PRIVILEGES;

-- Set the binary log format and enable logging
SET GLOBAL binlog_format = 'ROW';
RESET MASTER;
