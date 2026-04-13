-- 移动卫生执法系统 - 法律法规库扩展表结构
-- 版本：v1.0.0
-- 日期：2026-04-13
-- 说明：补充法律法规的章节表和条款表

SET FOREIGN_KEY_CHECKS = 0;

-- ============================
-- 1. 法律法规章节表
-- ============================
DROP TABLE IF EXISTS `law_regulation_chapter`;
CREATE TABLE `law_regulation_chapter` (
  `chapter_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '章节 ID',
  `regulation_id` bigint(20) NOT NULL COMMENT '法规 ID',
  `chapter_no` int(11) DEFAULT '0' COMMENT '章节序号',
  `chapter_title` varchar(200) DEFAULT NULL COMMENT '章节标题',
  `chapter_type` varchar(20) DEFAULT 'chapter' COMMENT '章节类型（chapter/section/part）',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父级章节 ID（用于嵌套章节）',
  `level` tinyint(4) DEFAULT '1' COMMENT '层级',
  `content` text COMMENT '章节前言/概述内容',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`chapter_id`),
  KEY `idx_regulation_id` (`regulation_id`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='法律法规章节表';

-- ============================
-- 2. 法律法规条款表
-- ============================
DROP TABLE IF EXISTS `law_regulation_article`;
CREATE TABLE `law_regulation_article` (
  `article_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '条款 ID',
  `regulation_id` bigint(20) NOT NULL COMMENT '法规 ID',
  `chapter_id` bigint(20) DEFAULT NULL COMMENT '所属章节 ID',
  `article_no` varchar(50) DEFAULT NULL COMMENT '条款编号（如：第一条、第 1 条）',
  `article_no_sort` int(11) DEFAULT '0' COMMENT '条款序号（用于排序）',
  `title` varchar(200) DEFAULT NULL COMMENT '条款标题',
  `content` text COMMENT '条款内容',
  `penalty_basis` text COMMENT '处罚依据',
  `penalty_type` varchar(100) DEFAULT NULL COMMENT '处罚种类',
  `discretion_level` varchar(50) DEFAULT NULL COMMENT '裁量阶次',
  `applicable_scenario` text COMMENT '适用情形',
  `penalty_range` varchar(200) DEFAULT NULL COMMENT '裁量幅度',
  `remark_text` varchar(500) DEFAULT NULL COMMENT '备注',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`article_id`),
  KEY `idx_regulation_id` (`regulation_id`),
  KEY `idx_chapter_id` (`chapter_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='法律法规条款表';

-- ============================
-- 3. 定性依据表（原 PRD 中的 basis）
-- ============================
DROP TABLE IF EXISTS `law_qualification_basis`;
CREATE TABLE `law_qualification_basis` (
  `basis_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '依据 ID',
  `regulation_id` bigint(20) DEFAULT NULL COMMENT '关联法规 ID',
  `article_id` bigint(20) DEFAULT NULL COMMENT '关联条款 ID',
  `title` varchar(200) NOT NULL COMMENT '依据标题',
  `content` text COMMENT '依据内容',
  `basis_type` varchar(50) DEFAULT NULL COMMENT '依据类型（定性/处罚/裁量）',
  `legal_basis` text COMMENT '法律依据',
  `penalty_basis` text COMMENT '处罚依据',
  `penalty_type` varchar(100) DEFAULT NULL COMMENT '处罚种类',
  `discretion_level` varchar(50) DEFAULT NULL COMMENT '裁量阶次',
  `applicable_scenario` text COMMENT '适用情形',
  `penalty_range` varchar(200) DEFAULT NULL COMMENT '裁量幅度',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`basis_id`),
  KEY `idx_regulation_id` (`regulation_id`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_basis_type` (`basis_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='定性依据表';

-- ============================
-- 4. 法律类型字典表
-- ============================
DROP TABLE IF EXISTS `law_legal_type_dict`;
CREATE TABLE `law_legal_type_dict` (
  `type_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '类型 ID',
  `type_name` varchar(50) NOT NULL COMMENT '类型名称（法律/法规/规章/规范性文件/批复文件/标准）',
  `type_code` varchar(50) NOT NULL COMMENT '类型代码',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`type_id`),
  UNIQUE KEY `idx_type_code` (`type_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='法律类型字典表';

-- ============================
-- 5. 监管类型字典表
-- ============================
DROP TABLE IF EXISTS `law_supervision_type_dict`;
CREATE TABLE `law_supervision_type_dict` (
  `type_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '类型 ID',
  `type_name` varchar(100) NOT NULL COMMENT '类型名称',
  `type_code` varchar(50) NOT NULL COMMENT '类型代码',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父级 ID',
  `level` tinyint(4) DEFAULT '1' COMMENT '层级',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`type_id`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='监管类型字典表';

-- ============================
-- 初始化数据
-- ============================

-- 法律类型字典数据
INSERT INTO `law_legal_type_dict` (`type_id`, `type_name`, `type_code`, `sort_order`) VALUES
(1, '法律', 'law', 1),
(2, '行政法规', 'regulation', 2),
(3, '部门规章', 'rule', 3),
(4, '地方法规', 'local_regulation', 4),
(5, '规范性文件', 'normative_document', 5),
(6, '批复文件', 'approval_document', 6),
(7, '标准', 'standard', 7);

-- 监管类型字典数据（12 个监管类型）
INSERT INTO `law_supervision_type_dict` (`type_id`, `type_name`, `type_code`, `parent_id`, `level`, `sort_order`) VALUES
(1, '食品安全', 'food_safety', 0, 1, 1),
(2, '公共场所', 'public_place', 0, 1, 2),
(3, '消毒产品', 'disinfection_product', 0, 1, 3),
(4, '生活饮用水及涉水产品', 'water_product', 0, 1, 4),
(5, '放射卫生', 'radiation_health', 0, 1, 5),
(6, '职业卫生', 'occupation_health', 0, 1, 6),
(7, '医疗机构', 'medical_institution', 0, 1, 7),
(8, '二次供水', 'secondary_water', 0, 1, 8),
(9, '学校卫生', 'school_health', 0, 1, 9),
(10, '放射诊疗卫生', 'radiation_diagnosis', 0, 1, 10),
(11, '小型医疗机构', 'small_medical', 0, 1, 11),
(12, '门诊/诊所传染病', 'clinic_infection', 0, 1, 12);

-- 恢复外键检查
SET FOREIGN_KEY_CHECKS = 1;
