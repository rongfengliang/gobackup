
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for userdemo
-- ----------------------------
DROP TABLE IF EXISTS `userdemo`;
CREATE TABLE `userdemo` (
  `username` varchar(255) DEFAULT NULL,
  `userage` int(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of userdemo
-- ----------------------------
BEGIN;
INSERT INTO `userdemo` VALUES ('dalong', 11);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
