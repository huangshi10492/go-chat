CREATE TABLE `user`
(
    `id`          bigint unsigned     NOT NULL AUTO_INCREMENT,
    `name`        varchar(100)        NOT NULL DEFAULT '默认用户' COMMENT '用户名称',
    `gender`      tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
    `avatar`      varchar(100)        NOT NULL DEFAULT '' COMMENT '用户头像',
    `mobile`      varchar(50)         NOT NULL DEFAULT '' COMMENT '电话号码',
    `email`       varchar(255)        NULL     DEFAULT '' COMMENT '邮箱',
    `password`    varchar(255)        NOT NULL DEFAULT '' COMMENT '用户密码',
    `create_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_mobile_unique` (`mobile`),
    INDEX `idx_mobile` (`mobile`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
