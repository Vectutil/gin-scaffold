CREATE TABLE IF NOT EXISTS `user_menu_data_scope` (
    `id`      BIGINT(20) NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(20) NOT NULL COMMENT '用户ID',
    `menu_id` BIGINT(20) NOT NULL COMMENT '菜单ID',
    `dept_id` BIGINT(20) NOT NULL COMMENT '部门ID',
    PRIMARY KEY (`user_id`, `menu_id`, `dept_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户在某功能下可访问的部门（个性化权限）';
