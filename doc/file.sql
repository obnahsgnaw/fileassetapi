# 文件资源
CREATE TABLE `f_assets`
(
    `id`         char(24) unique NOT NULL COMMENT 'ID',
    `module`     varchar(200)    NOT NULL COMMENT '模块',
    `target`     varchar(100)    NOT NULL COMMENT '对象标识',
    `path`       varchar(255)    NOT NULL COMMENT '文件路径',
    `confirm`    tinyint(1)      NOT NULL COMMENT '是否已确认保存',
    `enabled`    tinyint(2)      NOT NULL COMMENT '启用状态 -1=默认 0=禁用 1=启用',
    `created_at` datetime(3)     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='文件资源表';