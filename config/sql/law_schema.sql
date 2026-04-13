-- 移动卫生执法系统 - 数据库迁移脚本
-- 版本：v1.0.0
-- 日期：2026-04-13

-- 设置外键
SET FOREIGN_KEY_CHECKS = 0;

-- ============================
-- 1. 行业分类表
-- ============================
DROP TABLE IF EXISTS `law_industry`;
CREATE TABLE `law_industry` (
  `industry_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '行业 ID',
  `industry_code` varchar(50) DEFAULT NULL COMMENT '行业代码',
  `industry_name` varchar(100) DEFAULT NULL COMMENT '行业名称',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父级 ID',
  `level` tinyint(4) DEFAULT '1' COMMENT '层级（1:一级/2:二级）',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用（0:禁用/1:启用）',
  `order_num` int(11) DEFAULT '0' COMMENT '排序',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`industry_id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_level` (`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='行业分类表';

-- ============================
-- 2. 执法人员表
-- ============================
DROP TABLE IF EXISTS `law_official`;
CREATE TABLE `law_official` (
  `official_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '执法人员 ID',
  `user_id` bigint(20) DEFAULT NULL COMMENT '关联用户 ID',
  `badge_no` varchar(20) DEFAULT NULL COMMENT '执法证号',
  `department` varchar(100) DEFAULT NULL COMMENT '所属部门',
  `position` varchar(50) DEFAULT NULL COMMENT '职位',
  `law_type` varchar(50) DEFAULT NULL COMMENT '执法类型',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态（0:禁用/1:启用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`official_id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  KEY `idx_badge_no` (`badge_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='执法人员表';

-- ============================
-- 3. 设备表
-- ============================
DROP TABLE IF EXISTS `law_device`;
CREATE TABLE `law_device` (
  `device_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '设备 ID',
  `official_id` bigint(20) DEFAULT NULL COMMENT '关联执法人员 ID',
  `device_name` varchar(100) DEFAULT NULL COMMENT '设备名称',
  `device_model` varchar(50) DEFAULT NULL COMMENT '设备型号',
  `os_type` varchar(20) DEFAULT NULL COMMENT '操作系统类型（iOS/Android）',
  `os_version` varchar(20) DEFAULT NULL COMMENT '系统版本',
  `app_version` varchar(20) DEFAULT NULL COMMENT 'App 版本',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态（0:禁用/1:启用）',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(50) DEFAULT NULL COMMENT '最后登录 IP',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`device_id`),
  KEY `idx_official_id` (`official_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备表';

-- ============================
-- 4. 激活码表
-- ============================
DROP TABLE IF EXISTS `law_activate_code`;
CREATE TABLE `law_activate_code` (
  `code_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '激活码 ID',
  `activate_code` varchar(20) NOT NULL COMMENT '激活码',
  `official_id` bigint(20) DEFAULT NULL COMMENT '绑定执法人员 ID',
  `batch_no` varchar(50) DEFAULT NULL COMMENT '批次号',
  `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态（0:未使用/1:已激活/2:已过期/3:已禁用）',
  `activate_time` datetime DEFAULT NULL COMMENT '激活时间',
  `activate_device_id` bigint(20) DEFAULT NULL COMMENT '激活设备 ID',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`code_id`),
  UNIQUE KEY `idx_activate_code` (`activate_code`),
  KEY `idx_official_id` (`official_id`),
  KEY `idx_batch_no` (`batch_no`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='激活码表';

-- ============================
-- 5. 监管单位表
-- ============================
DROP TABLE IF EXISTS `law_subject`;
CREATE TABLE `law_subject` (
  `subject_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '单位 ID',
  `name` varchar(100) NOT NULL COMMENT '单位名称',
  `industry_id` bigint(20) DEFAULT NULL COMMENT '行业分类 ID',
  `industry_name` varchar(100) DEFAULT NULL COMMENT '行业名称（冗余）',
  `address` varchar(255) DEFAULT NULL COMMENT '经营地址',
  `contact_person` varchar(50) DEFAULT NULL COMMENT '联系人',
  `contact_phone` varchar(20) DEFAULT NULL COMMENT '联系电话',
  `license_no` varchar(50) DEFAULT NULL COMMENT '许可证号',
  `license_date` date DEFAULT NULL COMMENT '许可证日期',
  `license_expiry` date DEFAULT NULL COMMENT '许可证有效期',
  `business_scope` varchar(500) DEFAULT NULL COMMENT '经营范围',
  `lat` decimal(10,8) DEFAULT NULL COMMENT '纬度',
  `lng` decimal(11,8) DEFAULT NULL COMMENT '经度',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态（0:停用/1:正常）',
  `risk_level` tinyint(1) DEFAULT '1' COMMENT '风险等级（1:低/2:中/3:高）',
  `last_check_date` date DEFAULT NULL COMMENT '最后检查日期',
  `next_check_date` date DEFAULT NULL COMMENT '下次检查日期',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `sync_status` tinyint(1) DEFAULT '1' COMMENT '同步状态（0:待同步/1:已同步）',
  `sync_time` datetime DEFAULT NULL COMMENT '同步时间',
  PRIMARY KEY (`subject_id`),
  KEY `idx_industry_id` (`industry_id`),
  KEY `idx_name` (`name`),
  KEY `idx_status` (`status`),
  KEY `idx_sync_status` (`sync_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='监管单位表';

-- ============================
-- 6. 执法记录表
-- ============================
DROP TABLE IF EXISTS `law_enforcement_record`;
CREATE TABLE `law_enforcement_record` (
  `record_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录 ID',
  `record_no` varchar(50) DEFAULT NULL COMMENT '记录编号',
  `subject_id` bigint(20) DEFAULT NULL COMMENT '单位 ID',
  `subject_name` varchar(100) DEFAULT NULL COMMENT '单位名称（冗余）',
  `industry_id` bigint(20) DEFAULT NULL COMMENT '行业分类 ID',
  `check_date` datetime DEFAULT NULL COMMENT '检查日期',
  `check_type` varchar(20) DEFAULT NULL COMMENT '检查类型（日常/专项/复查/投诉）',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态（0:草稿/1:待上报/2:已上报/3:已审核/4:已归档）',
  `official_ids` json DEFAULT NULL COMMENT '参与执法人员 ID 列表',
  `official_names` json DEFAULT NULL COMMENT '参与执法人员姓名列表',
  `check_result` text COMMENT '检查结果',
  `problem_desc` text COMMENT '问题描述',
  `rectify_opinion` text COMMENT '整改意见',
  `rectify_deadline` date DEFAULT NULL COMMENT '整改期限',
  `latitude` decimal(10,8) DEFAULT NULL COMMENT '检查地点纬度',
  `longitude` decimal(11,8) DEFAULT NULL COMMENT '检查地点经度',
  `evidence_count` int(11) DEFAULT '0' COMMENT '证据数量',
  `document_count` int(11) DEFAULT '0' COMMENT '文书数量',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `sync_status` tinyint(1) DEFAULT '0' COMMENT '同步状态（0:待同步/1:已同步）',
  `sync_time` datetime DEFAULT NULL COMMENT '同步时间',
  PRIMARY KEY (`record_id`),
  KEY `idx_subject_id` (`subject_id`),
  KEY `idx_industry_id` (`industry_id`),
  KEY `idx_status` (`status`),
  KEY `idx_check_date` (`check_date`),
  KEY `idx_sync_status` (`sync_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='执法记录表';

-- ============================
-- 7. 证据材料表
-- ============================
DROP TABLE IF EXISTS `law_evidence`;
CREATE TABLE `law_evidence` (
  `evidence_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '证据 ID',
  `record_id` bigint(20) DEFAULT NULL COMMENT '执法记录 ID',
  `evidence_no` varchar(50) DEFAULT NULL COMMENT '证据编号',
  `type` varchar(20) DEFAULT NULL COMMENT '证据类型（photo/audio/video/document）',
  `title` varchar(100) DEFAULT NULL COMMENT '证据标题',
  `description` varchar(500) DEFAULT NULL COMMENT '证据描述',
  `file_path` varchar(255) DEFAULT NULL COMMENT '文件路径',
  `file_name` varchar(100) DEFAULT NULL COMMENT '文件名',
  `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小（字节）',
  `file_type` varchar(50) DEFAULT NULL COMMENT '文件类型（mime type）',
  `duration` int(11) DEFAULT '0' COMMENT '时长（秒，音频/视频）',
  `thumbnail_path` varchar(255) DEFAULT NULL COMMENT '缩略图路径',
  `latitude` decimal(10,8) DEFAULT NULL COMMENT '拍摄地点纬度',
  `longitude` decimal(11,8) DEFAULT NULL COMMENT '拍摄地点经度',
  `capture_time` datetime DEFAULT NULL COMMENT '采集时间',
  `upload_by` varchar(64) DEFAULT NULL COMMENT '上传人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `sync_status` tinyint(1) DEFAULT '0' COMMENT '同步状态（0:待同步/1:已同步）',
  PRIMARY KEY (`evidence_id`),
  KEY `idx_record_id` (`record_id`),
  KEY `idx_type` (`type`),
  KEY `idx_sync_status` (`sync_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='证据材料表';

-- ============================
-- 8. 文书模板表
-- ============================
DROP TABLE IF EXISTS `law_document_template`;
CREATE TABLE `law_document_template` (
  `template_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '模板 ID',
  `template_name` varchar(100) NOT NULL COMMENT '模板名称',
  `category_id` bigint(20) DEFAULT NULL COMMENT '分类 ID',
  `category_name` varchar(50) DEFAULT NULL COMMENT '分类名称（冗余）',
  `industry_id` bigint(20) DEFAULT NULL COMMENT '行业分类 ID',
  `industry_name` varchar(100) DEFAULT NULL COMMENT '行业名称（冗余）',
  `template_type` varchar(20) DEFAULT NULL COMMENT '模板类型（word/pdf）',
  `fields` json DEFAULT NULL COMMENT '填空项定义',
  `file_path` varchar(255) DEFAULT NULL COMMENT '模板文件路径',
  `file_content` longtext COMMENT '模板内容（base64）',
  `version` varchar(20) DEFAULT '1.0' COMMENT '版本号',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`template_id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_industry_id` (`industry_id`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文书模板表';

-- ============================
-- 9. 文书模板分类表
-- ============================
DROP TABLE IF EXISTS `law_template_category`;
CREATE TABLE `law_template_category` (
  `category_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类 ID',
  `category_name` varchar(50) NOT NULL COMMENT '分类名称',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父级 ID',
  `level` tinyint(4) DEFAULT '1' COMMENT '层级',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`category_id`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文书模板分类表';

-- ============================
-- 10. 生成文书表
-- ============================
DROP TABLE IF EXISTS `law_generated_document`;
CREATE TABLE `law_generated_document` (
  `doc_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文书 ID',
  `record_id` bigint(20) DEFAULT NULL COMMENT '执法记录 ID',
  `template_id` bigint(20) DEFAULT NULL COMMENT '模板 ID',
  `template_name` varchar(100) DEFAULT NULL COMMENT '模板名称（冗余）',
  `doc_title` varchar(200) DEFAULT NULL COMMENT '文书标题',
  `fill_data` json DEFAULT NULL COMMENT '填充数据',
  `file_path` varchar(255) DEFAULT NULL COMMENT '生成文件路径',
  `file_name` varchar(100) DEFAULT NULL COMMENT '文件名',
  `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小（字节）',
  `doc_type` varchar(20) DEFAULT NULL COMMENT '文书类型（word/pdf）',
  `print_count` int(11) DEFAULT '0' COMMENT '打印次数',
  `print_time` datetime DEFAULT NULL COMMENT '打印时间',
  `signatures` json DEFAULT NULL COMMENT '电子签名数据',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `sync_status` tinyint(1) DEFAULT '0' COMMENT '同步状态（0:待同步/1:已同步）',
  PRIMARY KEY (`doc_id`),
  KEY `idx_record_id` (`record_id`),
  KEY `idx_template_id` (`template_id`),
  KEY `idx_sync_status` (`sync_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='生成文书表';

-- ============================
-- 11. 同步队列表
-- ============================
DROP TABLE IF EXISTS `law_sync_queue`;
CREATE TABLE `law_sync_queue` (
  `queue_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '队列 ID',
  `table_name` varchar(50) NOT NULL COMMENT '表名',
  `record_id` bigint(20) NOT NULL COMMENT '记录 ID',
  `action` varchar(20) NOT NULL COMMENT '操作类型（insert/update/delete）',
  `sync_type` varchar(20) DEFAULT 'app_to_server' COMMENT '同步类型（app_to_server/server_to_app）',
  `data` json DEFAULT NULL COMMENT '变更数据',
  `priority` tinyint(4) DEFAULT '0' COMMENT '优先级（0:普通/1:重要/2:紧急）',
  `status` varchar(20) DEFAULT 'pending' COMMENT '状态（pending/success/failed/conflict）',
  `retry_count` int(11) DEFAULT '0' COMMENT '重试次数',
  `error_msg` varchar(500) DEFAULT NULL COMMENT '错误信息',
  `conflict_info` json DEFAULT NULL COMMENT '冲突信息',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `sync_time` datetime DEFAULT NULL COMMENT '同步时间',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  PRIMARY KEY (`queue_id`),
  KEY `idx_table_name` (`table_name`),
  KEY `idx_status` (`status`),
  KEY `idx_sync_type` (`sync_type`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='同步队列表';

-- ============================
-- 12. 同步日志表
-- ============================
DROP TABLE IF EXISTS `law_sync_log`;
CREATE TABLE `law_sync_log` (
  `log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志 ID',
  `device_id` bigint(20) DEFAULT NULL COMMENT '设备 ID',
  `official_id` bigint(20) DEFAULT NULL COMMENT '执法人员 ID',
  `sync_type` varchar(20) DEFAULT NULL COMMENT '同步类型（full/incremental）',
  `sync_tables` json DEFAULT NULL COMMENT '同步表列表',
  `record_count` int(11) DEFAULT '0' COMMENT '同步记录数',
  `success_count` int(11) DEFAULT '0' COMMENT '成功数',
  `failed_count` int(11) DEFAULT '0' COMMENT '失败数',
  `conflict_count` int(11) DEFAULT '0' COMMENT '冲突数',
  `duration` int(11) DEFAULT '0' COMMENT '耗时（秒）',
  `status` varchar(20) DEFAULT 'success' COMMENT '状态（success/failed/partial）',
  `error_msg` varchar(500) DEFAULT NULL COMMENT '错误信息',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`log_id`),
  KEY `idx_device_id` (`device_id`),
  KEY `idx_sync_type` (`sync_type`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='同步日志表';

-- ============================
-- 13. 法律法规表
-- ============================
DROP TABLE IF EXISTS `law_regulation`;
CREATE TABLE `law_regulation` (
  `regulation_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '法规 ID',
  `title` varchar(200) NOT NULL COMMENT '法规标题',
  `legal_type` varchar(50) DEFAULT NULL COMMENT '法律类型（法律/行政法规/部门规章/地方法规）',
  `supervision_types` json DEFAULT NULL COMMENT '监管类型列表',
  `industry_ids` json DEFAULT NULL COMMENT '关联行业 ID 列表',
  `publish_org` varchar(100) DEFAULT NULL COMMENT '发布机关',
  `publish_date` date DEFAULT NULL COMMENT '发布日期',
  `effective_date` date DEFAULT NULL COMMENT '生效日期',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态（0:废止/1:有效）',
  `content` longtext COMMENT '法规内容',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`regulation_id`),
  KEY `idx_legal_type` (`legal_type`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='法律法规表';

-- ============================
-- 14. 规范用语表
-- ============================
DROP TABLE IF EXISTS `law_standard_language`;
CREATE TABLE `law_standard_language` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(100) NOT NULL COMMENT '用语标题',
  `supervision_type` varchar(50) DEFAULT NULL COMMENT '监管类型',
  `category` varchar(50) DEFAULT NULL COMMENT '分类（检查前/检查中/检查后）',
  `scene` varchar(100) DEFAULT NULL COMMENT '适用场景',
  `content` text COMMENT '规范用语内容',
  `legal_basis` text COMMENT '法律依据',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_supervision_type` (`supervision_type`),
  KEY `idx_category` (`category`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='规范用语表';

-- ============================
-- 15. 监管事项表
-- ============================
DROP TABLE IF EXISTS `law_supervision_item`;
CREATE TABLE `law_supervision_item` (
  `item_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '事项 ID',
  `item_name` varchar(100) NOT NULL COMMENT '事项名称',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父级 ID',
  `level` tinyint(4) DEFAULT '1' COMMENT '层级',
  `supervision_type` varchar(50) DEFAULT NULL COMMENT '监管类型',
  `industry_ids` json DEFAULT NULL COMMENT '关联行业 ID 列表',
  `standard_language_ids` json DEFAULT NULL COMMENT '关联规范用语 ID 列表',
  `check_points` text COMMENT '检查要点',
  `legal_basis` text COMMENT '法律依据',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`item_id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_supervision_type` (`supervision_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='监管事项表';

-- ============================
-- 16. 通知公告表（增强）
-- ============================
DROP TABLE IF EXISTS `law_notice`;
CREATE TABLE `law_notice` (
  `notice_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '公告 ID',
  `title` varchar(200) NOT NULL COMMENT '公告标题',
  `content` text COMMENT '公告内容',
  `type` varchar(20) DEFAULT 'notice' COMMENT '类型（notice/announcement）',
  `priority` tinyint(4) DEFAULT '0' COMMENT '优先级（0:普通/1:重要/2:紧急）',
  `publish_org` varchar(100) DEFAULT NULL COMMENT '发布机关',
  `publish_time` datetime DEFAULT NULL COMMENT '发布时间',
  `target_audience` json DEFAULT NULL COMMENT '目标受众（null:全体/指定行业或人员）',
  `view_count` int(11) DEFAULT '0' COMMENT '浏览次数',
  `is_top` tinyint(1) DEFAULT '0' COMMENT '是否置顶',
  `is_enabled` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`),
  KEY `idx_type` (`type`),
  KEY `idx_priority` (`priority`),
  KEY `idx_is_enabled` (`is_enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通知公告表';

-- 恢复外键检查
SET FOREIGN_KEY_CHECKS = 1;

-- ============================
-- 初始化数据
-- ============================

-- 行业分类初始数据
INSERT INTO `law_industry` (`industry_id`, `industry_code`, `industry_name`, `parent_id`, `level`, `order_num`, `remark`) VALUES
(1, 'PUBLIC', '公共场所卫生', 0, 1, 1, '公共场所卫生监管'),
(2, 'PUBLIC_LODGE', '住宿场所', 1, 2, 1, '宾馆、酒店、旅馆等'),
(3, 'PUBLIC_BEAUTY', '美容美发场所', 1, 2, 2, '美容院、理发店等'),
(4, 'PUBLIC_BATH', '洗浴场所', 1, 2, 3, '浴场、桑拿中心等'),
(5, 'PUBLIC_POOL', '游泳场所', 1, 2, 4, '游泳馆、泳池等'),
(6, 'WATER', '生活饮用水卫生', 0, 1, 2, '生活饮用水监管'),
(7, 'WATER_SUPPLY', '供水单位', 6, 2, 1, '自来水厂、二次供水单位'),
(8, 'WATER_PURIFY', '涉及饮用水卫生安全产品', 6, 2, 2, '净水器、输配水设备等'),
(9, 'MEDICAL', '医疗机构卫生', 0, 1, 3, '医疗机构监管'),
(10, 'MEDICAL_HOSPITAL', '医院', 9, 2, 1, '综合医院、专科医院'),
(11, 'MEDICAL_CLINIC', '诊所', 9, 2, 2, '诊所、卫生室'),
(12, 'MEDICAL_DENTAL', '口腔诊所', 9, 2, 3, '口腔医院、口腔诊所'),
(13, 'SCHOOL', '学校卫生', 0, 1, 4, '学校卫生监管'),
(14, 'SCHOOL_PRIMARY', '小学', 13, 2, 1, '普通小学'),
(15, 'SCHOOL_MIDDLE', '中学', 13, 2, 2, '普通中学'),
(16, 'SCHOOL_UNI', '高等学校', 13, 2, 3, '大学、学院'),
(17, 'OCCUPATION', '职业卫生', 0, 1, 5, '职业卫生监管'),
(18, 'OCCUPATION_MINE', '煤矿卫生', 17, 2, 1, '煤矿作业场所'),
(19, 'OCCUPATION_CHEM', '化工卫生', 17, 2, 2, '化工企业');

-- 文书模板分类初始数据
INSERT INTO `law_template_category` (`category_id`, `category_name`, `parent_id`, `level`, `sort_order`, `remark`) VALUES
(1, '常用文书', 0, 1, 1, '日常执法常用文书'),
(2, '现场检查类', 1, 2, 1, '现场检查相关文书'),
(3, '调查取证类', 1, 2, 2, '调查取证相关文书'),
(4, '控制措施类', 1, 2, 3, '行政控制相关文书'),
(5, '处罚决定类', 1, 2, 4, '行政处罚相关文书'),
(6, '送达执行类', 1, 2, 5, '送达执行相关文书');

-- 监管类型初始数据
INSERT INTO `law_standard_language` (`id`, `title`, `supervision_type`, `category`, `content`) VALUES
(1, '出示执法证件', '检查前', '检查开始前必须出示执法证件', '您好！我们是 XX 卫生监督所卫生监督员，这是我们的执法证件。现依法对你单位进行卫生监督检查，请予以配合。'),
(2, '说明检查目的', '检查前', '检查开始前说明检查目的和依据', '根据《中华人民共和国基本医疗卫生与健康促进法》等法律法规的规定，我们对你单位进行例行卫生监督检查。'),
(3, '告知权利义务', '检查中', '检查过程中告知被检查人权利义务', '根据法律规定，你单位有义务配合卫生监督检查，如实提供有关情况和资料。同时，你单位有权要求我们出示执法证件，对检查结果有异议可以申请行政复议或提起行政诉讼。'),
(4, '检查结果反馈', '检查后', '检查结束后反馈检查结果', '本次检查发现你单位存在以下问题：1...2...3...。针对上述问题，我们提出以下整改意见：1...2...3...。请你单位于 XX 年 XX 月 XX 日前完成整改，并将整改情况报告我所。'),
(5, '告知救济途径', '检查后', '检查结束后告知救济途径', '如你对本检查结果有异议，可以在收到本结果之日起 60 日内向 XX 人民政府或 XX 卫生健康委员会申请行政复议，也可以在 6 个月内向 XX 人民法院提起行政诉讼。');
