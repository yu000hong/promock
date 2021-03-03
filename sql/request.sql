CREATE TABLE IF NOT EXISTS `request`(
    `id`                    BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `rid`                   BIGINT NOT NULL AUTO_INCREMENT COMMENT 'rule id',
    `path`                  VARCHAR(1024) NOT NULL DEFAULT '*' COMMENT 'request path',
    `mode`                  TINYINT NOT NULL DEFAULT '0' COMMENT 'mode: 0-proxy, 1-mock',
    `status`                INT NOT NULL DEFAULT '0' COMMENT 'response status',
    `response`              VARCHAR(65535) NOT NULL DEFAULT 'response body',
    `enabled`               TINYINT NOT NULL DEFAULT '0' COMMENT 'enabled: 0-禁用，1-启用',
    `create_time`           INT NOT NULL DEFAULT '0' COMMENT '创建时间(单位秒)',
    `update_time`           INT NOT NULL DEFAULT '0' COMMENT '更新时间(单位秒)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_rid_url` (`rid`,`url`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='请求';
