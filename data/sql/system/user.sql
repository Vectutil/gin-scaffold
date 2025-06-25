CREATE TABLE IF NOT EXISTS `user`
(
    `id`                    BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT '主键',
    `username`              VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`              VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
    `full_name`             VARCHAR(255) NOT NULL DEFAULT '' COMMENT '全名',
    `avatar`                VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像URL',
    `email`                 VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone`                 VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '手机号',
    `dept_id`               BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '所属主部门Id',
    `status`                TINYINT(1)   NOT NULL DEFAULT 1 COMMENT '状态：1启用 0禁用',
    `login_count`           INT(11)          NOT NULL DEFAULT 0 COMMENT '登录次数',
    `last_login_at`         INT(11)     NOT NULL DEFAULT 0 COMMENT '最后登录时间',
    `last_login_ip`         VARCHAR(45)  NOT NULL DEFAULT '' COMMENT '最后登录IP地址',
    `tenant_id`             BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '租户Id',
    `org_id`                BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '组织Id(暂时没能力构建)',
    `remark`                TEXT         NOT NULL COMMENT '备注信息',
    `created_at`            DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by`            BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '创建人Id',
    `updated_at`            DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by`            BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '更新人Id',
    `deleted_at`            DATETIME              DEFAULT NULL COMMENT '删除时间',
    `deleted_by`            BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '删除人Id',
    PRIMARY KEY (`id`),
    KEY `idx_dept_id` (`dept_id`),
    KEY `idx_tenant_id` (`tenant_id`),
    KEY `idx_deleted_status` (`deleted_at`, `status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';


