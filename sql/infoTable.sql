create database `systool`;
use `systool`;
create table `sysinfo`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `tt` INT NOT NULL,
    `cpu` INT NOT NULL,
    `mem` INT NOT NULL,
    `net` INT NOT NULL,
    `last_modify_date` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY(`id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT ="统计信息表";
