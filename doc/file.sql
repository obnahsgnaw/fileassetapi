# 文件资源
CREATE TABLE `f_assets`
(
    `id`           int(11) unsigned NOT NULL auto_increment COMMENT 'ID',
    `project`      varchar(100)    NOT NULL COMMENT '服务项目',
    `module`       varchar(100)    NOT NULL COMMENT '模块',
    `name`         varchar(100)    NOT NULL COMMENT '路径名称',
    `target`       varchar(100)    NOT NULL COMMENT '对象标识',
    `enabled`      tinyint(2)      NOT NULL COMMENT '文件状态 -1=待确认 0=禁用 1=启用',
    `created_at`   datetime(3)     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='文件资源表';