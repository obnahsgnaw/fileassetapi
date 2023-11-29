# 属性配置表
CREATE TABLE `attr_configs`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `target_id`  int(10) unsigned NOT NULL COMMENT '配置属性的对象.id',
    `attr`       varchar(100)     NOT NULL COMMENT '唯一的英文标识码',
    `title`      varchar(100)     NOT NULL COMMENT '标题',
    `value`      varchar(100)     NOT NULL COMMENT '默认值',
    `kind`       tinyint(2)       NOT NULL COMMENT '数据类型',
    `options`    varchar(255)     NOT NULL COMMENT '选项值',
    `enabled`    tinyint(2)       NOT NULL COMMENT '审核状态 -1=默认 0=禁用 1=启用',
    `created_at` datetime(3)      NOT NULL COMMENT '创建时间',
    `updated_at` datetime(3)      NOT NULL COMMENT '更新时间',
    `deleted_at` datetime(3)      NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_target` (`target_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='属性配置表';

# 属性表
CREATE TABLE `attrs`
(
    `target_id`  int(10) unsigned NOT NULL COMMENT '属性对象.id',
    `attr`       varchar(100)     NOT NULL COMMENT '唯一的英文标识码',
    `value`      varchar(100)     NOT NULL COMMENT '名称',
    `created_at` datetime(3)      NOT NULL COMMENT '创建时间',
    KEY `idx_target` (`target_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='属性表';