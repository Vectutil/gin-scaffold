CREATE TABLE IF NOT EXISTS `job_mq`
(
    `id`         bigint(20)       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `obj_id`     bigint(20)       NOT NULL DEFAULT 0 COMMENT '要实现的id',
    `target_id`  varchar(255) NOT NULL DEFAULT '' COMMENT '目标id',
    `type`       tinyint(1)      NOT NULL DEFAULT 0 COMMENT '类型',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `deleted_at` datetime              DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY          `idx_obj_id` (`obj_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息队列相关任务表'; 