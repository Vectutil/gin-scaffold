CREATE TABLE IF NOT EXISTS `role_permission` (
    `role_id`     BIGINT(20)     NOT NULL COMMENT '角色ID',
    `menu_id`     BIGINT(20)     NOT NULL COMMENT '菜单ID（按钮/菜单）',
    `scope_type` VARCHAR(20) NOT NULL DEFAULT 'SELF' COMMENT 'SELF | DEPT | DEPT_AND_CHILD | ALL | CUSTOM',
    PRIMARY KEY (`role_id`, `menu_id`),
    KEY `menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限映射表';
