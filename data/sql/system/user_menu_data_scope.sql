CREATE TABLE IF NOT EXISTS `user_menu_data_scope` (
    `id`        BIGINT(20) NOT NULL AUTO_INCREMENT,
    `tenant_id` BIGINT(20) NOT NULL COMMENT '租户id',
    `user_id`   BIGINT(20) NOT NULL COMMENT '用户Id',
    `menu_id`   BIGINT(20) NOT NULL COMMENT '菜单Id',
    `dept_ids`  json NOT NULL COMMENT '部门Ids',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_user_menu` (`user_id`, `menu_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_menu_id` (`menu_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户在某功能下可访问的部门（个性化权限）';