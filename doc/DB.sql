-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2019-06-21 05:43:19
-- 服务器版本： 8.0.13
-- PHP 版本： 7.3.5
-- MYSQL版本： 8.0.13 - Source distribution

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

--
-- 数据库： `jasonpsgo`
--

-- --------------------------------------------------------

--
-- 表的结构 `group`
--

CREATE TABLE `group` (
  `id` bigint(20) NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `title` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '2',
  `sort` int(11) NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转存表中的数据 `group`
--

INSERT INTO `group` (`id`, `name`, `title`, `status`, `sort`) VALUES
(1, 'APP', 'System', 2, 1);

-- --------------------------------------------------------

--
-- 表的结构 `node`
--

CREATE TABLE `node` (
  `id` bigint(20) NOT NULL,
  `title` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `level` int(11) NOT NULL DEFAULT '1',
  `pid` bigint(20) NOT NULL DEFAULT '0',
  `remark` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '2',
  `group_id` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转存表中的数据 `node`
--

INSERT INTO `node` (`id`, `title`, `name`, `level`, `pid`, `remark`, `status`, `group_id`) VALUES
(1, 'RBAC', 'rbac', 1, 0, '', 2, 1),
(2, 'Node', 'node/index', 2, 1, '', 2, 1),
(3, 'node list', 'index', 3, 2, '', 2, 1),
(4, 'add or edit', 'AddAndEdit', 3, 2, '', 2, 1),
(5, 'del node', 'DelNode', 3, 2, '', 2, 1),
(6, 'User', 'user/index', 2, 1, '', 2, 1),
(7, 'user list', 'Index', 3, 6, '', 2, 1),
(8, 'add user', 'AddUser', 3, 6, '', 2, 1),
(9, 'update user', 'UpdateUser', 3, 6, '', 2, 1),
(10, 'del user', 'DelUser', 3, 6, '', 2, 1),
(11, 'Group', 'group/index', 2, 1, '', 2, 1),
(12, 'group list', 'index', 3, 11, '', 2, 1),
(13, 'add group', 'AddGroup', 3, 11, '', 2, 1),
(14, 'update group', 'UpdateGroup', 3, 11, '', 2, 1),
(15, 'del group', 'DelGroup', 3, 11, '', 2, 1),
(16, 'Role', 'role/index', 2, 1, '', 2, 1),
(17, 'role list', 'index', 3, 16, '', 2, 1),
(18, 'add or edit', 'AddAndEdit', 3, 16, '', 2, 1),
(19, 'del role', 'DelRole', 3, 16, '', 2, 1),
(20, 'get roles', 'Getlist', 3, 16, '', 2, 1),
(21, 'show access', 'AccessToNode', 3, 16, '', 2, 1),
(22, 'add accsee', 'AddAccess', 3, 16, '', 2, 1),
(23, 'show role to userlist', 'RoleToUserList', 3, 16, '', 2, 1),
(24, 'add role to user', 'AddRoleToUser', 3, 16, '', 2, 1);

-- --------------------------------------------------------

--
-- 表的结构 `node_roles`
--

CREATE TABLE `node_roles` (
  `id` bigint(20) NOT NULL,
  `node_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- 表的结构 `role`
--

CREATE TABLE `role` (
  `id` bigint(20) NOT NULL,
  `title` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `remark` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '2'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转存表中的数据 `role`
--

INSERT INTO `role` (`id`, `title`, `name`, `remark`, `status`) VALUES
(1, 'Admin role', 'Admin', 'I\'m a admin role', 2);

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL,
  `username` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `nickname` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `email` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `remark` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '2',
  `lastlogintime` datetime DEFAULT NULL,
  `createtime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `remark`, `status`, `lastlogintime`, `createtime`) VALUES
(1, 'admin', '21232f297a57a5a743894a0e4a801fc3', 'ClownFish', 'osgochina@gmail.com', 'I\'m admin', 2, NULL, '2019-06-19 04:25:11');

-- --------------------------------------------------------

--
-- 表的结构 `user_roles`
--

CREATE TABLE `user_roles` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转储表的索引
--

--
-- 表的索引 `group`
--
ALTER TABLE `group`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `node`
--
ALTER TABLE `node`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `node_roles`
--
ALTER TABLE `node_roles`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `role`
--
ALTER TABLE `role`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `nickname` (`nickname`);

--
-- 表的索引 `user_roles`
--
ALTER TABLE `user_roles`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `group`
--
ALTER TABLE `group`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `node`
--
ALTER TABLE `node`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- 使用表AUTO_INCREMENT `node_roles`
--
ALTER TABLE `node_roles`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `role`
--
ALTER TABLE `role`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `user_roles`
--
ALTER TABLE `user_roles`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
COMMIT;
