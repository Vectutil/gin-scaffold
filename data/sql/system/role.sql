CREATE TABLE IF NOT EXISTS `role`
(
    `id`          BIGINT(20)  NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        VARCHAR(50) NOT NULL DEFAULT '' COMMENT '角色名称',
    `code`        VARCHAR(50) NOT NULL DEFAULT '' COMMENT '角色编码',
    `description` VARCHAR(255)         DEFAULT '' COMMENT '描述',
    `data_scope`  TINYINT(1)           DEFAULT 4 COMMENT '数据范围: 1-全部, 2-本部门, 3-本部门及子部门, 4-本人, 5-自定义',
    `status`      TINYINT(1)           DEFAULT 1 COMMENT '状态 1启用 0禁用',
    `created_at`  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by`  BIGINT(20)           DEFAULT NULL COMMENT '创建人',
    `updated_at`  DATETIME             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by`  BIGINT(20)           DEFAULT NULL COMMENT '更新人',
    `deleted_at`  DATETIME             DEFAULT NULL COMMENT '删除时间',
    `deleted_by`  BIGINT(20)           DEFAULT NULL COMMENT '删除人',
    PRIMARY KEY (`id`),
    KEY `code` (`code`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='角色表';
