CREATE TABLE IF NOT EXISTS `history`(
    `id`                    BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `rid`                   BIGINT NOT NULL AUTO_INCREMENT COMMENT 'rule id',
    `qid`                   BIGINT NOT NULL AUTO_INCREMENT COMMENT 'request id',
    `mode`                  TINYINT NOT NULL DEFAULT '0' COMMENT 'mode: 0-proxy, 1-mock',
    `url`                   VARCHAR(1024) NOT NULL DEFAULT '' COMMENT 'request url',
    `method`                VARCHAR(1024) NOT NULL DEFAULT '' COMMENT 'request method',
    `body`                  VARCHAR(1024) NOT NULL DEFAULT '' COMMENT 'request body',
    `status`                INT NOT NULL DEFAULT '0' COMMENT 'response status',
    `response`              VARCHAR(65535) NOT NULL DEFAULT 'response body',
    `create_time`           INT NOT NULL DEFAULT '0' COMMENT '创建时间(单位秒)',
    `update_time`           INT NOT NULL DEFAULT '0' COMMENT '更新时间(单位秒)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_rid_qid` (`rid`, `qid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='请求历史';
