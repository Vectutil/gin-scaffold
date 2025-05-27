CREATE TABLE IF NOT EXISTS `user_data_scope` (
    `user_id` BIGINT(20) NOT NULL COMMENT '用户ID',
    `dept_id` BIGINT(20) NOT NULL COMMENT '可访问的部门ID',
    PRIMARY KEY (`user_id`, `dept_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户数据权限扩展表';
