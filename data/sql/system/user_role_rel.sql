CREATE TABLE IF NOT EXISTS `user_role` (
        `user_id`     BIGINT(20)     NOT NULL COMMENT '用户ID',
        `role_id`     BIGINT(20)     NOT NULL COMMENT '角色ID',
    PRIMARY KEY (`user_id`, `role_id`),
    KEY `role_id` (`role_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关系表';
