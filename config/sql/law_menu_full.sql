-- 移动卫生执法系统 - 完整菜单数据 SQL
-- 日期：2026-04-13
-- 说明：包含卫生执法系统所有业务菜单

-- ============================
-- 一级菜单：卫生执法
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(200, '卫生执法', 0, 2, 'law', NULL, NULL, NULL, 1, 0, 'M', '0', '0', NULL, 'star', 'admin', NOW(), '', NULL, '卫生执法系统')
ON DUPLICATE KEY UPDATE `menu_name`='卫生执法';

-- ============================
-- 二级菜单：行业分类管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(201, '行业分类', 200, 1, 'law/industry', 'system/industry/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:industry:list', 'tree', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`='行业分类', `component`='system/industry/index', `perms`='law:industry:list', `icon`='tree';

-- ============================
-- 二级菜单：监管单位管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(202, '监管单位', 200, 2, 'law/subject', 'system/subject/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:subject:list', 'user', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`='监管单位', `component`='system/subject/index', `perms`='law:subject:list', `icon`='user';

-- ============================
-- 二级菜单：执法人员管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(203, '执法人员', 200, 3, 'law/official', 'system/official/index', NULL, NULL, 1, 0, 'C', '0', '1', 'law:official:list', 'peoples', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`='执法人员', `component`='system/official/index', `perms`='law:official:list', `icon`='peoples';

-- ============================
-- 二级菜单：设备管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(204, '设备管理', 200, 4, 'law/device', 'system/device/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:device:list', 'phone', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`='设备管理', `component`='system/device/index', `perms`='law:device:list', `icon`='phone';

-- ============================
-- 二级菜单：激活码管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(205, '激活码', 200, 5, 'law/activate-code', 'system/activate-code/index', NULL, NULL, 1, 0, 'C', '0', '1', 'law:activate-code:list', 'validCode', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`='激活码', `component`='system/activate-code/index', `perms`='law:activate-code:list', `icon`='validCode';

-- ============================
-- 二级菜单：文书模板管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(206, '文书模板', 200, 6, 'law/template', 'system/template/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:template:list', 'edit', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`='文书模板', `component`='system/template/index', `perms`='law:template:list', `icon`='edit';

-- ============================
-- 二级菜单：执法记录管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(207, '执法记录', 200, 7, 'law/record', 'system/record/index', NULL, NULL, 1, 0, 'C', '0', '1', 'law:record:list', 'form', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`='执法记录', `component`='system/record/index', `perms`='law:record:list', `icon`='form';

-- ============================
-- 二级菜单：数据同步管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(208, '数据同步', 200, 8, 'law/sync', 'system/sync/index', NULL, NULL, 1, 0, 'C', '0', '1', 'law:sync:list', 'link', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`='数据同步', `component`='system/sync/index', `perms`='law:sync:list', `icon`='link';

-- ============================
-- 二级菜单：法律法规库管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(209, '法律法规', 200, 9, 'law/regulation', 'system/regulation/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:regulation:list', 'book', 'admin', NOW(), '', NULL, '法律法规知识库')
ON DUPLICATE KEY UPDATE `menu_name`='法律法规', `component`='system/regulation/index', `perms`='law:regulation:list', `icon`='book';

-- ============================
-- 二级菜单：规范用语管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(210, '规范用语', 200, 10, 'law/standard-phrase', 'system/standard-phrase/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:standard-phrase:list', 'edit-table', 'admin', NOW(), '', NULL, '执法规范用语')
ON DUPLICATE KEY UPDATE `menu_name`='规范用语', `component`='system/standard-phrase/index', `perms`='law:standard-phrase:list', `icon`='edit-table';

-- ============================
-- 二级菜单：监管事项管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(211, '监管事项', 200, 11, 'law/supervision-item', 'system/supervision-item/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:supervision-item:list', 'list', 'admin', NOW(), '', NULL, '监管事项管理')
ON DUPLICATE KEY UPDATE `menu_name`='监管事项', `component`='system/supervision-item/index', `perms`='law:supervision-item:list', `icon`='list';

-- ============================
-- 行业分类 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2001, '行业分类查询', 201, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:industry:query', '#', 'admin', NOW(), '', NULL, ''),
(2002, '行业分类新增', 201, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:industry:add', '#', 'admin', NOW(), '', NULL, ''),
(2003, '行业分类修改', 201, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:industry:edit', '#', 'admin', NOW(), '', NULL, ''),
(2004, '行业分类删除', 201, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:industry:remove', '#', 'admin', NOW(), '', NULL, ''),
(2005, '行业分类导出', 201, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:industry:export', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 监管单位 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2011, '监管单位查询', 202, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:subject:query', '#', 'admin', NOW(), '', NULL, ''),
(2012, '监管单位新增', 202, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:subject:add', '#', 'admin', NOW(), '', NULL, ''),
(2013, '监管单位修改', 202, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:subject:edit', '#', 'admin', NOW(), '', NULL, ''),
(2014, '监管单位删除', 202, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:subject:remove', '#', 'admin', NOW(), '', NULL, ''),
(2015, '监管单位导出', 202, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:subject:export', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 执法人员 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2021, '执法人员查询', 203, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:official:query', '#', 'admin', NOW(), '', NULL, ''),
(2022, '执法人员新增', 203, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:official:add', '#', 'admin', NOW(), '', NULL, ''),
(2023, '执法人员修改', 203, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:official:edit', '#', 'admin', NOW(), '', NULL, ''),
(2024, '执法人员删除', 203, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:official:remove', '#', 'admin', NOW(), '', NULL, ''),
(2025, '执法人员导出', 203, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:official:export', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 设备管理 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2031, '设备查询', 204, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:device:query', '#', 'admin', NOW(), '', NULL, ''),
(2032, '设备新增', 204, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:device:add', '#', 'admin', NOW(), '', NULL, ''),
(2033, '设备修改', 204, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:device:edit', '#', 'admin', NOW(), '', NULL, ''),
(2034, '设备删除', 204, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:device:remove', '#', 'admin', NOW(), '', NULL, ''),
(2035, '设备导出', 204, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:device:export', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 激活码 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2041, '激活码查询', 205, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:activate-code:query', '#', 'admin', NOW(), '', NULL, ''),
(2042, '激活码新增', 205, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:activate-code:add', '#', 'admin', NOW(), '', NULL, ''),
(2043, '激活码修改', 205, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:activate-code:edit', '#', 'admin', NOW(), '', NULL, ''),
(2044, '激活码删除', 205, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:activate-code:remove', '#', 'admin', NOW(), '', NULL, ''),
(2045, '激活码导出', 205, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:activate-code:export', '#', 'admin', NOW(), '', NULL, ''),
(2046, '激活码生成', 205, 6, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:activate-code:generate', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 文书模板 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2051, '文书模板查询', 206, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:template:query', '#', 'admin', NOW(), '', NULL, ''),
(2052, '文书模板新增', 206, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:template:add', '#', 'admin', NOW(), '', NULL, ''),
(2053, '文书模板修改', 206, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:template:edit', '#', 'admin', NOW(), '', NULL, ''),
(2054, '文书模板删除', 206, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:template:remove', '#', 'admin', NOW(), '', NULL, ''),
(2055, '文书模板上传', 206, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:template:upload', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 执法记录 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2061, '执法记录查询', 207, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:record:query', '#', 'admin', NOW(), '', NULL, ''),
(2062, '执法记录新增', 207, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:record:add', '#', 'admin', NOW(), '', NULL, ''),
(2063, '执法记录修改', 207, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:record:edit', '#', 'admin', NOW(), '', NULL, ''),
(2064, '执法记录删除', 207, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:record:remove', '#', 'admin', NOW(), '', NULL, ''),
(2065, '执法记录导出', 207, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:record:export', '#', 'admin', NOW(), '', NULL, ''),
(2066, '执法记录上报', 207, 6, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:record:submit', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 数据同步 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2071, '数据同步查询', 208, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:sync:query', '#', 'admin', NOW(), '', NULL, ''),
(2072, '数据同步重试', 208, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:sync:retry', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 法律法规 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2081, '法律法规查询', 209, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:query', '#', 'admin', NOW(), '', NULL, ''),
(2082, '法律法规新增', 209, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:add', '#', 'admin', NOW(), '', NULL, ''),
(2083, '法律法规修改', 209, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:edit', '#', 'admin', NOW(), '', NULL, ''),
(2084, '法律法规删除', 209, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:remove', '#', 'admin', NOW(), '', NULL, ''),
(2085, '法律法规导出', 209, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:export', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 规范用语 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2091, '规范用语查询', 210, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:query', '#', 'admin', NOW(), '', NULL, ''),
(2092, '规范用语新增', 210, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:add', '#', 'admin', NOW(), '', NULL, ''),
(2093, '规范用语修改', 210, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:edit', '#', 'admin', NOW(), '', NULL, ''),
(2094, '规范用语删除', 210, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:remove', '#', 'admin', NOW(), '', NULL, ''),
(2095, '规范用语导出', 210, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:export', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;

-- ============================
-- 监管事项 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2101, '监管事项查询', 211, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:query', '#', 'admin', NOW(), '', NULL, ''),
(2102, '监管事项新增', 211, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:add', '#', 'admin', NOW(), '', NULL, ''),
(2103, '监管事项修改', 211, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:edit', '#', 'admin', NOW(), '', NULL, ''),
(2104, '监管事项删除', 211, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:remove', '#', 'admin', NOW(), '', NULL, ''),
(2105, '监管事项导出', 211, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:export', '#', 'admin', NOW(), '', NULL, '')
ON DUPLICATE KEY UPDATE `menu_name`=`menu_name`;
