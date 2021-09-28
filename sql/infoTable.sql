create database `systool`;
use `systool`;
create table `sysinfo`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `cpu_usr` INT NOT NULL,
    `cpu_sys` INT NOT NULL,
    `cpu_idl` INT NOT NULL,
    `cpu_wai` INT NOT NULL,
    `cpu_stl` INT NOT NULL,
    `mem_used` INT NOT NULL,
    `mem_free` INT NOT NULL,
    `mem_buff` INT NOT NULL,
    `mem_cach` INT NOT NULL,
    `net_recv` INT NOT NULL,
    `net_send` INT NOT NULL,
    `last_modify_date` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY(`id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT ="统计信息表";
