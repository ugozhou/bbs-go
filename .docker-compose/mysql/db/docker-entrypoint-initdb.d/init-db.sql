
SET NAMES UTF8;

-- 初始化用户表
CREATE TABLE IF NOT EXISTS `t_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(32) COLLATE utf8_unicode_ci DEFAULT NULL,
  `nickname` varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
  `avatar` text COLLATE utf8_unicode_ci,
  `password` varchar(512) COLLATE utf8_unicode_ci DEFAULT NULL,
  `paypassword` varchar(512) COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` int(11) NOT NULL,
  `invitecode` varchar(512) COLLATE utf8_unicode_ci NOT NULL,
  `roles` text COLLATE utf8_unicode_ci,
  `type` int(11) NOT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `introducer` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `mobile` (`mobile`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 初始化用户数据（用户名：admin、密码：123456）
INSERT INTO t_user (
    `id`,
    `mobile`,
    `nickname`,
    `avatar`,
    `password`,
    `paypassword`,
    `status`,
    `invitecode`,
    `create_time`,
    `update_time`,
    `roles`,
    `type`,
    `introducer`
) SELECT
    1,
    ''18910254412'',
    ''ugozhou'',
    '''',
    ''$2a$10$ofA39bAFMpYpIX/Xiz7jtOMH9JnPvYfPRlzHXqAtLPFpbE/cLdjmS'',
    ''$2a$10$ofA39bAFMpYpIX/Xiz7jtOMH9JnPvYfPRlzHXqAtLPFpbE/cLdjmS'',
    0,
    ''xxx'',
    1555419028975,
    1555419028975,
    ''管理员'',
    0,
    1
FROM
    DUAL
WHERE
    NOT EXISTS (
        SELECT
    `id`,
    `mobile`,
    `nickname`,
    `avatar`,
    `password`,
    `paypassword`,
    `status`,
    `invitecode`,
    `create_time`,
    `update_time`,
    `roles`,
    `type`,
    `introducer`
        FROM t_user
        WHERE id = 1
    );

-- 初始化系统配置表
CREATE TABLE IF NOT EXISTS `t_sys_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `key` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `value` text COLLATE utf8_unicode_ci,
  `name` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 初始化系统配置数据
INSERT INTO t_sys_config(
    `key`,
    `value`,
    `name`,
    `description`,
    `create_time`,
    `update_time`
) SELECT
    ''siteTitle'',
    ''wkycloud'',
    ''wkyapp'',
    ''站点标题'',
    1555419028975,
    1555419028975
FROM
    DUAL
WHERE
    NOT EXISTS (
        SELECT
            `key`,
            `value`,
            `name`,
            `description`,
            `create_time`,
            `update_time`
        FROM
            t_sys_config
        WHERE
            `key` = ''siteTitle''
    );

INSERT INTO t_sys_config (
    `key`,
    `value`,
    `name`,
    `description`,
    `create_time`,
    `update_time`
) SELECT
    ''siteDescription'',
    ''wkycloud，基于Go语言的开源社区系统'',
    ''站点描述'',
    ''站点描述'',
    1555419028975,
    1555419028975
FROM
    DUAL
WHERE
    NOT EXISTS (
        SELECT
            `key`,
            `value`,
            `name`,
            `description`,
            `create_time`,
            `update_time`
        FROM
            t_sys_config
        WHERE
            `key` = ''siteDescription''
    );

INSERT INTO t_sys_config (
    `key`,
    `value`,
    `name`,
    `description`,
    `create_time`,
    `update_time`
) SELECT
    ''siteKeywords'',
    ''bbs-go'',
    ''站点关键字'',
    ''站点关键字'',
    1555419028975,
    1555419028975
FROM
    DUAL
WHERE
    NOT EXISTS (
        SELECT
            `key`,
            `value`,
            `name`,
            `description`,
            `create_time`,
            `update_time`
        FROM
            t_sys_config
        WHERE
            `key` = ''siteKeywords''
    );


INSERT INTO t_sys_config (
    `key`,
    `value`,
    `name`,
    `description`,
    `create_time`,
    `update_time`
  )
SELECT
  ''siteNavs'',
  ''[{\"title\":\"首页\",\"url\":\"/\"},{\"title\":\"话题\",\"url\":\"/topics\"},{\"title\":\"文章\",\"url\":\"/articles\"}]'',
  ''站点导航'',
  ''站点导航'',
  1555419028975,
  1555419028975
FROM DUAL
WHERE
  NOT EXISTS (
    SELECT
      `key`,
      `value`,
      `name`,
      `description`,
      `create_time`,
      `update_time`
    FROM t_sys_config
    WHERE
      `key` = ''siteNavs''
  );