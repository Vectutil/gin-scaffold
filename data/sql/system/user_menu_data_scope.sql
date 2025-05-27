CREATE TABLE IF NOT EXISTS `user_menu_data_scope` (
    `user_id` BIGINT(20) NOT NULL,
    `menu_id` BIGINT(20) NOT NULL,
    `dept_id` BIGINT(20) NOT NULL,
    PRIMARY KEY (`user_id`, `menu_id`, `dept_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户在某功能下可访问的部门（个性化权限）';
