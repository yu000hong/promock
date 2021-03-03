CREATE TABLE IF NOT EXISTS `rule`(
    `id`                    BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `appid`                 BIGINT NOT NULL AUTO_INCREMENT COMMENT 'app id',
    `version`               VARCHAR(20) NOT NULL DEFAULT '*' COMMENT 'version',
    `platform`              VARCHAR(20) NOT NULL DEFAULT '*' COMMENT 'platform',
    `device`                VARCHAR(80) NOT NULL DEFAULT '*' COMMENT 'device',
    `uid`                   VARCHAR(20) NOT NULL DEFAULT '*' COMMENT 'uid',
    `enabled`               TINYINT NOT NULL DEFAULT '0' COMMENT 'enabled: 0-禁用，1-启用',
    `create_time`           INT NOT NULL DEFAULT '0' COMMENT '创建时间(单位秒)',
    `update_time`           INT NOT NULL DEFAULT '0' COMMENT '更新时间(单位秒)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_appid` (`appid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='规则';
