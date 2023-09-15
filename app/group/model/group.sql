CREATE TABLE `group`
(
    `id`   bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255)    NOT NULL DEFAULT '默认群组' COMMENT '群名称',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `user_group`
(
    `id`        bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`   bigint unsigned NOT NULL COMMENT '用户id',
    `group_id`  bigint unsigned NOT NULL COMMENT '群组id',
    `permission` int             NOT NULL DEFAULT 0 COMMENT '0:群主,1:管理员,2:群员,3:禁言',
    `remark`    varchar(255)    NOT NULL         DEFAULT '' COMMENT '备注',
    PRIMARY KEY (`id`),
    UNIQUE  KEY `uk_user_group` (`user_id`, `group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;