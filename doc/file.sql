# 文件资源
CREATE TABLE `f_assets`
(
    `id`           char(24) unique NOT NULL COMMENT 'ID',
    `module`       varchar(100)    NOT NULL COMMENT '模块',
    `target`       varchar(100)    NOT NULL COMMENT '对象标识',
    `md5`          varchar(255)    NOT NULL COMMENT '文件md5',
    `key`          varchar(255)    NOT NULL COMMENT '上传key',
    `content_type` varchar(100)    NOT NULL COMMENT '类型',
    `path`         varchar(255)    NOT NULL COMMENT '文件路径',
    `ext`          varchar(20)     NOT NULL COMMENT '文件后缀',
    `state`        tinyint(2)      NOT NULL COMMENT '文件状态 -1=待确认 -2=待禁用 0=禁用 1=启用',
    `created_at`   datetime(3)     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='文件资源表';