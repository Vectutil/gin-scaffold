CREATE TABLE IF NOT EXISTS `user_role_rel`
(
    `tenant_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
    `user_id`   BIGINT(20) NOT NULL COMMENT '用户ID',
    `role_id`   BIGINT(20) NOT NULL COMMENT '角色ID',
    PRIMARY KEY (`tenant_id`, `user_id`, `role_id`),
    KEY `idx_tenant_user` (`tenant_id`, `user_id`),
    KEY `idx_tenant_role` (`tenant_id`, `role_id`),
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户角色关系表';
