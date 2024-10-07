/*
 Navicat Premium Data Transfer

 Source Server         : my-server
 Source Server Type    : MySQL
 Source Server Version : 80037 (8.0.37)
 Source Host           : localhost:3306
 Source Schema         : rbac_system

 Target Server Type    : MySQL
 Target Server Version : 80037 (8.0.37)
 File Encoding         : 65001

 Date: 07/10/2024 16:40:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户姓名',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户邮箱',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户密码',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户电话',
  `department_id` int NULL DEFAULT NULL COMMENT '部门ID',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '用户状态（1=启用，0=禁用）',
  `last_login` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `email`(`email` ASC) USING BTREE,
  INDEX `department_id`(`department_id` ASC) USING BTREE,
  CONSTRAINT `admin_users_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`) ON DELETE SET NULL ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '后台用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_users
-- ----------------------------
INSERT INTO `admin_users` VALUES (1, 'admin', 'lin123456@gmail.com', '123456', '13800138000', 1, 1, '2024-10-07 16:13:42', '2024-09-13 15:36:18', '2024-10-07 16:21:24');
INSERT INTO `admin_users` VALUES (3, '王芳', 'wangfang@example.com', 'password3', '13800138002', 9, 1, '2024-01-12 11:00:00', '2024-09-13 15:36:18', '2024-09-28 12:19:20');
INSERT INTO `admin_users` VALUES (4, '赵敏', 'zhaomin@example.com', 'password4', '13800138003', 5, 1, '2024-01-13 15:20:00', '2024-09-13 15:36:18', '2024-09-21 14:34:51');
INSERT INTO `admin_users` VALUES (5, '钱强', 'qianqiang@example.com', 'password5', '13800138004', 5, 1, '2024-01-14 10:15:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (6, '孙丽', 'sunli@example.com', '123456', '13800138005', 6, 1, '2024-10-07 00:30:31', '2024-09-13 15:36:18', '2024-10-07 00:30:31');
INSERT INTO `admin_users` VALUES (7, '周涛', 'zhoutao@example.com', 'password7', '13800138006', 7, 1, '2024-01-16 13:00:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (8, '吴敏', 'wumin@example.com', '123456', '13800138007', 8, 1, '2024-10-07 00:26:36', '2024-09-13 15:36:18', '2024-10-07 00:26:36');
INSERT INTO `admin_users` VALUES (9, '郑杰', 'zhengjie@example.com', 'password9', '13800138008', 9, 1, '2024-01-18 14:00:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (10, '冯楠11', 'fengnan@example.com', 'password10', '13800138009', 1, 1, '2024-01-19 11:10:00', '2024-09-13 15:36:18', '2024-09-28 14:44:13');
INSERT INTO `admin_users` VALUES (11, '陈英', 'chenying@example.com', 'password11', '13800138099', 2, 1, '2024-01-20 15:45:00', '2024-09-13 15:36:18', '2024-09-28 14:44:00');
INSERT INTO `admin_users` VALUES (12, '杨阳', 'yangyang@example.com', 'password12', '13800138011', 10, 1, '2024-01-21 16:30:00', '2024-09-13 15:36:18', '2024-09-21 14:33:38');
INSERT INTO `admin_users` VALUES (13, '何勇', 'heyong@example.com', 'password13', '13800138012', 5, 1, '2024-01-22 10:50:00', '2024-09-13 15:36:18', '2024-09-21 14:34:42');
INSERT INTO `admin_users` VALUES (14, '吕佳', 'lvjia@example.com', 'password14', '13800138013', 5, 1, '2024-01-23 09:00:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (15, '施涛', 'shitao@example.com', 'password15', '13800138014', 6, 1, '2024-01-24 08:30:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (16, '范冰', 'fanbing@example.com', 'password16', '13800138015', 7, 1, '2024-01-25 14:30:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (17, '彭明', 'pengming@example.com', 'password17', '13800138016', 8, 1, '2024-01-26 11:00:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (18, '高峰', 'gaofeng@example.com', 'password18', '13800138017', 9, 1, '2024-01-27 17:00:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (19, '马红', 'mahong@example.com', 'password19', '13800138018', 1, 1, '2024-01-28 15:30:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (20, '刘艳', 'liuyan@example.com', 'password20', '13800138019', 2, 1, '2024-01-29 16:45:00', '2024-09-13 15:36:18', '2024-09-13 15:36:18');
INSERT INTO `admin_users` VALUES (21, 'lwf', 'linwe8@gmail.com', '123456', '19121237808', 10, 1, NULL, '2024-09-18 23:32:51', '2024-10-07 16:21:53');
INSERT INTO `admin_users` VALUES (24, '林总', '17326@qq.com', '123456', '13984888616', 10, 1, NULL, '2024-09-21 14:15:43', '2024-10-07 16:21:47');

-- ----------------------------
-- Table structure for departments
-- ----------------------------
DROP TABLE IF EXISTS `departments`;
CREATE TABLE `departments`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `parent_id` int NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `parent_id`(`parent_id` ASC) USING BTREE,
  CONSTRAINT `departments_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `departments` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '部门表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of departments
-- ----------------------------
INSERT INTO `departments` VALUES (1, '总部', NULL, 'ApartmentOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (2, '人事部', 1, 'TeamOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (3, '技术部', 1, 'TeamOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (4, 'unicube', NULL, 'ApartmentOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (5, 'web端', 4, 'TeamOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (6, 'admin端', 4, 'TeamOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (7, '招聘组', 2, 'UserOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (8, '培训组', 2, 'UserOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (9, '前端组', 3, 'UserOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');
INSERT INTO `departments` VALUES (10, '后端组', 3, 'UserOutlined', '2024-09-22 22:48:33', '2024-09-22 22:48:33');

-- ----------------------------
-- Table structure for dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `dictionaries`;
CREATE TABLE `dictionaries`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '字典类型',
  `label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '显示名称',
  `value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '值',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '描述',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `sort` int NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `type_value`(`type` ASC, `value` ASC) USING BTREE COMMENT '保证同一类型下值唯一',
  INDEX `idx_type`(`type` ASC) USING BTREE,
  INDEX `idx_sort_id`(`sort` ASC, `id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '数据字典表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dictionaries
-- ----------------------------
INSERT INTO `dictionaries` VALUES (1, 'ROLE_TYPE', '超级管理员', '1', '拥有系统所有权限', '2024-09-16 23:43:36', '2024-09-29 16:12:04', 1);
INSERT INTO `dictionaries` VALUES (2, 'ROLE_TYPE', '初审员', '2', '负责初步审核', '2024-09-16 23:43:36', '2024-09-29 16:26:36', 2);
INSERT INTO `dictionaries` VALUES (3, 'ROLE_TYPE', '复审员', '3', '负责复审', '2024-09-16 23:43:36', '2024-09-29 16:26:37', 3);
INSERT INTO `dictionaries` VALUES (4, 'ROLE_TYPE', '财务管理员', '4', '管理财务相关事务', '2024-09-16 23:43:36', '2024-09-29 16:12:06', 4);
INSERT INTO `dictionaries` VALUES (5, 'ROLE_TYPE', '人事管理员', '5', '管理人事相关事务', '2024-09-16 23:43:36', '2024-09-29 16:12:06', 5);
INSERT INTO `dictionaries` VALUES (6, 'ROLE_TYPE', '技术支持', '6', '提供技术支持', '2024-09-16 23:43:36', '2024-09-29 16:12:07', 6);
INSERT INTO `dictionaries` VALUES (7, 'ROLE_TYPE', '市场推广', '7', '负责市场推广', '2024-09-16 23:43:36', '2024-09-29 16:12:08', 7);
INSERT INTO `dictionaries` VALUES (8, 'ROLE_TYPE', '运营人员', '8', '负责运营事务', '2024-09-16 23:43:36', '2024-09-29 16:12:10', 8);
INSERT INTO `dictionaries` VALUES (9, 'ROLE_TYPE', '产品经理', '9', '负责产品规划与设计', '2024-09-16 23:43:36', '2024-09-29 16:12:12', 9);
INSERT INTO `dictionaries` VALUES (14, '3123', '12312', '31231', '', '2024-09-29 17:39:33', '2024-09-29 17:39:33', 3213);

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `component_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `parent_id` int UNSIGNED NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT 1 COMMENT '0-禁用 1-启用',
  `permissions` json NULL COMMENT '菜单默认权限',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (1, '系统管理', '/system', NULL, 'emojione:gear', NULL, 1, '2024-09-21 16:12:18.000', '2024-09-21 16:12:18.000', 1, NULL);
INSERT INTO `menus` VALUES (2, '用户管理', '/system/user', 'system/UserManagement', 'emojione:busts-in-silhouette', 1, 1, '2024-09-21 16:12:39.000', '2024-10-07 00:29:37.273', 1, '[\"system:user:update\", \"system:user:add\", \"system:user:delete\"]');
INSERT INTO `menus` VALUES (3, '角色管理', '/system/role', 'system/RoleManagement', 'emojione:shield', 1, 2, '2024-09-21 16:12:39.000', '2024-09-29 15:48:28.694', 1, NULL);
INSERT INTO `menus` VALUES (4, '菜单管理', '/system/menu', 'system/MenuManagement', 'emojione:clipboard', 1, 3, '2024-09-21 16:12:39.000', '2024-09-28 14:55:36.300', 1, NULL);
INSERT INTO `menus` VALUES (7, '字典管理', '/system/dict', 'system/DictManagement', 'emojione:books', 1, 4, '2024-09-29 16:15:50.397', '2024-09-29 16:16:37.425', 1, NULL);

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '权限代码',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '权限描述',
  `menu_id` int NULL DEFAULT NULL COMMENT '关联的菜单ID',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code`(`code` ASC) USING BTREE,
  INDEX `fk_permissions_menu`(`menu_id` ASC) USING BTREE,
  CONSTRAINT `fk_permissions_menu` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of permissions
-- ----------------------------
INSERT INTO `permissions` VALUES (2, 'system:user:update', '更新用户按钮', 2, '2024-10-06 15:41:24', '2024-10-07 00:29:37');
INSERT INTO `permissions` VALUES (10, 'system:user:add', '新增用户按钮', 2, '2024-10-06 23:07:49', '2024-10-07 00:29:37');
INSERT INTO `permissions` VALUES (11, 'system:user:delete', '删除用户按钮', 2, '2024-10-06 23:28:49', '2024-10-07 00:29:37');

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus`  (
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `fk_role_menus_menu`(`menu_id` ASC) USING BTREE,
  CONSTRAINT `fk_role_menus_menu` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_role_menus_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role_menus
-- ----------------------------
INSERT INTO `role_menus` VALUES (1, 1);
INSERT INTO `role_menus` VALUES (6, 1);
INSERT INTO `role_menus` VALUES (1, 2);
INSERT INTO `role_menus` VALUES (2, 2);
INSERT INTO `role_menus` VALUES (3, 2);
INSERT INTO `role_menus` VALUES (6, 2);
INSERT INTO `role_menus` VALUES (1, 3);
INSERT INTO `role_menus` VALUES (6, 3);
INSERT INTO `role_menus` VALUES (1, 4);
INSERT INTO `role_menus` VALUES (6, 4);
INSERT INTO `role_menus` VALUES (1, 7);
INSERT INTO `role_menus` VALUES (6, 7);

-- ----------------------------
-- Table structure for role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `role_permissions`;
CREATE TABLE `role_permissions`  (
  `role_id` int NOT NULL,
  `permission_id` int NOT NULL,
  PRIMARY KEY (`role_id`, `permission_id`) USING BTREE,
  INDEX `fk_role_permissions_permission`(`permission_id` ASC) USING BTREE,
  CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色权限关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role_permissions
-- ----------------------------
INSERT INTO `role_permissions` VALUES (2, 2);
INSERT INTO `role_permissions` VALUES (6, 2);
INSERT INTO `role_permissions` VALUES (2, 10);
INSERT INTO `role_permissions` VALUES (6, 10);
INSERT INTO `role_permissions` VALUES (3, 11);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '角色描述',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, '超级管理员', '拥有系统的所有权限1', '2024-09-13 15:36:37', '2024-09-28 14:51:09');
INSERT INTO `roles` VALUES (2, '初审员', '负责初步审核用户的申请', '2024-09-13 15:36:37', '2024-10-06 21:01:10');
INSERT INTO `roles` VALUES (3, '复审员', '负责复审通过初审的用户申请', '2024-09-13 15:36:37', '2024-10-06 21:01:10');
INSERT INTO `roles` VALUES (4, '财务管理员', '负责处理财务相关事务', '2024-09-13 15:36:37', '2024-09-13 15:36:37');
INSERT INTO `roles` VALUES (5, '人事管理员', '负责管理人事事务', '2024-09-13 15:36:37', '2024-09-29 16:45:47');
INSERT INTO `roles` VALUES (6, '技术支持', '提供技术帮助和支持', '2024-09-13 15:36:37', '2024-09-13 15:36:37');
INSERT INTO `roles` VALUES (7, '市场推广', '负责市场推广和营销', '2024-09-13 15:36:37', '2024-09-13 15:36:37');
INSERT INTO `roles` VALUES (8, '运营人员', '负责系统的日常运营', '2024-09-13 15:36:37', '2024-09-13 15:36:37');
INSERT INTO `roles` VALUES (9, '产品经理', '负责产品规划和管理', '2024-09-13 15:36:37', '2024-09-13 15:36:37');

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles`  (
  `user_id` int NOT NULL,
  `role_id` int NOT NULL,
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE,
  INDEX `role_id`(`role_id` ASC) USING BTREE,
  CONSTRAINT `user_roles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `admin_users` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `user_roles_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户角色关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO `user_roles` VALUES (1, 1);
INSERT INTO `user_roles` VALUES (9, 1);
INSERT INTO `user_roles` VALUES (18, 1);
INSERT INTO `user_roles` VALUES (21, 1);
INSERT INTO `user_roles` VALUES (10, 2);
INSERT INTO `user_roles` VALUES (19, 2);
INSERT INTO `user_roles` VALUES (3, 3);
INSERT INTO `user_roles` VALUES (11, 3);
INSERT INTO `user_roles` VALUES (20, 3);
INSERT INTO `user_roles` VALUES (3, 4);
INSERT INTO `user_roles` VALUES (4, 4);
INSERT INTO `user_roles` VALUES (12, 4);
INSERT INTO `user_roles` VALUES (5, 5);
INSERT INTO `user_roles` VALUES (13, 5);
INSERT INTO `user_roles` VALUES (6, 6);
INSERT INTO `user_roles` VALUES (14, 6);
INSERT INTO `user_roles` VALUES (24, 6);
INSERT INTO `user_roles` VALUES (7, 7);
INSERT INTO `user_roles` VALUES (15, 7);
INSERT INTO `user_roles` VALUES (8, 8);
INSERT INTO `user_roles` VALUES (16, 8);
INSERT INTO `user_roles` VALUES (17, 9);

SET FOREIGN_KEY_CHECKS = 1;
