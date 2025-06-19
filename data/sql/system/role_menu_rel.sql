CREATE TABLE IF NOT EXISTS `role_menu_rel` (
    `tenant_id`   BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '租户ID',
    `role_id`     BIGINT(20)     NOT NULL COMMENT '角色ID',
    `menu_id`     BIGINT(20)     NOT NULL COMMENT '菜单ID（按钮/菜单）',
    `scope_type`  TINYINT(1)     NOT NULL DEFAULT 4 COMMENT '数据范围: 1-全部, 2-本部门, 3-本部门及子部门, 4-本人',
    PRIMARY KEY (`tenant_id`, `role_id`, `menu_id`),
    KEY `idx_tenant_role` (`tenant_id`, `role_id`),
    KEY `idx_tenant_menu` (`tenant_id`, `menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限映射表';
