-- 移动卫生执法系统 - 监管事项菜单数据
-- 版本：v1.0.0
-- 日期：2026-04-13

-- ============================
-- 二级菜单：监管事项管理
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(209, '监管事项', 200, 9, 'law/supervision-item', 'system/supervision-item/index', NULL, NULL, 1, 0, 'C', '0', '0', 'law:supervisionItem:list', 'list', 'admin', NOW(), '', NULL, '');

-- ============================
-- 监管事项管理按钮权限
-- ============================
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(210, '监管事项查询', 209, 1, '', '', NULL, NULL, 1, 0, 'F', '0', '0', 'law:supervisionItem:query', '#', 'admin', NOW(), '', NULL, ''),
(211, '监管事项新增', 209, 2, '', '', NULL, NULL, 1, 0, 'F', '0', '0', 'law:supervisionItem:add', '#', 'admin', NOW(), '', NULL, ''),
(212, '监管事项修改', 209, 3, '', '', NULL, NULL, 1, 0, 'F', '0', '0', 'law:supervisionItem:edit', '#', 'admin', NOW(), '', NULL, ''),
(213, '监管事项删除', 209, 4, '', '', NULL, NULL, 1, 0, 'F', '0', '0', 'law:supervisionItem:remove', '#', 'admin', NOW(), '', NULL, ''),
(214, '监管事项导出', 209, 5, '', '', NULL, NULL, 1, 0, 'F', '0', '0', 'law:supervisionItem:export', '#', 'admin', NOW(), '', NULL, '');

-- ============================
-- 监管事项基础数据示例
-- ============================

-- 公共场所监管事项
INSERT INTO `law_supervision_item` (`item_id`, `item_name`, `parent_id`, `level`, `supervision_type`, `industry_ids`, `standard_language_ids`, `check_points`, `legal_basis`, `sort_order`, `is_enabled`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
(1, '公共场所监管', 0, 1, 'GGCS', NULL, NULL, '公共场所卫生监督执法事项', '《公共场所卫生管理条例》', 1, 1, 'admin', NOW(), '', NULL, ''),
(2, '住宿场所监管', 1, 2, 'GGCS', NULL, NULL, '宾馆、酒店、旅店、招待所等住宿场所的卫生监督执法', '《公共场所卫生管理条例》', 1, 1, 'admin', NOW(), '', NULL, ''),
(3, '美容美发场所监管', 1, 2, 'GGCS', NULL, NULL, '理发店、美容院、美发厅等美容美发场所的卫生监督执法', '《公共场所卫生管理条例》', 2, 1, 'admin', NOW(), '', NULL, ''),
(4, '洗浴场所监管', 1, 2, 'GGCS', NULL, NULL, '浴室、桑拿中心、足浴店等洗浴场所的卫生监督执法', '《公共场所卫生管理条例》', 3, 1, 'admin', NOW(), '', NULL, ''),
(5, '游泳场所监管', 1, 2, 'GGCS', NULL, NULL, '游泳馆、游泳池等游泳场所的卫生监督执法', '《公共场所卫生管理条例》', 4, 1, 'admin', NOW(), '', NULL, '');
