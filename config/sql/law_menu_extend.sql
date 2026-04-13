-- 移动卫生执法系统 - 法律法规库、规范用语、监管事项菜单数据
-- 日期：2026-04-13

-- ============================
-- 二级菜单：法律法规库管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(209, '法律法规', 200, 9, 'law/regulation', 'system/regulation/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:regulation:list', 'book', 'admin', NOW(), '', NULL, '法律法规知识库');

-- ============================
-- 二级菜单：规范用语管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(210, '规范用语', 200, 10, 'law/standard-phrase', 'system/standard-phrase/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:standard-phrase:list', 'edit-table', 'admin', NOW(), '', NULL, '执法规范用语');

-- ============================
-- 二级菜单：监管事项管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(211, '监管事项', 200, 11, 'law/supervision-item', 'system/supervision-item/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:supervision-item:list', 'list', 'admin', NOW(), '', NULL, '监管事项管理');

-- ============================
-- 法律法规 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2081, '法律法规查询', 209, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:query', '#', 'admin', NOW(), '', NULL, ''),
(2082, '法律法规新增', 209, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:add', '#', 'admin', NOW(), '', NULL, ''),
(2083, '法律法规修改', 209, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:edit', '#', 'admin', NOW(), '', NULL, ''),
(2084, '法律法规删除', 209, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:remove', '#', 'admin', NOW(), '', NULL, ''),
(2085, '法律法规导出', 209, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:regulation:export', '#', 'admin', NOW(), '', NULL, '');

-- ============================
-- 规范用语 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2091, '规范用语查询', 210, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:query', '#', 'admin', NOW(), '', NULL, ''),
(2092, '规范用语新增', 210, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:add', '#', 'admin', NOW(), '', NULL, ''),
(2093, '规范用语修改', 210, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:edit', '#', 'admin', NOW(), '', NULL, ''),
(2094, '规范用语删除', 210, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:remove', '#', 'admin', NOW(), '', NULL, ''),
(2095, '规范用语导出', 210, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:standard-phrase:export', '#', 'admin', NOW(), '', NULL, '');

-- ============================
-- 监管事项 - 按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(2101, '监管事项查询', 211, 1, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:query', '#', 'admin', NOW(), '', NULL, ''),
(2102, '监管事项新增', 211, 2, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:add', '#', 'admin', NOW(), '', NULL, ''),
(2103, '监管事项修改', 211, 3, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:edit', '#', 'admin', NOW(), '', NULL, ''),
(2104, '监管事项删除', 211, 4, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:remove', '#', 'admin', NOW(), '', NULL, ''),
(2105, '监管事项导出', 211, 5, '#', '', NULL, NULL, 1, 0, 'F', '0', '1', 'law:supervision-item:export', '#', 'admin', NOW(), '', NULL, '');
