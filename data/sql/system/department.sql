CREATE TABLE IF NOT EXISTS `department`
(
    `id`         BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT '部门Id',
    `name`  VARCHAR(255) NOT NULL COMMENT '部门名称',
    `parent_id`  BIGINT(20)    NOT NULL DEFAULT 0 COMMENT '上级部门Id，NULL 表示顶级',
    `level`     TINYINT(1)   NOT NULL DEFAULT 1 COMMENT '深度',
    `status`     TINYINT(1)   NOT NULL DEFAULT 1 COMMENT '状态：1启用 0禁用',
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by` BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '创建人Id',
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by` BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '更新人Id',
    `deleted_at` DATETIME              DEFAULT NULL COMMENT '删除时间',
    `deleted_by` BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '删除人Id',

    PRIMARY KEY (`id`),
    INDEX `idx_tenant_parent` (`tenant_id`, `parent_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COMMENT = '部门表';
