CREATE TABLE IF NOT EXISTS `user_role_rel`
(
    `id`        BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `tenant_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '租户ID',
    `user_id`   BIGINT(20) NOT NULL COMMENT '用户ID',
    `role_id`   BIGINT(20) NOT NULL COMMENT '角色ID',
    PRIMARY KEY (`id`),
    KEY `tenant_id` (`tenant_id`),
    KEY `role_id` (`role_id`),
    KEY `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户角色关系表';
