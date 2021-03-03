CREATE TABLE IF NOT EXISTS `app`(
    `id`                    BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `host`                  VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'app host',
    `name`                  VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'app name',
    `create_time`           INT NOT NULL DEFAULT '0' COMMENT '创建时间(单位秒)',
    `update_time`           INT NOT NULL DEFAULT '0' COMMENT '更新时间(单位秒)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_key` (`host`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用';
