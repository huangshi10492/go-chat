CREATE TABLE `friend`
(
    `id`        bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`   bigint unsigned NOT NULL,
    `friend_id` bigint unsigned NOT NULL,
    `remark`    varchar(255)    NOT NULL DEFAULT '' COMMENT '好友备注',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_id` (`user_id`, `friend_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `friend_request`
(
    `id`        bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`   bigint unsigned NOT NULL,
    `friend_id` bigint unsigned NOT NULL,
    `reason`    varchar(255)    NOT NULL DEFAULT '' COMMENT '申请原因',
    `remark`    varchar(255)    NOT NULL DEFAULT '' COMMENT '好友请求备注',
    `status`     int             NOT NULL DEFAULT 0 COMMENT '0: 未处理, 1: 同意, 2: 拒绝',
    PRIMARY KEY (`id`),
    INDEX `idx_user_id` (`friend_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;