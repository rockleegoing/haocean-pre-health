-- 规范用语库 - 数据库迁移脚本
-- 版本：v1.0.0
-- 日期：2026-04-13
-- 说明：4 级结构（监管类型 → 规范类别 → 规范条目 → 规范内容）

-- 设置外键
SET FOREIGN_KEY_CHECKS = 0;

-- ============================
-- 1. 监管类型表（12 类监管类型）
-- ============================
DROP TABLE IF EXISTS `law_standard_phrase_supervision_type`;
CREATE TABLE `law_standard_phrase_supervision_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '监管类型 ID',
  `name` varchar(100) NOT NULL COMMENT '监管类型名称',
  `code` varchar(50) DEFAULT NULL COMMENT '监管类型代码',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标',
  `description` varchar(500) DEFAULT NULL COMMENT '类型描述',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用（0:禁用/1:启用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_code` (`code`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='规范用语 - 监管类型表';

-- ============================
-- 2. 规范类别表（如：监督用语、处罚用语等）
-- ============================
DROP TABLE IF EXISTS `law_standard_phrase_category`;
CREATE TABLE `law_standard_phrase_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '规范类别 ID',
  `supervision_type_id` bigint(20) DEFAULT NULL COMMENT '监管类型 ID',
  `name` varchar(100) NOT NULL COMMENT '规范类别名称',
  `code` varchar(50) DEFAULT NULL COMMENT '规范类别代码',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_supervision_type_id` (`supervision_type_id`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='规范用语 - 规范类别表';

-- ============================
-- 3. 规范条目表（具体的规范条目）
-- ============================
DROP TABLE IF EXISTS `law_standard_phrase_item`;
CREATE TABLE `law_standard_phrase_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '规范条目 ID',
  `category_id` bigint(20) DEFAULT NULL COMMENT '规范类别 ID',
  `title` varchar(200) NOT NULL COMMENT '条目标题',
  `scene` varchar(200) DEFAULT NULL COMMENT '适用场景',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='规范用语 - 规范条目表';

-- ============================
-- 4. 规范内容表（详细的规范表达内容）
-- ============================
DROP TABLE IF EXISTS `law_standard_phrase_content`;
CREATE TABLE `law_standard_phrase_content` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '规范内容 ID',
  `item_id` bigint(20) DEFAULT NULL COMMENT '规范条目 ID',
  `content` text NOT NULL COMMENT '规范内容',
  `legal_basis` text COMMENT '法律依据',
  `tips` varchar(500) DEFAULT NULL COMMENT '提示要点',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_item_id` (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='规范用语 - 规范内容表';

-- 恢复外键检查
SET FOREIGN_KEY_CHECKS = 1;

-- ============================
-- 初始化数据 - 12 类监管类型
-- ============================
INSERT INTO `law_standard_phrase_supervision_type` (`id`, `name`, `code`, `icon`, `description`, `sort_order`, `is_enabled`) VALUES
(1, '公共场所卫生', 'PUBLIC', 'public-health', '宾馆、美容美发、洗浴、游泳等场所卫生监管', 1, 1),
(2, '生活饮用水卫生', 'WATER', 'water', '供水单位、涉水产品卫生监管', 2, 1),
(3, '学校卫生', 'SCHOOL', 'school', '学校教学环境、饮用水、传染病防控监管', 3, 1),
(4, '医疗机构卫生', 'MEDICAL', 'medical', '医院、诊所等医疗机构执业监管', 4, 1),
(5, '职业卫生', 'OCCUPATION', 'occupation', '职业病危害因素检测、职业健康检查监管', 5, 1),
(6, '放射卫生', 'RADIATION', 'radiation', '放射诊疗、放射防护检测监管', 6, 1),
(7, '传染病防治', 'INFECTION', 'infection', '传染病疫情报告、防控措施落实监管', 7, 1),
(8, '消毒产品', 'DISINFECTANT', 'disinfectant', '消毒剂、消毒器械卫生监管', 8, 1),
(9, '涉水产品', 'WATER_PRODUCT', 'water-product', '涉及饮用水卫生安全产品监管', 9, 1),
(10, '妇幼保健', 'MATERNAL', 'maternal', '妇幼保健机构执业监管', 10, 1),
(11, '计划生育', 'FAMILY_PLAN', 'family-plan', '计划生育技术服务监管', 11, 1),
(12, '中医诊所', 'TCM_CLINIC', 'tcm', '中医诊所备案、执业监管', 12, 1);

-- ============================
-- 初始化数据 - 规范类别（以公共场所为例）
-- ============================
INSERT INTO `law_standard_phrase_category` (`id`, `supervision_type_id`, `name`, `code`, `sort_order`, `is_enabled`) VALUES
(1, 1, '监督检查用语', 'INSPECTION', 1, 1),
(2, 1, '行政处罚用语', 'PENALTY', 2, 1),
(3, 1, '行政强制用语', 'ENFORCEMENT', 3, 1),
(4, 1, '整改指导用语', 'RECTIFY', 4, 1);

-- ============================
-- 初始化数据 - 规范条目（以监督检查用语为例）
-- ============================
INSERT INTO `law_standard_phrase_item` (`id`, `category_id`, `title`, `scene`, `sort_order`, `is_enabled`) VALUES
(1, 1, '出示执法证件', '检查开始前', 1, 1),
(2, 1, '说明检查目的', '检查开始前', 2, 1),
(3, 1, '告知权利义务', '检查过程中', 3, 1),
(4, 1, '现场检查记录', '检查过程中', 4, 1),
(5, 1, '检查结果反馈', '检查结束后', 5, 1),
(6, 1, '告知救济途径', '检查结束后', 6, 1);

-- ============================
-- 初始化数据 - 规范内容（以出示执法证件为例）
-- ============================
INSERT INTO `law_standard_phrase_content` (`id`, `item_id`, `content`, `legal_basis`, `tips`) VALUES
(1, 1, '您好！我们是 XX 卫生监督所的卫生监督员，这是我们的执法证件（出示证件）。现依法对你单位进行卫生监督检查，请予以配合。', '《中华人民共和国基本医疗卫生与健康促进法》第八十七条', '1. 必须由 2 名以上执法人员\n2. 主动出示有效执法证件\n3. 用语规范、态度文明'),
(2, 2, '根据《中华人民共和国传染病防治法》《公共场所卫生管理条例》等法律法规的规定，我们今天对你单位进行例行卫生监督检查。检查内容包括：卫生许可证持有情况、从业人员健康证明、卫生管理制度落实情况等。', '《公共场所卫生管理条例实施细则》第二十八条', '1. 说明检查法律依据\n2. 列出检查主要内容\n3. 语气平和、表达清晰'),
(3, 3, '根据法律规定，你单位有义务配合卫生监督检查，如实提供有关情况和资料。同时，你单位有权要求我们出示执法证件，对检查结果有异议可以申请行政复议或提起行政诉讼。', '《行政处罚法》第四十四条', '1. 告知配合义务\n2. 告知合法权利\n3. 表达公正执法态度'),
(4, 4, '根据现场检查情况，我们记录如下：1. 卫生许可证在有效期内；2. 从业人员健康证明齐全；3. 发现以下问题...（详细说明）。以上记录是否属实，请签字确认。', '《行政处罚法》第四十七条', '1. 如实记录检查情况\n2. 问题描述客观准确\n3. 请当事人签字确认'),
(5, 5, '本次检查发现你单位存在以下问题：1...2...3...。针对上述问题，我们提出以下整改意见：1...2...3...。请你单位于 XX 年 XX 月 XX 日前完成整改，并将整改情况报告我所。', '《公共场所卫生管理条例实施细则》第三十七条', '1. 问题表述清晰\n2. 整改意见具体\n3. 明确整改期限'),
(6, 6, '如你对本检查结果有异议，可以在收到本结果之日起 60 日内向 XX 人民政府或 XX 卫生健康委员会申请行政复议，也可以在 6 个月内向 XX 人民法院提起行政诉讼。', '《行政复议法》第九条', '1. 准确告知复议期限\n2. 准确告知诉讼期限\n3. 告知受理机关');
