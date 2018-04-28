/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50717
Source Host           : localhost:3306
Source Database       : movie_site

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2018-04-28 19:12:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` varchar(50) NOT NULL COMMENT '账户',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `role` int(11) NOT NULL COMMENT '1超级管理员2普通管理员',
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='管理员表';

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES ('1', 'hxy', '123456', '1', '1');
INSERT INTO `admin` VALUES ('2', 'zym', '123456', '2', '2');
INSERT INTO `admin` VALUES ('3', 'zyt', '123456', '2', '1');
INSERT INTO `admin` VALUES ('4', 'zyf', '123456', '2', '1');
INSERT INTO `admin` VALUES ('5', 'zgf', '123456', '2', '1');

-- ----------------------------
-- Table structure for advise
-- ----------------------------
DROP TABLE IF EXISTS `advise`;
CREATE TABLE `advise` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` varchar(50) NOT NULL COMMENT '账户',
  `role` enum('resource','build') DEFAULT 'resource' COMMENT '建议类型',
  `title` varchar(50) DEFAULT NULL COMMENT '标题',
  `content` text COMMENT '内容',
  `status` enum('wait','pass','fair') DEFAULT NULL COMMENT '建议状态',
  `advise_date` date DEFAULT NULL COMMENT '建议时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 COMMENT='看电影历史记录表';

-- ----------------------------
-- Records of advise
-- ----------------------------
INSERT INTO `advise` VALUES ('1', '958752538', 'resource', '《起跑线》', '最近特别火的一部印度电影起跑线。希望该网站能够观看。谢谢！', 'wait', '2018-04-23');
INSERT INTO `advise` VALUES ('2', '958752538', 'build', '没有搜索功能', '网站是挺好的，但是好像没有电影搜索功能呀！这就有点可惜啊，希望能快点增加该功能', 'wait', '2018-04-23');
INSERT INTO `advise` VALUES ('5', '958752538', 'resource', '头号玩家', '故事我给3星\n特效我给4星\n希望能够看到', 'pass', '2018-04-23');
INSERT INTO `advise` VALUES ('6', '958752538', 'build', '网站的推荐功能', '网站的推荐功能我是觉得不错的\n但是感觉算法还是有那么一些缺陷\n希望能够改进吧', 'pass', '2018-04-23');
INSERT INTO `advise` VALUES ('7', '141842291', 'build', '资讯问题', '我觉得资讯的登录好像用不了啊', 'pass', '2018-04-24');
INSERT INTO `advise` VALUES ('8', '958752538', 'resource', '《起跑线》', '最近特别火的一部印度电影起跑线。希望该网站能够观看。谢谢！', 'pass', '2018-04-23');
INSERT INTO `advise` VALUES ('9', '958752538', 'build', '没有搜索功能', '网站是挺好的，但是好像没有电影搜索功能呀！这就有点可惜啊，希望能快点增加该功能', 'pass', '2018-04-23');
INSERT INTO `advise` VALUES ('10', '958752538', 'resource', '头号玩家', '故事我给3星\n特效我给4星\n希望能够看到', 'pass', '2018-04-23');
INSERT INTO `advise` VALUES ('11', '958752538', 'build', '网站的推荐功能', '网站的推荐功能我是觉得不错的\n但是感觉算法还是有那么一些缺陷\n希望能够改进吧', 'wait', '2018-04-23');
INSERT INTO `advise` VALUES ('12', '141842291', 'build', '资讯问题', '我觉得资讯的登录好像用不了啊', 'pass', '2018-04-24');
INSERT INTO `advise` VALUES ('13', '958752538', 'resource', '《起跑线》', '最近特别火的一部印度电影起跑线。希望该网站能够观看。谢谢！', 'wait', '2018-04-23');
INSERT INTO `advise` VALUES ('14', '958752538', 'build', '没有搜索功能', '网站是挺好的，但是好像没有电影搜索功能呀！这就有点可惜啊，希望能快点增加该功能', 'wait', '2018-04-23');
INSERT INTO `advise` VALUES ('15', '958752538', 'resource', '头号玩家', '故事我给3星\n特效我给4星\n希望能够看到', 'wait', '2018-04-23');
INSERT INTO `advise` VALUES ('16', '958752538', 'build', '网站的推荐功能', '网站的推荐功能我是觉得不错的\n但是感觉算法还是有那么一些缺陷\n希望能够改进吧', 'wait', '2018-04-23');
INSERT INTO `advise` VALUES ('17', '141842291', 'build', '资讯问题', '我觉得资讯的登录好像用不了啊', 'pass', '2018-04-24');
INSERT INTO `advise` VALUES ('18', '141842291', 'resource', '哈哈哈', '休息西区', 'wait', '2018-04-27');

-- ----------------------------
-- Table structure for area
-- ----------------------------
DROP TABLE IF EXISTS `area`;
CREATE TABLE `area` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(11) NOT NULL COMMENT '地区名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='地区表';

-- ----------------------------
-- Records of area
-- ----------------------------
INSERT INTO `area` VALUES ('1', '大陆');
INSERT INTO `area` VALUES ('2', '香港');
INSERT INTO `area` VALUES ('3', '台湾');
INSERT INTO `area` VALUES ('4', '日本');
INSERT INTO `area` VALUES ('5', '韩国');
INSERT INTO `area` VALUES ('6', '英国');
INSERT INTO `area` VALUES ('7', '西班牙');
INSERT INTO `area` VALUES ('8', '法国');
INSERT INTO `area` VALUES ('9', '美国');
INSERT INTO `area` VALUES ('10', '德国');
INSERT INTO `area` VALUES ('11', '印度');
INSERT INTO `area` VALUES ('12', '其他');

-- ----------------------------
-- Table structure for film_review
-- ----------------------------
DROP TABLE IF EXISTS `film_review`;
CREATE TABLE `film_review` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `movie_id` int(11) NOT NULL COMMENT '电影id',
  `user_id` varchar(50) NOT NULL COMMENT '用户id',
  `title` varchar(100) NOT NULL COMMENT '影评标题',
  `content` text NOT NULL COMMENT '内容',
  `date` datetime NOT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='电影影评';

-- ----------------------------
-- Records of film_review
-- ----------------------------
INSERT INTO `film_review` VALUES ('1', '1', '141842291', '第一次的影评（求赞(〃\'▽\'〃)）', '<p>日本的电影总是有一种特殊的魔力</p><p>第一眼总是平凡无奇</p><p>但是到了最后却有让人那个忘怀</p><p>缓缓的2小时的电影，也缓缓的流入了我的内心</p><p><img src=\"https://img1.doubanio.com/view/photo/l/public/p2516296918.webp\">&nbsp;&nbsp;</p>', '2018-04-03 09:40:54');
INSERT INTO `film_review` VALUES ('2', '7', '141842291', '长泽雅美图', '<p><span style=\"color: rgb(249, 150, 59);\">相信很多小伙伴看过这部影片以后首先会被雅美桑的颜值给惊呆了！会被这个风景所感染</span></p><p><span style=\"color: rgb(139, 170, 74);\">都会被里面的女主角所吸引</span></p><p><span style=\"color: rgb(70, 172, 200);\">所以我特地在这个地方分享一下该女神的一些写真</span></p><p><img src=\"https://timgsa.baidu.com/timg?image&amp;quality=80&amp;size=b9999_10000&amp;sec=1522813142953&amp;di=a0a4af6f6da3532d76da3368e4883f33&amp;imgtype=0&amp;src=http%3A%2F%2Ff9.topitme.com%2F9%2Fbf%2Fb0%2F1149011417788b0bf9o.jpg\" style=\"max-width: 100%;\"><span style=\"color: rgb(77, 128, 191);\">&nbsp;！！！ ----------------------------------------</span></p><p><img src=\"https://timgsa.baidu.com/timg?image&amp;quality=80&amp;size=b9999_10000&amp;sec=1522813142953&amp;di=a607860a0e222e6ec9ef84a47325b798&amp;imgtype=0&amp;src=http%3A%2F%2Fimg3.duitang.com%2Fuploads%2Fitem%2F201205%2F28%2F20120528121844_CNemU.jpeg\" style=\"max-width:100%;\"><br></p>', '2018-04-04 09:41:28');
INSERT INTO `film_review` VALUES ('3', '1', '141842291', '第二次了（太好看了）', '<p>这是我弟二次写这部电影的评论了！</p><p>讲真的太好看了</p><p>我超级喜欢看</p><p>没有之一</p><p>这的</p><p>他就是用这种魔力</p><p>令我留恋</p><p>嗯，很nice</p><p>相当的不错哦</p><p>强烈推荐给大家！</p><p>强烈推荐给大家 ！</p><p>强烈推荐给大家！&nbsp;</p>', '2018-04-04 09:48:01');
INSERT INTO `film_review` VALUES ('4', '1', '141842261', '我的第一次', '<p>想不到第一次写影评</p><p>是因为测试这个下这个功能是不是有bug窝草</p><p>醉了</p>', '2018-04-04 09:50:50');
INSERT INTO `film_review` VALUES ('5', '2', '141842261', '期待', '<p>感觉还是很令人期待</p><p>pv我就很期待了</p><p>没想到超出期待值好多呢</p><p>好开心啊</p>', '2018-04-08 10:33:33');
INSERT INTO `film_review` VALUES ('6', '9', '141842261', '烧脑', '<p>借着推理片</p><p>描绘出一个深沉爱</p><p>每个人都是不同的</p><p>但都是殊途同归</p><p>为了自己爱的人啊</p>', '2018-04-08 10:36:38');
INSERT INTO `film_review` VALUES ('7', '12', '141842261', '恋人', '<p>所谓恋人</p><p>就是。。。也不能乱下定义吧</p><p>泪目</p>', '2018-04-08 10:38:30');
INSERT INTO `film_review` VALUES ('8', '20', '141842261', '战争啊！', '<p>战争</p><p>这个词语</p><p>光是听到就会令许多了两股战战</p><p>很庆幸自己能活在个以和平的时代</p><p>希望永恒</p>', '2018-04-08 10:41:11');
INSERT INTO `film_review` VALUES ('9', '22', '141842261', '永恒', '<p>这部影片对我来说</p><p>就是对爱情的最好的诠释</p><p><br></p>', '2018-04-08 10:43:53');
INSERT INTO `film_review` VALUES ('10', '1', '141842261', '我是露', '<p>这部电影</p><p>在一处让我看到了音乐的力量</p><p>是都多么的伟大啊<img src=\"https://timgsa.baidu.com/timg?image&amp;quality=80&amp;size=b9999_10000&amp;sec=1523173045813&amp;di=0e29046df1ddea9d19231c0a43565aa9&amp;imgtype=0&amp;src=http%3A%2F%2Fimgsrc.baidu.com%2Fimgad%2Fpic%2Fitem%2Fcf1b9d16fdfaaf5154a06bb7865494eef01f7a73.jpg\" style=\"max-width: 100%;\"></p>', '2018-04-08 12:49:46');
INSERT INTO `film_review` VALUES ('13', '1', '141842261', '听说写影评有奖励', '<p>那我就测试起来</p><p>积满四行</p><p>意思一下</p><p>加油</p>', '2018-04-08 17:00:13');
INSERT INTO `film_review` VALUES ('14', '1', '141842261', '王企鹅', '<p>大事的群无若</p><p>asdasd1223</p><p>12312</p><p>dasd</p><p>asdasd</p>', '2018-04-08 17:00:59');
INSERT INTO `film_review` VALUES ('15', '4', '958752538', '我是新人(〃\'▽\'〃)', '<p>今天看了声之形</p><p>真的特别感动呢</p><p>也想跟大家一起分享一下我的喜悦</p><p>嘻嘻嘻嘻</p>', '2018-04-20 14:42:50');
INSERT INTO `film_review` VALUES ('16', '42', '141842291', '学长大傻逼', '<p>哈哈哈哈</p>', '2018-04-27 16:35:56');

-- ----------------------------
-- Table structure for fr_attribute
-- ----------------------------
DROP TABLE IF EXISTS `fr_attribute`;
CREATE TABLE `fr_attribute` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `fr_id` int(11) NOT NULL COMMENT '影评id',
  `movie_id` int(11) NOT NULL COMMENT '电影id',
  `like` int(11) DEFAULT '0' COMMENT '喜欢人数',
  `hate` int(11) DEFAULT '0' COMMENT '讨厌人数',
  `talk_num` int(11) DEFAULT '0' COMMENT '谈论次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COMMENT='影评属性表';

-- ----------------------------
-- Records of fr_attribute
-- ----------------------------
INSERT INTO `fr_attribute` VALUES ('1', '13', '1', '1', '0', '1');
INSERT INTO `fr_attribute` VALUES ('2', '14', '1', '3', '0', '13');
INSERT INTO `fr_attribute` VALUES ('3', '1', '1', '0', '0', '0');
INSERT INTO `fr_attribute` VALUES ('4', '2', '7', '20', '0', '0');
INSERT INTO `fr_attribute` VALUES ('5', '3', '1', '0', '0', '0');
INSERT INTO `fr_attribute` VALUES ('6', '4', '1', '0', '0', '0');
INSERT INTO `fr_attribute` VALUES ('7', '5', '2', '0', '0', '1');
INSERT INTO `fr_attribute` VALUES ('8', '6', '9', '1', '0', '0');
INSERT INTO `fr_attribute` VALUES ('9', '7', '12', '0', '0', '0');
INSERT INTO `fr_attribute` VALUES ('10', '8', '20', '0', '0', '0');
INSERT INTO `fr_attribute` VALUES ('11', '9', '22', '0', '0', '0');
INSERT INTO `fr_attribute` VALUES ('12', '10', '1', '0', '1', '0');
INSERT INTO `fr_attribute` VALUES ('13', '15', '4', '10', '0', '4');
INSERT INTO `fr_attribute` VALUES ('14', '16', '42', '1', '0', '1');

-- ----------------------------
-- Table structure for fr_talk
-- ----------------------------
DROP TABLE IF EXISTS `fr_talk`;
CREATE TABLE `fr_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `fr_id` int(11) NOT NULL COMMENT '影评id',
  `movie_id` int(11) NOT NULL COMMENT '电影id',
  `user_id` varchar(50) NOT NULL COMMENT '评论人id',
  `content` text COMMENT '评论内容',
  `date` datetime DEFAULT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8 COMMENT='影评评论表';

-- ----------------------------
-- Records of fr_talk
-- ----------------------------
INSERT INTO `fr_talk` VALUES ('1', '14', '1', '141842291', '秀起来', '2018-04-09 11:37:10');
INSERT INTO `fr_talk` VALUES ('2', '14', '1', '141842291', '我就是我不一样的烟火', '2018-04-09 11:37:47');
INSERT INTO `fr_talk` VALUES ('3', '14', '1', '141842291', '我居然是第三！！！', '2018-04-09 11:38:01');
INSERT INTO `fr_talk` VALUES ('4', '14', '1', '141842291', '( • ̀ω•́ )✧', '2018-04-09 11:38:22');
INSERT INTO `fr_talk` VALUES ('5', '14', '1', '141842291', '!!!∑(ﾟДﾟノ)ノ', '2018-04-09 11:38:28');
INSERT INTO `fr_talk` VALUES ('6', '14', '1', '141842291', '(╥╯^╰╥)', '2018-04-09 11:38:31');
INSERT INTO `fr_talk` VALUES ('7', '14', '1', '141842291', '╮(╯﹏╰）╭', '2018-04-09 11:38:36');
INSERT INTO `fr_talk` VALUES ('8', '14', '1', '141842291', '(╥╯^╰╥)', '2018-04-09 11:38:38');
INSERT INTO `fr_talk` VALUES ('9', '14', '1', '141842291', '1231', '2018-04-09 11:38:40');
INSERT INTO `fr_talk` VALUES ('10', '14', '1', '141842291', '(>ω･* )ﾉ', '2018-04-09 11:38:46');
INSERT INTO `fr_talk` VALUES ('11', '14', '1', '141842291', 'o(╯□╰)o', '2018-04-09 11:40:01');
INSERT INTO `fr_talk` VALUES ('12', '14', '1', '141842291', '哭瞎(ಥ_ಥ) ', '2018-04-09 12:46:21');
INSERT INTO `fr_talk` VALUES ('13', '14', '1', '141842291', '嘻嘻徐', '2018-04-09 12:52:46');
INSERT INTO `fr_talk` VALUES ('14', '5', '2', '141842291', '第一', '2018-04-09 12:53:06');
INSERT INTO `fr_talk` VALUES ('15', '13', '1', '141842291', '<script>alert(123)</script>', '2018-04-09 14:06:28');
INSERT INTO `fr_talk` VALUES ('16', '15', '4', '958752538', '顶顶顶', '2018-04-20 14:47:13');
INSERT INTO `fr_talk` VALUES ('17', '15', '4', '141842261', '新人你好呀', '2018-04-23 16:17:46');
INSERT INTO `fr_talk` VALUES ('18', '15', '4', '141842291', '姐姐来啦！', '2018-04-23 16:19:18');
INSERT INTO `fr_talk` VALUES ('19', '15', '4', '958752538', '谢谢姐姐', '2018-04-23 16:19:51');
INSERT INTO `fr_talk` VALUES ('20', '16', '42', '141842291', '哈哈哈哈', '2018-04-27 16:36:10');

-- ----------------------------
-- Table structure for history_record
-- ----------------------------
DROP TABLE IF EXISTS `history_record`;
CREATE TABLE `history_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '账户',
  `movie_id` int(11) NOT NULL COMMENT '电影id',
  `date` datetime DEFAULT NULL COMMENT '观看时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='看电影历史记录表';

-- ----------------------------
-- Records of history_record
-- ----------------------------
INSERT INTO `history_record` VALUES ('5', '958752538', '1', '2018-04-28 19:08:38');
INSERT INTO `history_record` VALUES ('6', '958752538', '4', '2018-04-20 17:35:46');
INSERT INTO `history_record` VALUES ('7', '958752538', '2', '2018-04-20 17:35:50');
INSERT INTO `history_record` VALUES ('8', '958752538', '22', '2018-04-20 17:36:03');

-- ----------------------------
-- Table structure for language
-- ----------------------------
DROP TABLE IF EXISTS `language`;
CREATE TABLE `language` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(11) NOT NULL COMMENT '语言',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='语言表';

-- ----------------------------
-- Records of language
-- ----------------------------
INSERT INTO `language` VALUES ('1', '国语');
INSERT INTO `language` VALUES ('2', '粤语');
INSERT INTO `language` VALUES ('3', '英语');
INSERT INTO `language` VALUES ('4', '韩语');
INSERT INTO `language` VALUES ('5', '日语');
INSERT INTO `language` VALUES ('6', '西班牙语');
INSERT INTO `language` VALUES ('7', '德语');
INSERT INTO `language` VALUES ('8', '法语');
INSERT INTO `language` VALUES ('9', '其他');

-- ----------------------------
-- Table structure for movie
-- ----------------------------
DROP TABLE IF EXISTS `movie`;
CREATE TABLE `movie` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `movie_name` varchar(50) DEFAULT NULL COMMENT '电影名字',
  `area` varchar(50) DEFAULT NULL COMMENT '地区',
  `language` varchar(50) DEFAULT NULL COMMENT '语言',
  `type` varchar(50) DEFAULT NULL COMMENT '电影类型',
  `director` varchar(50) DEFAULT NULL COMMENT '导演',
  `release_date` date DEFAULT NULL COMMENT '上映日期',
  `update_date` date DEFAULT NULL COMMENT '更新日期',
  `introuduce` varchar(50) DEFAULT NULL COMMENT '电影简介',
  `year` varchar(50) DEFAULT NULL COMMENT '年份',
  `plot` text COMMENT '剧情',
  `source` enum('exist','disappear') DEFAULT NULL COMMENT '是否有片源 1 有 0 无',
  `source_path` varchar(255) DEFAULT NULL COMMENT '资源路径',
  `source_from` int(11) DEFAULT NULL COMMENT '资源出处 1官方0个人 （暂定）',
  `integral` int(11) DEFAULT NULL COMMENT '积分（暂定）',
  `movie_img` varchar(50) DEFAULT 'default.jpg' COMMENT '电影图片',
  `img_head` varchar(50) DEFAULT 'default.jpg' COMMENT '电影图片',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8 COMMENT='电影表';

-- ----------------------------
-- Records of movie
-- ----------------------------
INSERT INTO `movie` VALUES ('1', '宣告黎明的露之歌', '日本', '日语', '动画', '汤浅政明234', '2017-05-19', '2018-03-27', '郁郁寡欢的少年“海”与人鱼少女“露”相遇的故事 ', '2017-05-19', '因父母离异从东京搬回老家渔港町，郁郁寡欢的少年凯，唯一的心灵寄托就是将自己创作的音乐传到网上，某天他遇见了人鱼少女露，而人鱼自古就被视为带来灾祸的存在。', 'exist', '1.mp4', '1', '0', '1524882856.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('2', '你的名字', '日本', '日语', '动画', '新海诚', '2016-12-02', '2018-02-05', '爱无力治愈病例：命运感的爱情，才是时空旅行的意义 ', '2016', ' 在远离大都会的小山村，住着巫女世家出身的高中女孩宫水三叶（上白石萌音 配音）。校园和家庭的原因本就让她充满烦恼，而近一段时间发生的奇怪事件，又让三叶摸不清头脑', 'exist', '2.mp4', '1', '0', '2.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('3', '春宵苦短', '日本', '日语', '动画', '汤浅政明', '2017-04-07', '2018-03-22', '著名作家森见登美彦最高的代表作 ', '2017', '古色古香的京都，奇妙的恋爱之旅正在展开。春天，黑发少女踏上寻找伪电气白兰地的夜之旅，学长却被从天而降的锦鲤砸晕；夏天，少女在旧书市寻找心爱的童年绘本，学长则面临残酷的火锅地狱', 'disappear', null, '0', '0', '3.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('4', '声之形', '日本', '日语', '动画', '山田尚子', '2016-09-07', '2017-09-07', '修复电视格式，没法用声音交流传达自己的想法 ', '2016', ' 《声之形》是由日本漫画家大今良时创作的漫画作品，故事主题是没法用声音交流来传达自己的想法，故事以男主角石田将也的视角出发，讲述一位患有听力障碍的女生西宫硝子在学校备受霸凌，得不到友情和关爱，后来男主良心发现，帮助她找到幸福', 'exist', '4.mp4', '1', '0', '4.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('5', '老师我可以喜欢你吗', '日本', '日语', '爱情', '三木孝浩', '2017-10-28', '2018-03-20', '最楚楚动人的丝丝,再加上师生恋,就够了! ', '2017', ' 响受朋友小惠之托，将情书放入关矢老师的鞋箱，但却误放到伊藤老师的鞋箱里…一次的阴错阳差，造成一段师生恋…', 'disappear', null, null, null, '5.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('6', '热血高校2', '日本', '日语', '动作', ' 三池崇史 ', '2009-09-04', '2018-04-28', '传说将华丽地落下帷幕 ', '2009', '讲述了暴力斗殴不再仅仅局限于铃兰高中内部，凤仙学院——铃兰曾经的死对头蠢蠢欲动，一帮后人为给死去的老大报仇又再度发起挑战的故事。 在GPS与芹泽军团的恶战之后，铃兰高中进入短暂的和平期。芹泽多摩雄（山田孝之 饰）让出老大宝座，泷谷源治（小栗旬 饰）则持续向恐怖的林田惠（深水元基 饰）发起冲击', 'exist', '1524894983.mp4', null, null, '6.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('7', '哪啊哪啊神去村', '日本', '日语', '喜剧', ' 矢口史靖 ', '2014-03-22', '2018-03-03', '日本知识青年上山下乡的励志喜剧', '2014', '讲述了刚刚高中毕业对未来没有任何期望的主人公平野勇气（染谷将太饰），在母亲以及导师的联手策划下，被送往没有手机信号的深山里的“神去村”从事伐木工作。从一开始抵触“伐木生活”到逐渐喜欢上了这个自然村落，从勇气口中终于说出一句当地村民的口头语“哪啊哪啊”。', 'disappear', null, null, null, '7.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('8', '虐杀器官', '日本', '日语', '科幻', '村濑修功 ', '2017-02-12', '2017-10-15', '改编自伊藤计划原作的同名小说 ', '2017', '《屠杀器官》的舞台将设定在赛拉耶佛核爆后的地球，此时部分落后国家爆发了一连串内战及民族冲突，在这血腥风暴之中，美国特殊情报部队克拉维斯发现了在阴影之下有一名叫约翰·保罗的美国人参与其中，约翰路过之处必定掀起腥风血雨，这让众人疯狂的 「屠杀器官」到底是什么？ 虐杀器官剧场版讲述的是在911之后，世界因核爆、战争而破败不堪', 'disappear', null, null, null, '8.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('9', '嫌疑犯X的献身', '日本', '日语', '悬疑', ' 西谷弘 ', '2008-10-04', '2014-06-23', '浪漫到无可救药但却愚钝的爱 ', '2008', '天才物理学家汤川学（福山雅治 饰）此次又受内海薰（柴崎幸 饰）所托，调查一件不可思议命案。受害人富樫慎二被发现死在河边，面目指纹尽毁。最大的嫌疑人是他的前妻花冈靖子（松雪泰子 饰），然而靖子却有着完美的不在场证明', 'disappear', null, null, null, '9.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('10', '浪矢解忧杂货店', '日本', '日语', '剧情', ' 广木隆一 ', '2018-02-02', '2018-02-19', '又一部治愈的电影，这鸡汤我先干为敬！ ', '2018', ' 2012年的一个夜晚，自幼在某孤儿院长大的敦也（山田凉介 饰）、翔太（村上虹郎 饰）和幸平（宽一郎 饰）对某女企业家实施了抢劫，由于汽车意外抛锚，他们仓皇逃到了一家早已关张的杂货店内躲藏。这家名为浪矢杂货店的铺子此前由风趣和善的老人浪矢雄治（西田敏行 饰）经营，数十年前他出于玩笑帮小朋友们解答各种困扰，后来名气越来越大，以至于有人专程向他咨询人生中的难题', 'disappear', null, null, null, '10.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('11', '釜山行', '韩国', '韩语', '灾难', '延尚昊 ', '2016-07-20', '2016-09-30', '由延相昊执导的灾难片 ', '2016', '影片讲述主人公单亲爸爸石宇与女儿秀安乘坐KTX高速列车往釜山，列车上由一位少女身上带来的僵尸病毒开始肆虐且不断扩散，倾刻间列车陷入灾难的故事。一场神秘的病毒疫情爆发，让整个韩国陷入紧急状态！ 一种尚未识别的病毒席卷了整个国家，韩国政府宣布全国戒严。而那些乘坐火车前往釜山（成功避免病毒爆发的城市）的人们，则必须为自己的生存而战…… 453公里，从首尔到釜山的距离。 为了保护亲人朋友，他们将为生存殊死一搏！ 要活命，快上车！', 'disappear', null, null, null, '11.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('12', '触不到的恋人', '韩国', '韩语', '爱情', '李铉升 ', '2000-01-19', '2016-08-23', '缘分天注定 ', '2000', '讲述了分别在两个时段住进同一栋屋子里的男女主人公，通过一个神秘的邮箱互相书信交流的跨越时空的爱情故事。 韩星贤（李政宰饰）搬到一个海边小屋，他为房屋取了一个意大利文的名字“Il Mare”，意思是“海”。整理房子时，他在信箱发现了一封内容很奇怪的信，信上写着“我是你搬来前的上一个房客，如果有收到我的信，请寄来”，更奇怪的是这封信寄出的日期是在1999年，也就是两年后——而星贤现在的时间是在1997年。星贤就立刻回信给这莫名其妙的人，询问是怎么回事', 'disappear', null, null, null, '12.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('13', '男与女', '韩国', '韩语', '爱情', '李胤基 ', '2016-02-25', '2018-08-02', '全度妍、孔刘主演的爱情电影 ', '2016', '在芬兰国际学校里偶遇的尚敏（全度妍饰）和基洪（孔刘饰）被安排一同前往北校区，由于暴雪导致交通瘫痪，二人更是在路上迷失方向，在一望无际的雪山里虽然不知道对方姓名，但却互相依偎挨过困境，但也就此告别。8个月后回到首尔的基洪回到一切自己正常的生活，但他内心挂念的尚敏却突然出现，二人注定陷入一段不可告人的关系中 。', 'disappear', null, null, null, '13.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('14', '我的野蛮女友', '韩国', '韩语', '爱情', '郭在容 ', '2001-07-27', '2015-04-05', '你是我的全智贤，我只偷看你一眼 ', '2001', '讲述了既美丽又野蛮的女主与大学生牵牛之间发生的有趣而又浪漫爱情故事。牵牛（车太贤饰）第一次与她（全智贤饰）邂逅，就毫无浪漫感可言。在地铁上她就语出惊人，喝的醉醺醺的她大声训斥不自觉给老人让座的人，吐完之后就软趴趴地倒了下来。善良的牵牛上前接住，于是就只好背着她到旅馆投宿了', 'disappear', null, null, null, '14.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('15', '记忆之夜', '韩国', '韩语', '悬疑', '张恒俊 ', '2017-11-29', '2017-12-29', '一场车祸，一次金融危机，两个家庭悲剧！ ', '2017', ' 《记忆之夜》是一部讲述兄弟故事的惊悚片，哥哥因绑架事件失去了记忆，弟弟则一直在为哥哥找寻记忆 ', 'disappear', null, null, null, '15.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('16', '勇敢者游戏', '美国', '英语', '动作', '杰克·卡斯丹 ', '2017-12-20', '2018-03-14', '爆笑十足也惊险刺激，非常不错的爆米花电影 ', '2017', '讲述了四名性格迥异的高中生意外穿越到险象环生的勇敢者游戏中，变身成为与自身性格外貌截然不同的游戏角色。面对丛林猛兽的袭击和邪恶势力的追捕，四人必须在冒险家（道恩·强森 Dwayne Johnson 饰）的带领下，战胜重重危机，赢得游戏，才能获得重返现实的机会的故事', 'disappear', null, null, null, '16.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('17', '疯狂动物城', '美国', '英语', '动画', ' 拜伦·霍华德  瑞奇·摩尔 ', '2016-03-04', '2016-08-14', '第89届奥斯卡金像奖 ', '2016', '影片以兔子朱迪通过自己努力奋斗，终于完成自己儿时的梦想，成为动物城的一个警察的故事为主线，传达了要摒弃偏见、为了梦想不抛弃不放弃的精神。 疯狂动物城是一座独一无二的现代动物都市。每种动物在这里都有自己的居所，比如富丽堂皇的撒哈拉广场，或者常年严寒的冰川镇。', 'disappear', null, null, null, '17.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('18', '雄狮', '美国', '英语', '剧情', '加斯·戴维斯 ', '2017-01-19', '2017-02-17', '爱是指引前进的光明，妮可基德曼主演 ', '2017', '五岁时萨罗在印度的火车上不小心与家人走散，他在加尔各答的街头流浪了几个星期，之后被送进一所孤儿院，并被澳大利亚的一对夫妇收养。25年后，他开始凭借隐约的记忆，用谷歌地图去寻找过去的家', 'disappear', null, null, null, '18.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('19', '爱乐之城', '美国', '英语', '爱情', ' 达米安·沙泽勒 ', '2016-12-25', '2018-01-25', '二十年后，再一次用爱情感动全世界 ', '2016', '故事发生在当代的洛杉矶，寂寥的小演员米娅的志向是女演员兼剧作家，她沉迷老电影明星。她在华纳兄弟的片场当咖啡师，经常翘班去试镜，如果真的接到戏，哪怕再小的角色也会让她欣喜若狂', 'disappear', null, null, null, '19.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('20', '敦刻尔克', '美国', '英语', '战争', ' 克里斯托弗·诺兰 ', '2017-07-21', '2018-01-17', '面对战争与大海，个体可以何等无助，又何等坚持 ', '2017', '故事改编自著名的二战军事事件“敦刻尔克大撤退”。二战初期，40万英法盟军被敌军围困于敦刻尔克的海滩之上，面对敌军步步逼近的绝境，形势万分危急。英国政府和海军发动大批船员，动员人民起来营救军队', 'disappear', null, null, null, '20.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('21', '末日崩塌', '美国', '英语', '灾难', '布拉德·佩顿 ', '2015-06-02', '2015-08-26', '加州大地震', '2015', '影片讲述了一场大地震从洛杉矶一直延伸至旧金山，巨石强森饰演的角色是消防部空军中队的一名飞行员，他必须驾驶飞机努力寻找并营救自己的女儿', 'disappear', null, null, null, '21.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('22', '泰坦尼克号', '美国', '英语', '爱情', '詹姆斯·卡梅隆 ', '1997-12-19', '2016-10-29', '影史上的精品之作，同类影片中的极品 ', '1997', '片以1912年泰坦尼克号邮轮在其处女启航时触礁冰山而沉没的事件为背景，描述了处于不同阶层的两个人——穷画家杰克和贵族女露丝抛弃世俗的偏见坠入爱河，最终杰克把生命的机会让给了露丝的感人故事', 'exist', '22.mp4', null, null, '22.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('23', '橄榄球传奇', '美国', '英语', '运动', '安德鲁·欧文', '2015-10-16', '2016-05-17', '根据橄榄球员托尼·内森的真实故事改编 ', '2015', '讲述了20世纪70年代的美国，刚刚废除种族隔离，但南部地区的种族冲突仍然十分严重。伯明翰伍德劳恩高中的天才橄榄球球员托尼·内森（迦勒·卡斯提尔 Caleb Castille 饰）与其他黑人同学因种族冲突而无法加入球队。橄榄球教练坦迪（尼克·毕肖普 Nic Bishop 饰）率先组成黑人白人学生混合的球队', 'disappear', null, null, null, '23.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('24', '寻梦环球记', '美国', '英语', '动画', ' 李·昂克里奇  阿德里安·', '2017-11-22', '2018-03-01', '供奉的遗像是牵引家人回家的通道，驻留的记忆是保持亡灵存续的神力', '2017', '热爱音乐的米格尔（安东尼·冈萨雷兹 Anthony Gonzalez 配音）不幸地出生在一个视音乐为洪水猛兽的大家庭之中，一家人只盼着米格尔快快长大，好继承家里传承了数代的制鞋产业。一年一度的亡灵节即将来临，每逢这一天，去世的亲人们的魂魄便可凭借着摆在祭坛上的照片返回现世和生者团圆', 'disappear', null, null, null, '24.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('25', '王牌特工', '美国', '英语', '动作', '马修·沃恩 ', '2015-01-29', '2017-11-27', '逼要这么装 特工电影的新标杆 向特工们致敬 ', '2015', ' 哈里（科林·费斯 Colin Firth 饰）是英国秘密特工组织金士曼中的一员，某次行动中，他的战友不幸牺牲，哈里将一枚徽章和一个电话号码交给了战友年幼的小儿子艾格西（亚历克斯·尼科洛夫 Alex Nikolov 饰），叮嘱他将来如果遇到了什么麻烦可以拨打这个号码，然而，这样的机会只能使用一次', 'disappear', null, null, null, '25.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('26', '一切尽失', '美国', '英语', '灾难', ' J·C·尚多尔 ', '2013-10-18', '2014-01-30', '自然无常，生命至上', '2013', ' 罗伯特·雷德福饰演一名搭乘帆船出海的男子，先是和一艘货船相撞、失去对外联络的装备，又遇上了暴风雨，他只好弃船求生……', 'disappear', null, null, null, '26.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('27', '时间之旅', '美国', '英语', '记录', '泰伦斯·马力克 ', '2016-09-07', '2017-10-03', '在天地万物间来趟时间之旅 ', '2016', '水母、鳄鱼胚胎、星云、犹他州的峡谷和木卫三；第一缕生命的信号、细菌、细胞起源、初恋感觉、意识诞生、人类进化、生与死——本片将带你走过漫长时间之旅，从宇宙的诞生到最终灭亡', 'disappear', null, null, null, '27.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('28', ' 当幸福来敲门', '美国', '英语', '励志', ' 加布里尔·穆奇诺 ', '2006-12-15', '2011-12-20', '幸福明天就会来临 ', '2006', '克里斯•加纳（威尔•史密斯饰）用尽全部积蓄买下了高科技治疗仪，到处向医院推销，可是价格高昂，接受的人不多。就算他多努力都无法提供一个良好的生活环境给妻儿，妻子（桑迪•纽顿）最终选择离开家', 'disappear', null, null, null, '28.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('29', 'C罗 ', '英国', '英语', '记录', 'Anthony Wonke ', '2015-11-09', '2015-11-13', '“有人爱我，就有人恨我”你就是你！罗纳尔多 ', '2015', ' 这是一份送给球迷的礼物，他并不想通过这部纪录片去说服谁，讨厌他的人或者对他有偏见的人，他只是讲述自己的生活、自己的想法、自己的家人和朋友，展现出自己最柔软而真实的部分给那些爱他的人们。你们赠予我爱，我还以真实，如此温柔，谢谢这份礼物！', 'disappear', null, null, null, '29.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('30', '一呼一吸', '英国', '英语', '爱情', '安迪·瑟金斯 ', '2017-10-27', '2018-01-04', '爱如空气，一呼一吸', '2017', '阳光开朗，热爱冒险的罗宾（安德鲁·加菲尔德 Andrew Garfield 饰）邂逅了温柔美丽的戴安娜（克莱尔·福伊 Claire Foy 饰），两人坠入了情网并且很快便决定携手步入婚姻的殿堂', 'disappear', null, null, null, '30.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('31', '银翼杀手2049', '英国', '英语', '科幻', '丹尼斯·维伦纽瓦 ', '2017-10-06', '2017-12-26', '生命的解放，带来人性的脆弱，应然与实然的未来 ', '2017', ' 故事发生在大断电30年后。复制人K（瑞恩·高斯林 Ryan Gosling 饰）是新一代的银翼杀手，在如今的世界里，人类和复制人之间的界限划分的更加明确，复制人从刚一制造出来就被灌输了服务于人类的思想，绝对不被允许产生人类的感情', 'disappear', null, null, null, '31.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('32', '星际穿越', '英国', '英语', '科幻', '克里斯托弗·诺兰 ', '2014-11-07', '2016-07-18', '重新洗版，并增加电视格式。诺兰导演科幻巨制', '2014', '星际穿越》（Interstellar）是克里斯托弗·诺兰执导的一部原创科幻冒险电影，由马修·麦康纳、安妮·海瑟薇、杰西卡·查斯坦及迈克尔·凯恩主演，基于知名理论物理学家基普·索恩的黑洞理论经过合理演化之后，加入人物和相关情节改编而成', 'disappear', null, null, null, '32.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('33', '比利·林恩的中场战事', '英国', '英语', '战争', '李安 ', '2016-10-11', '2017-01-31', '李安挑战观影极限之作 ', '2016', '《比利·林恩的中场战事》根据本·芳汀（Ben Fountain）同名小说改编，被认为是伊拉克战争版的“第22条军规”（Catch-22）。故事主角是一名19岁的得克萨斯男孩比利·林（Billy Lynn），他在加入美军步兵部队后，被派往伊拉克战场', 'disappear', null, null, null, '33.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('34', '赎罪', '英国', '英语', '战争', '乔·怀特 ', '2007-09-07', '2018-02-03', '灵魂悲歌，用精致优美打动你的古典爱情悲剧 ', '2007', '1935年夏天，来自一个宽裕的英国小康家庭的13岁的少女布里奥妮（斯奥里兹•罗南 饰）刚刚开始尝试写作,想象力丰富。一天，她暗中发现仆人的儿子——罗比•特纳（詹姆斯•迈克沃伊 饰）和她姐姐塞西利亚（凯特•奈特莉 饰）之间有暧昧关系，而且发现他给她写了一封充满情色意味的情书', 'disappear', null, null, null, '34.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('35', '神探夏洛克', '英国', '英语', '悬疑', '道格拉斯·麦金农 ', '2016-01-01', '2016-01-03', '神探夏洛克：2016新年特别篇 ', '2016', '福尔摩斯回归原著维多利亚时代，调查一宗神秘的幽灵新娘谋杀案。Thomas Ricoletti惊讶地发现，几小时前自杀身亡的妻子Emelia Ricoletti身着旧婚纱再次出现了，她的幽灵带着强烈的复仇欲望在街上徘徊', 'disappear', null, null, null, '35.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('36', '地球脉动', '英国', '英语', '记录', '艾雷斯泰·法瑟吉尔 ', '2006-02-27', '2016-11-25', '行星地球 , 我们的地球 , 地球无限 ', '2006', 'BBC曾经制作出《蓝色星球》的纪录片摄影团队，再次集结奉上了这部堪称难以超越的经典纪录片《地球脉动》。从南极到北极，从赤道到寒带，从非洲草原到热带雨林，再从荒凉峰顶到深邃大海，难以数计的生物以极其绝美的身姿呈现在世人面前', 'disappear', null, null, null, '36.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('37', '愿上帝宽恕我们', '西班牙', '西班牙语', '悬疑', '罗德里戈·索罗戈延 ', '2016-10-28', '2018-01-13', '西语悬疑惊悚胜在氛围营造 好精彩！', '2016', ' 马德里，2011年，有史以来最热的八月，一百五十万朝圣者正等待着教宗的到访。整个城市正处于安全部队的监视之下。几个老妇女被残酷奸杀，两名不想打交道的督察不得不彼此合作解决此案，他们接到明确指令：尽快解决此案，但是不能声张。究竟隐藏了什么事呢', 'disappear', null, null, null, '37.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('38', '杰出公民', '西班牙', '西班牙语', '喜剧', '加斯顿·杜帕拉特 ', '2016-09-08', '2017-04-13', '本片获得第73届威尼斯国际电影节最佳男演员奖 ', '2016', ' 沉默寡言的小说家丹尼尔·曼托瓦尼（奥斯卡·马丁内兹饰）与成功有着一种说不清、道不明的关系：成功给他带来了同时代的其他作家难以企及的财富，却让他担忧自己已经没有成名前那宝贵的棱角。 五年后，这位定居巴塞罗那的作家依然炙手可热，他不得不让助手礼貌地回绝一堆盛情邀请，直到收到了一封来自距离布宜诺斯艾利斯车程七小时的家乡小镇萨拉斯的信。丹尼尔所有创作的灵感全部来自这片故土，但却已经四十年没有回乡——正像他自己所说的那样，他从未离开家乡，但也从未回归', 'disappear', null, null, null, '38.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('39', '午夜巴塞罗那', '西班牙', '西班牙语', '爱情', '伍迪·艾伦 ', '2008-09-19', '2014-08-10', '美丽的巴塞罗，惬意的时光，一段错综复杂的迷情爱恋 ', '2008', '故事发生在美丽的巴塞罗那，美国女孩维姬（丽贝卡·豪尔 Rebecca Hall 饰）和克里斯蒂娜（斯嘉丽·约翰逊 Scarlett Johansson 饰）在度假时认识了名声并不太好的艺术家胡安（哈维尔·巴登 Javier Bardem 饰）。性格迥异的两姐妹在第二次遇到胡安后，对于胡安结伴去奥维耶多过周末的的邀请，热情奔放的克里斯蒂娜一下子就爽快答应了，而即将结婚的维姬却感到深深的不安', 'disappear', null, null, null, '39.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('40', '佣兵传奇 ', '西班牙', '西班牙语', '战争', '奥古斯汀·迪亚兹·亚内', '2006-09-01', '2016-05-04', '以草根的放大镜窥探昨日帝国的衰荣 ', '2006', 'Diego Alatriste (Viggo Mortensen饰)不是一个高尚的人，但他却是拥有荣誉与勇猛的男子汉。在17世纪的西班牙军队中担任上校，他逐渐成为这个国家帝国战争的英雄', 'disappear', null, null, null, '40.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('41', '天使爱美丽', '法国', '法语', '喜剧', ' 让-皮埃尔·热内 ', '2001-04-25', '2013-12-15', '过去的只能重现而不会再来 ', '2001', '艾米莉（奥黛丽·塔图 Audrey Tautou 饰）有着别人看来不幸的童年——父亲给她做健康检查时，发现她心跳过快，便断定她患上心脏病，从此艾米莉与学校绝缘', 'disappear', null, null, null, '41.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('42', '这个杀手不太冷', '法国', '法语', '剧情', ' 吕克·贝松 ', '1994-09-14', '2014-12-28', '吕克·贝松经典电影 ', '1994', '昂（让•雷诺饰）是名孤独的职业杀手，受人雇佣。一天，邻居家小姑娘马蒂尔达（纳塔丽•波特曼饰)敲开他的房门，要求在他那里暂避杀身之祸', 'disappear', null, null, null, '42.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('43', '触不可及', '法国', '法语', '剧情', ' 奥利维·那卡什', '2011-11-02', '2012-02-23', '当“奥巴马”遇到“霍金”', '2011', ' 因为一次跳伞事故，白人富翁菲利普Philippe（弗朗索瓦·克鲁塞 François Cluzet 饰）瘫痪在床，欲招聘一名全职陪护。由于薪酬高，应聘者云集，个个舌灿莲花，却无法打动他的心。直到黑人德希斯Driss（奥玛·赛 Omar Sy 饰）的出现才让他作出决定', 'disappear', null, null, null, '43.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('44', '露西亚的情人', '法国', '法语', '爱情', '胡里奥·密谭 ', '2001-08-24', '2013-04-29', '法国女人热情奔放,爱情戏也拍得自然动人 ', '2001', '风流倜傥的小说家洛伦佐（Tristán Ulloa 饰）在遭遇创作瓶颈之际邂逅了自己的崇拜者露西亚（Paz Vega 帕兹·维嘉 饰）——一名在咖啡馆工作的女招待。她点燃了洛伦佐沉寂已久的激情，两人开始一段愉快刺激的同居生活。偶然机会，洛伦佐得知自己拥有一个女儿，那是他一年前在海边夜晚的所留下的艳遇结晶', 'disappear', null, null, null, '44.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('45', ' 只要在一起 ', '法国', '法语', '爱情', '克劳德·贝里', '2007-03-21', '2014-02-11', '情节简单,却因为真实而感人 ', '2007', '26岁的美丽女孩卡米尔（奥黛丽·塔图饰）是一个清洁女工，非常热爱绘画。但多年来的孤独和童年时不堪回首的经历，让她对自己的才华和人生都产生怀疑', 'disappear', null, null, null, '45.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('46', '地雷区 ', '德国', '德语', '战争', ' 马丁·赞里维特 ', '2015-12-03', '2016-04-16', '丹麦德国合拍，根据二战中的真实历史事件改编 ', '2015', '第二次世界大战之后，丹麦当局迫使成千上万的德国战俘排除埋在丹麦海滩的地雷。这些年轻人只得用赤裸的双手，挖掘数以百万计的残余地雷', 'disappear', null, null, null, '46.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('47', '极速风流 ', '德国', '德语', '运动', ' 朗·霍华德 ', '2013-09-13', '2014-03-06', '雨道狂飙惨烈竞速 ', '2013', '詹姆斯（克里斯·海姆斯沃斯 Chris Hemsworth 饰）和尼基（丹尼尔·布鲁赫 Daniel Brühl 饰）最初相识在三级方程式锦标赛的赛道上，二人有着截然不同的风格', 'disappear', null, null, null, '47.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('48', '窃听风暴', '德国', '德语', '悬疑', '弗洛里安·亨克尔·冯·多', '2006-03-23', '2012-02-14', '第79届奥斯卡最佳外语片奖 ', '2006', '984年的东德，整个社会笼罩在国家安全局的高压统治之下，特工魏斯曼（乌尔利希·穆厄 饰）奉命监听剧作家德莱曼（塞巴斯蒂安·寇治 饰）和他妻子演员克里斯蒂娜（玛蒂娜·杰蒂克 饰）的生活，监听过程中，魏斯曼渐渐对这家人的生活产生了兴趣，开始暗中帮助他们', 'disappear', null, null, null, '48.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('49', '海蒂和爷爷', '德国', '德语', '剧情', '阿兰·葛斯彭纳 ', '2015-12-10', '2016-08-12', '根据Johanna Spyri 的小说改编的 ', '2015', '这部电影是2015年德国瑞士合拍的。海蒂自幼失去双亲，於是就由姨妈照顾，后来由於姨妈要到外处工作，海蒂便去投靠住在阿尔卑斯山上的爷爷，爷爷性情古怪，跟村民不合，长久以来就离群索居，独自住在山区的木屋中，当海蒂来到爷爷家，她的纯真及率直终於融化了爷爷的心', 'disappear', null, null, null, '49.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('50', '朗读者', '德国', '英语', '爱情', '史蒂芬·戴德利 ', '2008-12-12', '2012-09-06', '回来不是为了留下，而是为了重新出发 ', '2008', '15岁的少年米夏·伯格（大卫·克劳斯 David Kross 饰）偶遇36岁的中年神秘女列车售票员汉娜（凯特·温丝莱特 Kate Winslet 饰），后来两个发展出一段秘密的情人关系', 'disappear', null, null, null, '50.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('51', ' 三傻大闹宝莱坞', '印度', '英语', '励志', '拉库马·希拉尼 ', '2011-12-08', '2016-06-14', '三分傻气，七分传奇 ', '2011', ' 本片根据印度畅销书作家奇坦·巴哈特（Chetan Bhagat）的处女作小说《五点人》（Five Point Someone）改编而成。法兰（马德哈万 R Madhavan 饰）、拉杜（沙曼·乔希 Sharman Joshi 饰）与兰乔（阿米尔·汗 Aamir Khan 饰）是皇家工程学院的学生，三人共居一室，结为好友。在以严格著称的学院里，兰乔是个非常与众不同的学生，他不死记硬背，甚至还公然顶撞校长“病毒”（波曼·伊拉尼 Boman Irani 饰），质疑他的教学方法', 'disappear', null, null, null, '51.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('52', ' 我的个神啊', '印度', '英语', '喜剧', ' 拉吉库马尔·希拉尼 ', '2015-05-22', '2015-05-21', '外星醉汉PK地球神 ', '2015', '影片通过一个外星人在地球的奇幻旅行，对印度的文化进行了深刻的探讨。 贾古（安努舒卡·莎玛 Anushka Sharma 饰）和男友相恋多年，感情十分要好的两人终于决定步入婚姻的殿堂，然而，一场意外的突然降临让贾古所期望的一切都化为了泡影，因此，伤心欲绝的贾古决定返回家乡，成为了一名记者', 'disappear', null, null, null, '52.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('53', '未知死亡', '印度', '其他', '悬疑', 'A.R. Murugadoss', '2008-12-25', '2014-03-16', '记忆，是爱你最安全的方式 ', '2008', '电信大亨杰辛哈尼亚（阿米尔·汗 Aamir Khan 饰）为人低调，某日却见报纸爆料他向一个叫卡尔帕谢蒂（阿米尔·汗 Aamir Khan 饰）的小广告模特大胆示爱，气愤之余杰辛哈尼亚去向这个造谣生事的女孩儿问责，谁料他途中看见一个美丽善良的女孩机智地帮助残疾儿童，而这个女孩就是卡尔帕谢蒂，卡尔帕谢蒂并不认识杰辛哈尼亚，一场误会令别人以为她就是大亨的女友', 'disappear', null, null, null, '53.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('54', '看了又看', '印度', '其他', '爱情', 'Nitya Mehra ', '2016-09-09', '2017-03-04', '《降临》+《人生遥控器》 ', '2016', '《看了又看》是一部印度爱情电影。', 'disappear', null, null, null, '54.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('55', '小萝莉的猴神大叔 ', '印度', '英语', '动作', ' 卡比尔·汗 ', '2015-07-17', '2015-10-28', '爱，发声 ', '2015', '《小萝莉的猴神大叔》又名（《猴神老哥》），讲述了一个拥有虔诚宗教信仰的单纯印度男人（萨尔曼·汗 饰）承诺帮一个巴基斯坦哑女（哈尔莎莉·马尔霍特拉 饰）与父母重聚的故事', 'disappear', null, null, null, '55.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('56', '摔跤吧！爸爸', '印度', '英语', '运动', '涅提·蒂瓦里', '2016-12-23', '2017-06-09', '热血逗乐竞技惊心动魄的宝莱坞电影 ', '2016', '摔跤吧！爸爸》是由尼特什·提瓦瑞执导、阿米尔·汗领衔主演的印度电影。 《摔跤吧！爸爸》根据真人真事改编，讲述曾经的摔跤冠军辛格培养两个女儿成为女子摔跤冠军，打破印度传统的励志故事', 'disappear', null, null, null, '56.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('57', '狐妖小红娘剧场版', '大陆', '国语', '动画', '王昕 ', '2015-12-11', '2015-12-12', null, '2015', '《狐妖小红娘》的故事围绕人与妖之间的爱情展开。根据古典小说记载，世上有人有妖，妖会与人相恋，妖寿命千万年，人的寿命有限，人死了，妖活着', 'disappear', null, null, null, '57.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('58', '中国惊奇先生', '大陆', '国语', '动画', ' 李豪凌 ', '2015-07-25', '2015-07-31', 'TV版总集篇 ', '2015', '著名道人袁天罡集合麻衣派与茅山派道术所创的“神鬼七杀令”在现今时代悄然出世…… 北京某胡同内，近来常发生下班女子惨遭色狼强暴、碎尸案件', 'disappear', null, null, null, '58.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('59', '捉妖记2 ', '大陆', '国语', '喜剧', '许诚毅 ', '2018-02-16', '2018-01-05', '萌萌的胡巴又来了！！ ', '2018', '上一集与胡巴分别后，天荫带着小岚踏上寻父之路，在义薄云天的天师堂堂主云大哥的帮助下，二人得知天荫父亲宋戴天的护妖轨迹；而重回永宁村的胡巴再度被妖王追杀，颠沛流离逃亡时结识大赌徒屠四谷和一只妖怪，三人一起过着相依为命的生活，但又因屠四谷欠下的巨额赌债横生诸多波折', 'disappear', null, null, null, '59.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('60', '南极之恋', '大陆', '国语', '爱情', '吴有音 ', '2018-02-02', '2018-03-24', '空气戒指求婚那段真挺感人的 ', '2018', '南极，一场坠机，婚庆公司老板吴富春（赵又廷 饰）和高空物理学家荆如意（杨子姗 饰）相遇，两个毫无共同语言的男女在南极腹地无人区冒险生存75天', 'disappear', null, null, null, '60.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('61', '芳华', '大陆', '国语', '战争', '冯小刚', '2017-12-15', '2018-02-17', '一个始终不被善待的人，最能识别善良，也最珍惜善良”练功房、成分、“解放”……太多感同身受的情节', '2017', '根据严歌苓同名小说改编，讲述了上世纪七十到八十年代充满理想和激情的军队文工团，一群正值芳华的青春少年，经历着成长中的爱情萌发与充斥变数的人生命运', 'disappear', null, null, null, '61.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('62', '羞羞的铁拳 ', '大陆', '国语', '喜剧', '宋阳 ', '2017-09-30', '2017-12-22', '拳拳都是笑点', '2017', '该片根据同名话剧改编，主要讲述了一个搏击选手艾迪生和一个体育女记者马小，因为一场意外的电击灵魂互换的搞笑爱情故事。 靠打假拳混日子的艾迪生（艾伦 饰），本来和正义感十足的体育记者马小（马丽 饰）是一对冤家，没想到因为一场意外的电击，男女身体互换', 'disappear', null, null, null, '62.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('63', '战狼', '大陆', '国语', '动作', ' 吴京 ', '2015-04-02', '2017-08-20', '太极宗师来当兵 ', '2015', '《战狼》是由吴京执导的现代军事战争片，该片由吴京，余男、倪大红、斯科特·阿金斯 、周晓鸥等主演。该影片属于国内首部3D动作战争电影，历时七年全力打造，《战狼》真实呈现了一场中外边境战争，也让堪称“东方之狼”的特种兵战队及高能战士首次登陆大银幕 。讲述的是小人物成长为拯救国家和民族命运的孤胆英雄的传奇故事', 'disappear', null, null, null, '63.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('64', '唐人街探案', '大陆', '国语', '悬疑', '陈思诚 ', '2015-12-31', '2016-02-15', '喜剧外表下掩藏着杀人案里子 ', '2015', '《唐人街探案》由万达影视传媒有限公司、上海骋亚影视文化传媒有限公司出品；湖南芒果娱乐有限公司、合一影业有限公司等联合出品', 'disappear', null, null, null, '64.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('65', '无问西东', '大陆', '国语', '爱情', '李芳芳 ', '2018-01-12', '2018-01-05', '2018最值得期待的华语电影', '2018', '如果提前了解了你所要面对的人生，你是否还会有勇气前来？吴岭澜、沈光耀、王敏佳、陈鹏、张果果，几个年轻人满怀诸多渴望，在四个非同凡响的时空中一路前行', 'disappear', null, null, null, '65.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('66', '重返·狼群 ', '大陆', '国语', '记录', '亦风 ', '2017-06-16', '2017-11-05', '人与自然本是一体 ', '2017', '美女画家李微漪在一次草原采风中意外收养了狼王遗孤，为其取名格林，并带回成都喂养。但繁华的都市无法容纳一匹野性的草原狼，李微漪带着格林重返草原，一路追寻狼群的踪迹，历经严寒酷雪', 'disappear', null, null, null, '66.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('67', '喜马拉雅天梯 ', '大陆', '国语', '记录', '萧寒 ', '2015-10-16', '2015-11-18', '登顶对游客来说是生命的高潮 ', '2015', '《喜马拉雅天梯》是首部纪录珠峰攀登全程的4K超高清极限纪实电影，本片由萧寒、梁君健两名资深纪录片导演联合执导，拍摄制作历时4年，将镜头对准西藏年轻的登山向导，完整记录了珠峰北坡登顶全过程', 'disappear', null, null, null, '67.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('68', '我在故宫修文物', '大陆', '国语', '记录', '萧寒', '2016-12-16', '2017-01-24', '国宝复活术，搬上大荧幕的纪录片', '2016', '《我在故宫修文物》第一次将镜头对准了故宫的文物修复师们，他们已经存在了几百年，却始终不为人知。书中以口述的形式撰写了12位顶级文物修复师的对历史、对人生的回顾和感悟，同时也是一本故宫几百年文物修复历史的缩影', 'disappear', null, null, null, '68.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('69', '闪光少女', '大陆', '国语', '喜剧', '王冉 ', '2017-07-20', '2017-09-22', '你人多你有理？来自二次元的少数派报告 ', '2017', ' 陈惊（徐璐饰）是一个古灵精怪、敢爱敢恨的姑娘，江湖人称“神经”。作为音乐附中的奇葩人物，陈惊人缘欠佳只有男闺蜜“油渣”（彭昱畅饰）甘愿为她鞍前马后', 'disappear', null, null, null, '69.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('70', '前任3：再见前任', '大陆', '国语', '爱情', '田羽生 ', '2017-12-29', '2018-03-06', '那些在你生命中出现又离去的前任，是为了让你成为更好的自己 ', '2017', '一对好基友孟云（韩庚 饰）和余飞（郑恺 饰）跟女友都因为一点小事宣告分手，并且“拒绝挽回，死不认错”。两人在夜店、派对与交友软件上放飞人生第二春，大肆庆祝“黄金单身期”，从而引发了一系列好笑的故事', 'disappear', null, null, null, '70.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('71', '那些年', '台湾', '国语', '爱情', '九把刀', '2011-08-19', '2015-12-17', '一起来寻找最纯真的感动吧 ', '2011', ' 青春是一场大雨。即使感冒了，还盼望回头再淋它一次。人生就是不停的战斗，在还没有获得女神青睐时，左手永远都只是辅助！！！ ', 'disappear', null, null, null, '71.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('72', '海角七号', '台湾', '国语', '爱情', '魏德圣 ', '2008-08-22', '2012-01-03', '每个人心中都有一封寄不出的情书，不管是到天涯，还是。。。 ', '2008', '依傍大海的美丽小镇恒春上，音乐梦想在台北遭受打击的年轻人阿嘉（范逸臣）冷眼旁观着一切。虽在民代主席后父的帮忙下成了代班邮差，信件和邮包却被他乱堆在房间，而对一个寄自日本的无法送至的邮包，他也没按老邮差茂伯（林宗仁）的吩咐将之退回，而是私自拆阅查看：里面除一张泛黄的少女照片，是几封写于60年前的信件。其时日本在二战中战败，在台日籍教师（中孝介）随日军撤退时遗弃了相约私奔的女友，归日途中，他将爱意和悔意化为文字，但信件直到去世才被其女儿代为寄出', 'disappear', null, null, null, '72.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('73', '十二夜', '台湾', '国语', '记录', 'Raye ', '2013-11-29', '2014-05-06', '九把刀首次监制纪录片,流浪狗的收容所日记 ', '2013', ' 九把刀首次监制的电影，与演员隋棠兼任出品人。 　　一个叫Raye的女新导演拍的一部推广「以领养代替购买」的精神，去帮助处境可怜的流浪狗的影片', 'disappear', null, null, null, '73.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('74', '视线之外', '台湾', '国语', '动画', '虞雅婷 ', '2010-07-31', '2013-10-03', '如果一直走下去，你还会记得你是来找小狗的吗 ', '2010', '在一个阳光明媚、清新爽朗的日子里，一尘不染的街旁，茂密的大树随风摇曳着安逸与清凉。穿着天蓝色短裙、戴着白色帽子的小女孩和她的小狗悠然漫步，全然不知危险将近', 'disappear', null, null, null, '74.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('75', '我的少女时代', '台湾', '国语', '爱情', '陈玉珊 ', '2015-11-19', '2015-12-20', '青春里有过你，总不负少女一场 ', '2015', '《我的少女时代》是由陈玉珊执导，宋芸桦、王大陆、李玉玺、简廷芮主演的青春校园爱情电影。该片以90年代的台湾高中为背景，讲述了平凡少女林真心和校园老大徐太宇的初恋故事', 'disappear', null, null, null, '75.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('76', '追龙', '香港', '粤语', '动作', '王晶', '2017-09-30', '2018-01-23', '久违的王晶与英雄片，这是港式浪漫化的江湖 ', '2017', '电影讲述了香港现代史上两大著名狠角色大毒枭跛豪（甄子丹饰）、五亿探长雷洛（刘德华饰）的传奇故事。 上世纪六七十年代，香港由英国殖民，权势腐败、社会混乱', 'disappear', null, null, null, '76.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('77', '美人鱼', '香港', '粤语', '爱情', '周星驰', '2016-02-08', '2016-05-28', '自成一格，周星驰版美人鱼 ', '2016', '《美人鱼》是由周星驰执导的喜剧片，邓超、罗志祥 、张雨绮、林允等领衔出演。该片讲述了富豪轩少和美人鱼珊在填海工程的开发过程中发生的一系列故事。', 'disappear', null, null, null, '77.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('78', '叶问', '香港', '粤语', '动作', '叶伟信 ', '2015-12-24', '2016-03-24', '甄子丹时代谢幕前的最后突围 ', '2015', '该片以叶问晚年为背景，讲述了一位豪杰落入种种困境后通过自强不息最终成为一代宗师的历程。 1959年，叶问（甄子丹饰）与张永成（熊黛林饰）将大儿子叶准送回广东，小儿子叶正继续在香港读书', 'disappear', null, null, null, '78.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('79', '春娇救志明', '香港', '粤语', '爱情', ' 彭浩翔 ', '2017-04-28', '2017-06-12', '爱情是几兆亿运算都无解的问题 ', '2014', ' 2009年，春娇志明后巷邂逅，演绎烟火世界中的极致浪漫。 2017年，饮食男女分分合合，志明仍未长大，春娇急需爱情救火。 相恋多年的春娇与志明，渐渐从“可遇不可求”的惊涛骇浪变成了普通的情侣。面临“中女危机”的春娇，深感颜值危机，可志明却还如同孩子一般无法长大', 'disappear', null, null, null, '79.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('80', '杀破狼2', '香港', '粤语', '动作', '郑保瑞', '2015-06-18', '2015-10-11', '你们穿西装的真能打 ', '2015', '香港卧底探员志杰（吴京 饰）不惜变成瘾君子，潜伏犯罪集团侦查幕后主脑洪先生（古天乐 饰）；直属上司兼叔父华哥（任达华 饰）决定终止行动，志杰却身份败露且人间蒸发。华哥违抗命令私自搜寻，辗转得悉志杰已被关进泰国某监狱，遂只身前赴营救。 泰国警察阿猜（托尼·贾 饰）为筹钱医治患血癌的女儿转任狱警，负责看管志杰', 'disappear', null, null, null, '80.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('81', '初恋这件小事', '其他', '其他', '爱情', '普特鹏·普罗萨卡·那·萨', '2010-08-12', '2012-05-24', '女屌丝蜕变为白天鹅的小清新之作，曾经疯传校内网 ', '2010', '中生小水相貌平平，家境一般，所以在学校里并不受重视。但是她心地善良，又有一群死党，日子过得倒也开心。某天，她遇见了帅气的学长阿亮（马里奥·毛瑞尔 Mario Maurer 饰），春心萌动，无 法遏制。她喜欢看他踢足球，看他拍照，如痴如狂。上英语课时，她不停地给死党传纸条 表达对阿亮的爱慕', 'disappear', null, null, null, '81.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('82', '完美陌生人', '其他', '其他', '喜剧', '保罗·格诺维瑟', '2016-02-11', '2016-07-06', '相爱没有那么容易，每个人有他的手机 ', '2016', '讲诉七个常年的好朋友聚在一起吃晚餐。忽然他们决定与对方分享每一个短信的内容，包括他们收到的电子邮件和电话，由此许多秘密引出了彼此看似平淡无奇生活之下暗涌的层层情绪和欲望，这里面朋友和夫妻的关系变成一个让人力有不逮的存在，表面上是大家是惯常社会中常见的亲密的朋友，但最开始时不时出现的对他人生活的蠢蠢欲动的好奇和毒舌已经显示出大家压抑不住的情绪攻击，到后来一切的发展愈加失控，随之而来的真相让大家几乎全线崩溃', 'disappear', null, null, null, '82.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('83', ' 博格对战麦肯罗', '其他', '其他', '运动', '扬努斯·梅兹·彼得森', '2017-09-08', '2018-01-09', '网坛史上名局，真的印证了阿加西的话，好一段迷你人生', '2017', '希亚·拉博夫要在银幕上正儿八经干坏事了，他宣布加盟传记片《博格/麦肯罗》，演绎网球球坛坏小子麦肯罗，瑞典演员斯维尔·古德纳松将扮演“冰山”球手博格，该片由扬努斯·梅兹·彼得森执导，讲述两位网球天才之间的巅峰对决', 'disappear', null, null, null, '83.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('84', '我的天使', '其他', '法语', '爱情', '哈里·克莱文', '2017-05-03', '2017-12-09', '小成本电影，但视觉冲击力极强;无形胜有型，几处都拍的很惊艳 ', '2017', '比利时导演哈里·克莱文作品。描述隐形人魔幻爱情的电影，轰动「世界三大奇幻影展之首」布鲁塞尔国际奇幻影展，将潜藏人类心底的愿望「隐形」，激荡出更多的幻想与欲望', 'disappear', null, null, null, '84.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('85', '曾经', '其他', '英语', '爱情', '约翰·卡尼 ', '2007-03-23', '2013-09-05', '渺小的我们曾经相遇 ', '2007', '这是一部由音乐开启的爱尔兰电影。　 卖花女（Marketa Irglova）被街头艺人（Glen Hansard饰）的音乐所吸引，开始了一段荡气回肠的浪漫故事', 'disappear', null, null, null, '85.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('93', '肖申克的救赎', '英国', '国语', '悬疑', '弗兰克·德拉邦特', '1994-09-23', '2018-04-28', '拍出了同类作品罕见的人情味和温馨感觉', '1994-09-23', ' 1946年，年青的银行家安迪（蒂姆•罗宾斯 Tim Robbins 饰）被冤枉杀了他的妻子和其情人，这意味着他要在肖申克的监狱渡过余生。 银行家出身的安迪很快就在监狱里很吃得开，他懂得如何帮助狱卒逃税，懂得如何帮监狱长将他收到的非法收入“洗白”，很快，安迪就成为了狱长的私人助理。 一名小偷因盗窃入狱，他知道安迪妻子和她情人的死亡真相，兴奋的安迪找到了狱长，希望狱长能帮他翻案。虚伪的狱长表面上答应了安迪，暗中却派人杀死了小偷，只因他想安迪一直留在监狱帮他做帐。 安迪知道真相后，决定通过自己的救赎去获得自由！', 'exist', '1524900339.mp4', null, null, '1524900339.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('94', '现在去见你', '韩国', '韩语', '爱情', '李章焄', '2018-03-04', null, '有笑点有泪点的电影，孙艺珍怎么还是那么美', '2018-03-04', '本片根据市川拓司小说《相约在雨季》改编，讲述男子(苏志燮饰)的妻子秀雅(孙艺珍饰)去世前告诉他，自己将在下一个雨季回来。果不其然，雨季来临之际，妻子如约而至。可奇怪的是，妻子似乎忘记了他和家庭。影片预计明年上映', null, null, null, null, '1524901316.jpg', 'default.jpg');
INSERT INTO `movie` VALUES ('95', '后来的我们', '大陆', '国语', '爱情', '刘若英', '2018-04-28', null, '超有才华的奶茶导演处女作！期待！爱情文艺片！ ', '2018-04-28', '这是一个爱情故事，关于一对异乡漂泊的年轻人。 十年前，见清和小晓偶然地相识在归乡过年的火车上。两人怀揣着共同的梦想，一起在北京打拼，并开始了一段相聚相离的情感之路。 十年后，见清和小晓在飞机上再次偶然重逢…… 命运似乎是一个轮回。在一次次的偶然下，平行线交叉，再平行，故事始终有“然后”。可后来的他们，学会如何去爱了吗？', null, null, null, null, '1524902023.jpg', 'default.jpg');

-- ----------------------------
-- Table structure for movie_attribute
-- ----------------------------
DROP TABLE IF EXISTS `movie_attribute`;
CREATE TABLE `movie_attribute` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `movie_id` int(11) NOT NULL COMMENT '电影id',
  `popularity` int(11) DEFAULT '0' COMMENT '人气',
  `collect_num` int(11) DEFAULT '0' COMMENT '收藏人数',
  `score` decimal(2,1) DEFAULT '0.0' COMMENT '评分',
  `year` varchar(50) DEFAULT NULL COMMENT '年份',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8 COMMENT='电影属性表';

-- ----------------------------
-- Records of movie_attribute
-- ----------------------------
INSERT INTO `movie_attribute` VALUES ('1', '1', '537', '2', '8.0', '2018');
INSERT INTO `movie_attribute` VALUES ('2', '2', '0', '1', '4.0', '2016');
INSERT INTO `movie_attribute` VALUES ('3', '3', '40', '1', '6.0', '2017');
INSERT INTO `movie_attribute` VALUES ('4', '4', '1', '2', '6.0', '2016');
INSERT INTO `movie_attribute` VALUES ('5', '5', '2', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('6', '6', '6', '3', '9.0', '2009');
INSERT INTO `movie_attribute` VALUES ('7', '7', '0', '3', '8.0', '2014');
INSERT INTO `movie_attribute` VALUES ('8', '8', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('9', '9', '2', '0', '0.0', '2008');
INSERT INTO `movie_attribute` VALUES ('10', '10', '0', '0', '8.0', '2018');
INSERT INTO `movie_attribute` VALUES ('11', '11', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('12', '12', '8', '0', '0.0', '2000');
INSERT INTO `movie_attribute` VALUES ('13', '13', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('14', '14', '0', '0', '0.0', '2001');
INSERT INTO `movie_attribute` VALUES ('15', '15', '0', '1', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('16', '16', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('17', '17', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('18', '18', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('19', '19', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('20', '20', '0', '0', '9.0', '2017');
INSERT INTO `movie_attribute` VALUES ('21', '21', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('22', '22', '0', '1', '5.0', '1997');
INSERT INTO `movie_attribute` VALUES ('23', '23', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('24', '24', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('25', '25', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('26', '26', '0', '0', '0.0', '2013');
INSERT INTO `movie_attribute` VALUES ('27', '27', '1', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('28', '28', '0', '0', '0.0', '2006');
INSERT INTO `movie_attribute` VALUES ('29', '29', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('30', '30', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('31', '31', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('32', '32', '0', '0', '0.0', '2014');
INSERT INTO `movie_attribute` VALUES ('33', '33', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('34', '34', '0', '0', '0.0', '2007');
INSERT INTO `movie_attribute` VALUES ('35', '35', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('36', '36', '0', '0', '0.0', '2006');
INSERT INTO `movie_attribute` VALUES ('37', '37', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('38', '38', '1', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('39', '39', '0', '0', '4.0', '2008');
INSERT INTO `movie_attribute` VALUES ('40', '40', '0', '0', '0.0', '2006');
INSERT INTO `movie_attribute` VALUES ('41', '41', '0', '0', '0.0', '2001');
INSERT INTO `movie_attribute` VALUES ('42', '42', '2', '0', '0.0', '1994');
INSERT INTO `movie_attribute` VALUES ('43', '43', '0', '0', '0.0', '2011');
INSERT INTO `movie_attribute` VALUES ('44', '44', '0', '0', '0.0', '2001');
INSERT INTO `movie_attribute` VALUES ('45', '45', '0', '0', '7.0', '2007');
INSERT INTO `movie_attribute` VALUES ('46', '46', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('47', '47', '0', '0', '0.0', '2013');
INSERT INTO `movie_attribute` VALUES ('48', '48', '0', '0', '0.0', '2006');
INSERT INTO `movie_attribute` VALUES ('49', '49', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('50', '50', '0', '0', '0.0', '2008');
INSERT INTO `movie_attribute` VALUES ('51', '51', '0', '0', '0.0', '2011');
INSERT INTO `movie_attribute` VALUES ('52', '52', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('53', '53', '0', '0', '0.0', '2008');
INSERT INTO `movie_attribute` VALUES ('54', '54', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('55', '55', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('56', '56', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('57', '57', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('58', '58', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('59', '59', '0', '0', '0.0', '2018');
INSERT INTO `movie_attribute` VALUES ('60', '60', '0', '0', '0.0', '2018');
INSERT INTO `movie_attribute` VALUES ('61', '61', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('62', '62', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('63', '63', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('64', '64', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('65', '65', '0', '0', '0.0', '2018');
INSERT INTO `movie_attribute` VALUES ('66', '66', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('67', '67', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('68', '68', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('69', '69', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('70', '70', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('71', '71', '1', '0', '0.0', '2011');
INSERT INTO `movie_attribute` VALUES ('72', '72', '0', '0', '0.0', '2008');
INSERT INTO `movie_attribute` VALUES ('73', '73', '0', '0', '0.0', '2013');
INSERT INTO `movie_attribute` VALUES ('74', '74', '0', '0', '0.0', '2010');
INSERT INTO `movie_attribute` VALUES ('75', '75', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('76', '76', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('77', '77', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('78', '78', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('79', '79', '0', '0', '0.0', '2014');
INSERT INTO `movie_attribute` VALUES ('80', '80', '0', '0', '0.0', '2015');
INSERT INTO `movie_attribute` VALUES ('81', '81', '0', '0', '0.0', '2010');
INSERT INTO `movie_attribute` VALUES ('82', '82', '0', '0', '0.0', '2016');
INSERT INTO `movie_attribute` VALUES ('83', '83', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('84', '84', '0', '0', '0.0', '2017');
INSERT INTO `movie_attribute` VALUES ('85', '85', '0', '0', '0.0', '2007');
INSERT INTO `movie_attribute` VALUES ('86', '93', '5', '1', '0.0', '1994');
INSERT INTO `movie_attribute` VALUES ('87', '94', '2', '0', '0.0', '2018-03-04');
INSERT INTO `movie_attribute` VALUES ('88', '95', '0', '0', '0.0', '2018-04-28');

-- ----------------------------
-- Table structure for movie_commentary
-- ----------------------------
DROP TABLE IF EXISTS `movie_commentary`;
CREATE TABLE `movie_commentary` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `movie_id` int(11) NOT NULL COMMENT '电影id',
  `user_id` varchar(50) NOT NULL COMMENT '用户id',
  `date` datetime NOT NULL COMMENT '短评时间',
  `content` varchar(100) NOT NULL COMMENT '评论内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8 COMMENT='电影短评表';

-- ----------------------------
-- Records of movie_commentary
-- ----------------------------
INSERT INTO `movie_commentary` VALUES ('1', '1', '141842261', '2018-04-02 18:06:37', '我是大帅哥');
INSERT INTO `movie_commentary` VALUES ('2', '1', '141842261', '2018-04-02 18:08:00', '很喜欢的一部片，很美，很治愈');
INSERT INTO `movie_commentary` VALUES ('3', '3', '141842261', '2018-04-03 09:06:25', '没有资源啊！好难受');
INSERT INTO `movie_commentary` VALUES ('4', '4', '141842261', '2018-04-03 10:24:00', '我是第一名');
INSERT INTO `movie_commentary` VALUES ('5', '1', '141842261', '2018-04-03 11:10:36', '还有谁？');
INSERT INTO `movie_commentary` VALUES ('6', '1', '141842291', '2018-04-03 11:19:14', '我也很感兴趣啊');
INSERT INTO `movie_commentary` VALUES ('7', '1', '141842291', '2018-04-03 11:20:10', 'what?');
INSERT INTO `movie_commentary` VALUES ('8', '1', '141842291', '2018-04-03 11:25:06', '666');
INSERT INTO `movie_commentary` VALUES ('9', '1', '141842261', '2018-04-03 11:25:58', '还是你们会玩啊');
INSERT INTO `movie_commentary` VALUES ('10', '1', '141842261', '2018-04-03 11:26:18', '真的不要太厉害');
INSERT INTO `movie_commentary` VALUES ('11', '1', '141842261', '2018-04-03 11:26:54', '听说评论超过15个字可以加积分，不知道是不是真的');
INSERT INTO `movie_commentary` VALUES ('12', '1', '141842291', '2018-04-03 11:38:43', '我也想上评论啊');
INSERT INTO `movie_commentary` VALUES ('13', '1', '141842291', '2018-04-03 11:40:45', '天哪，我的呢？');
INSERT INTO `movie_commentary` VALUES ('14', '7', '141842291', '2018-04-03 11:52:51', '1楼');
INSERT INTO `movie_commentary` VALUES ('15', '1', '141842291', '2018-04-03 14:29:43', '今天面试成功，是招财猫的客服哦666的');
INSERT INTO `movie_commentary` VALUES ('16', '5', '141842291', '2018-04-03 14:50:41', 'wow,禁断的爱情，令人着迷');
INSERT INTO `movie_commentary` VALUES ('17', '1', '141842291', '2018-04-04 08:46:07', '2018-4-4第一');
INSERT INTO `movie_commentary` VALUES ('18', '48', '141842291', '2018-04-04 09:36:41', '一楼');
INSERT INTO `movie_commentary` VALUES ('19', '1', '141842261', '2018-04-04 12:49:05', 'shit');
INSERT INTO `movie_commentary` VALUES ('20', '3', '141842261', '2018-04-04 12:54:20', '秀秀秀');
INSERT INTO `movie_commentary` VALUES ('21', '1', '141842261', '2018-04-08 13:09:48', '1');
INSERT INTO `movie_commentary` VALUES ('22', '1', '958752538', '2018-04-20 14:39:14', '听说增加了更改头像功能厉害了');
INSERT INTO `movie_commentary` VALUES ('23', '1', '958752538', '2018-04-20 14:39:41', '话说楼下的俩个不会是百合吧');
INSERT INTO `movie_commentary` VALUES ('24', '93', '958752538', '2018-04-28 15:45:21', '这是刚上传的电影吗？好好看啊\n');

-- ----------------------------
-- Table structure for msg_record
-- ----------------------------
DROP TABLE IF EXISTS `msg_record`;
CREATE TABLE `msg_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `content` text COMMENT '内容',
  `status` enum('false','true') DEFAULT 'false' COMMENT '查看状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COMMENT='消息记录表';

-- ----------------------------
-- Records of msg_record
-- ----------------------------
INSERT INTO `msg_record` VALUES ('9', '141842291', '消息回复', '尊敬的会员你好，你的资讯问题消息已被查看！谢谢支持本网站', 'false');
INSERT INTO `msg_record` VALUES ('10', '958752538', '消息回复', '尊敬的会员你好，你的网站的推荐功能消息已被查看！谢谢支持本网站', 'false');
INSERT INTO `msg_record` VALUES ('11', '958752538', '消息回复', '尊敬的会员你好，你的头号玩家消息已被查看！谢谢支持本网站', 'false');
INSERT INTO `msg_record` VALUES ('12', '958752538', '消息回复', '尊敬的会员你好，你的头号玩家消息已被查看！谢谢支持本网站', 'false');
INSERT INTO `msg_record` VALUES ('13', '958752538', '消息回复', '尊敬的会员你好，你的《起跑线》消息已被查看！谢谢支持本网站', 'false');
INSERT INTO `msg_record` VALUES ('14', '141842291', '消息回复', '尊敬的会员你好，你的资讯问题消息已被查看！谢谢支持本网站', 'false');

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL COMMENT '新闻标题',
  `date` date DEFAULT NULL COMMENT '发布日期',
  `content` text COMMENT '内容',
  `img` varchar(100) DEFAULT NULL COMMENT '宣传海报',
  `point` varchar(100) DEFAULT NULL COMMENT '简要',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='资讯表';

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news` VALUES ('3', '加热比海盗来啦！', '2018-04-26', '<p>\r\n\r\n<p>《加勒比海盗》是全球最具知名度和最卖座的电影系列之一，自2003年首部曲横空出世以来，已累计推出四部作品，全球总票房突破37亿美元。《加勒比海盗》以冒险奇幻交织的动人故事、恢弘神秘的海盗世界为全球观众带来独一无二的观影体验，约翰尼·德普领衔的豪华群星阵容更是锦上添花，影史经典地位毋庸置疑。距四部曲问世六年之后，第五部终于王者归来，在保留系列经典元素的同时，势必将带来更多惊喜。日前迪士尼影业正式宣布，《加勒比海盗5：死无对证》将于今年5月26日在内地同步北美隆重上映，中国影迷能够第一时间与全球观众一起先睹为快。</p><p>此外，迪士尼影业同时宣布对中国影迷的独家喜讯，《加勒比海盗5：死无对证》全球首映礼将于5月11日在上海迪士尼度假区盛大举行，约翰尼·德普、奥兰多·布鲁姆、杰弗里·拉什、哈维尔·巴登、布兰顿·思怀兹（Brenton Thwaites）等群星主演将集体现身辉映红毯，零距离约会中国影迷。一手打造整个系列的王牌制片人杰瑞·布鲁克海默也将携导演艾斯彭·山德伯格(Espen Sandberg)、乔阿吉姆·罗恩尼(Joachim R?nning)齐齐出席。红毯仪式后，《加勒比海盗5：死无对证》正片放映则将在位于上海迪士尼度假区迪士尼小镇内的华特迪士尼大剧院震撼上演。全明星阵容还将现身上海迪士尼乐园内的加勒比海盗主题园区“宝藏湾”，将首映活动推向最高潮。“宝藏湾”是全球迪士尼乐园中首个以加勒比海盗为主题的园区，也是上海迪士尼乐园独有的园区，因此这次《加勒比海盗5：死无对证》首映派对选址在宝藏湾举行，绝对堪称不二之选。</p><p></p><p><img alt=\"\" src=\"http://inews.gtimg.com/newsapp_bt/0/1463309288/641\"></p><p></p><p>《加勒比海盗5：死无对证》演员阵容强大星光璀璨。约翰尼·德普演绎的杰克船长作为系列灵魂人物，魅力独具，深受全球影迷热爱，并为德普赢得奥斯卡影帝提名殊荣；奥斯卡最佳男配角哈维尔·巴登此次全新加盟饰演亡灵萨拉查船长，奥斯卡影帝杰弗里·拉什再度化身与杰克船长相爱相杀的巴博萨船长。几位奥斯卡巨星飙戏“互怼”势必成为一大看点。前三部曲“铁三角”成员奥兰多·布鲁姆也惊喜回归，饰演备受全球观众喜爱的“小铁匠”威尔·特纳，牵动观众更多回忆；两名“小鲜肉”布兰顿·思怀兹、卡雅·斯考达里奥也全新加盟，化身杰克船长此番大冒险的亲密搭档，皆是举重若轻的关键角色。</p><p>影片幕后也有金奖团队保驾护航。一手打造《加勒比海盗》系列的好莱坞金牌制作人杰瑞·布鲁克海默继续掌控全局，此番担任导演的艾斯彭·山德伯格、乔阿吉姆·罗恩尼曾共同执导荣获奥斯卡最佳外语片提名的《孤筏重洋》，新鲜血液的注入将为《加勒比海盗5：死无对证》带来令人耳目一新的惊喜创意。</p><p>此次《加勒比海盗5：死无对证》不但将呈现一位令人胆战心惊的全新反派，更将完美融合幽默、惊悚、诙谐、史诗等《加勒比海盗》系列深受全球观众追捧的经典元素，呈现出2017年最独一无二的观影体验。在《加勒比海盗5：死无对证》中，亡灵萨拉查船长率领亡灵军队逃出百慕大三角区，扬言杀尽所有海盗，头号目标就是杰克船长。为对抗命运，杰克船长将起航寻找传说中海神波塞冬的三叉戟，因为这是能统治整个海洋的神器。波澜壮阔的海战争霸和史诗冒险就此展开。在先期曝光的物料中，影片对奇幻海洋背景的刻画巨细无靡、叹为观止；亡灵军队与亡灵鲨鱼等设计惊艳、见之难忘；海战场景的规模也更上一层楼。此外，青年时代的杰克船长在电影系列中首度登场也是一大惊喜，英气勃发又颜值爆表，令粉丝期待不已。</p>\r\n\r\n<br></p>', '1524732631.jpg', '5年的史诗巨作');
INSERT INTO `news` VALUES ('4', '劳拉O(∩_∩)O哈哈~', '2018-04-26', '<p>\r\n\r\n<p>史克威尔艾尼克斯日前在东京电玩展（Tokyo Game Show）上公布了《古墓丽影之崛起 20周年庆》的最新宣传片。</p><p></p><div><img alt=\"TGS 2016 20\" src=\"http://img1.gtimg.com/gamezone/pics/hv1/226/132/2128/138407086.jpg\" title=\"Image: http://img1.gtimg.com/gamezone/pics/hv1/226/132/2128/138407086.jpg\"></div><p></p><p></p><div><img alt=\"TGS 2016 20\" src=\"http://img1.gtimg.com/gamezone/pics/hv1/228/132/2128/138407088.jpg\"></div><p></p><p>本作登陆PS4平台，预定将于10月11日上市。游戏将在PC和XBO两个版本的基础上新增下载内容“血缘”（Blood Ties）、新模式“劳拉的噩梦”（Lara’s Nightmare）、新难度“极限求生”（Extreme Survivor）、五款“经典劳拉”（Classic Lara）皮肤、以及《古墓丽影3》的主题装备和武器，并支持PlayStation VR和PlayStation 4 Pro。</p><p>而PlayStation 4 Pro将提供三种显示方案：4K电视用户可使用“4K模式”（4K Mode）；没有4K电视的话，可选用“高帧率模式”（High Frame Rate Mode），游戏将保持1080p和解锁帧率的运行状态；还有一个“增强图形模式”（Enhanced Visuals Mode），游戏将保持1080p和每秒30帧的运行状态，并实现渲染技术全开。</p>\r\n\r\n<br></p>', '1524733360.jpg', '古墓古墓古墓重要的事情说三遍！！！');
INSERT INTO `news` VALUES ('5', '环太平洋', '2018-04-26', '<p>\r\n\r\n<p>《环太平洋：雷霆再起》（《环太2》） 正在全国热映中，首周末3天共拿下4.09亿票房，成为3月的“票房炸弹”。</p><p><img alt=\"\" src=\"http://n.sinaimg.cn/translate/132/w1080h652/20180328/46Xd-fysqfni0761078.jpg\" title=\"Image: http://n.sinaimg.cn/translate/132/w1080h652/20180328/46Xd-fysqfni0761078.jpg\"></p><p>想看电影的小心心，是不是因为女生节被掏空的钱包而望而止步呢？好像只能吃土了。然而朋友，以实力宠粉著称的京东爸爸怎么可能舍得让你们吃土？即使昨天一场冰雹让泥土飘满花香。</p>\r\n\r\n<br></p>', '1524733595.jpg', '特效风暴来袭');
INSERT INTO `news` VALUES ('6', '宫老的片', '2018-04-26', '<p>\r\n\r\n<p>《起风了》是宫崎骏执导的第11部动画长片，继2008年<a target=\"_blank\" rel=\"nofollow\" href=\"http://movie.mtime.com/73219/\">《悬崖上的金鱼姬》</a>后时隔五年宫崎骏推出的又一动画长篇力作，根据他本人在《Model Graphix》(大日本绘画)的2009年4月号-2010年1月号中连载的同名人气漫画改编。预告片中我们可以看到，这部二战题材的动画，保持着宫崎骏一贯细腻清新的画风，为我们展现了那个时代日本社会的宏大画卷。虽然本片基本没有描写战争的场景，却蕴含了对战争、科技的反思，女主人公掩面溅血的残忍戏份，更让人隐隐感到强大的催泪指数。在本片的完成报告会上，宫崎骏曾坦言“看自己的作品看哭了，这还是第一次……经过常年的积累，在不可思议的机缘和幸运的帮助下，终于完成了一部电影。所以不自觉的落泪。”</p><p>&nbsp;</p><p></p><p>　　本片主人公原形是日本航空之父、零式战机的开发者堀越二郎，同时融合了同时代文学家堀辰雄的生平。故事描绘了二郎的技师生活以及他与不幸少女菜穗子的相遇与爱情，以及他从一个单纯喜爱航空机的少年，成长为一个为了战争不得不陷入“是否应该制造（杀人机器）零战”的工程师。在展示主人公堀越二郎为实现儿时梦想而努力成为飞机设计师的同时，也揭露了战争年代堀越二郎“想设计优美的飞机”的单纯设计理念与先进的零式战斗机在战争中起的作用之间的矛盾。</p><p>&nbsp;</p><p></p><p><a target=\"_blank\" rel=\"nofollow\" href=\"http://movie.mtime.com/186716/posters_and_images/2310927/\"><img alt=\"\" src=\"http://img31.mtime.cn/pi/2013/06/24/175821.34657859.jpg\"></a></p><p><a target=\"_blank\" rel=\"nofollow\" href=\"http://movie.mtime.com/186716/posters_and_images/2306957/\"><img alt=\"\" src=\"http://img31.mtime.cn/pi/2013/06/17/134359.99364287.jpg\"></a></p><p>男女主人公的凄美爱情将成为本片的一大泪弹，片名“風立ちぬ”正是影片中男主角邂逅女主角时吟诵的诗句“風立ちぬ、いざ生きめやも（纵有疾风起，人生不言弃）”</p><p>&nbsp;</p><p></p><p>　　为片中二郎配音的是现年52岁，以<a target=\"_blank\" rel=\"nofollow\" href=\"http://movie.mtime.com/23779/\">《新世纪福音战士》</a>系列扬名的动画导演<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/892801/\">庵野秀明</a>。对于曾经在<a target=\"_blank\" rel=\"nofollow\" href=\"http://movie.mtime.com/55470/\">《欢迎来到隔离病房》</a>、<a target=\"_blank\" rel=\"nofollow\" href=\"http://movie.mtime.com/50323/\">《无厘头森林的第一次接触》</a>等影片中有过真人表演经验，并在动画片《阿倍野桥魔法商店街》第12集客串过声优的庵野导演来说，担任长篇动画电影的主演声优还实属首次。世界瞩目的动画大师宫崎骏新作中有另一位重量级动画大师加盟声优，这种卡司可谓难得一见。庵野导演早在动画从业者时期就参与过<a target=\"_blank\" rel=\"nofollow\" href=\"http://movie.mtime.com/10606/\">《风之谷》</a>的设计工作中，一直视宫崎骏为自己的老师。</p><p>&nbsp;</p><p></p><p>　　此外，本片汇聚的声优卡司还有<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/1818454/\">泷本美织</a>（饰演菜穗子）、<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/920996/\">西岛秀俊</a>（饰演本庄季郎）、<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/936605/\">西村雅彦</a>（饰演黑川）、<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/952841/\">风间杜夫</a>（饰演里见）、<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/952845/\">竹下景子</a>（饰演二郎的母亲）、<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/1096466/\">志田未来</a>（饰演堀越加代）、<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/920900/\">国村隼</a>（饰演服部）、<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/940744/\">大竹忍</a>（饰演黑川夫人）、<a target=\"_blank\" rel=\"nofollow\" href=\"http://people.mtime.com/938426/\">野村万斋</a>（饰演卡普罗尼）。</p>\r\n\r\n<br></p>', '1524733982.jpg', '起风了 ');
INSERT INTO `news` VALUES ('7', '头号玩家', '2018-04-26', '<p>\r\n\r\n<p><img alt=\"\" src=\"http://cms-bucket.nosdn.127.net/7fc3d9d483fc46d4acc07a335ab50cd020180421131714.jpeg?imageView&amp;thumbnail=550x0\" title=\"Image: http://cms-bucket.nosdn.127.net/7fc3d9d483fc46d4acc07a335ab50cd020180421131714.jpeg?imageView&amp;thumbnail=550x0\"><br></p><p>《头号玩家》的宣传海报</p><p>这段时间，《头号玩家》意外成为一部爆款电影。说它“意外”，是因为这样一部充斥着游戏元素的电影不仅受到广泛关注，还赢得了大量好评。在豆瓣上，这部电影收获了33.6万次评价和8.9的高分，超过了豆瓣电影条目中98%的科幻片。</p><p>当然，观影者不可能都是游戏迷，《头号玩家》也不单纯是一部游戏迷的电影——它由华纳制作，斯皮尔伯格指导，面向大众，上映10天即在中国收获10亿票房……一切都表明：《头号玩家》里有着超乎游戏、彩蛋、怀旧的更多意义，值得深究。</p><p>《头号玩家》并不是第一部与游戏有关的电影，但该片的确进行了一些颇具新意的尝试。比如，《魔兽世界》是将已经存在的游戏内容转化为电影，其中的人物是虚拟世界的人物；《刺客信条》略有不同，同时出现了现实和虚拟两个空间，但重心仍然放在虚拟空间上；《头号玩家》则更进一步，其真正的价值在于预言性地为虚拟世界赋予了和现实世界同等重要的地位，并以此来探讨现实和虚拟的相互作用。这是十年前的电影不敢想的，虽然去年《星际特工》也在某些段落做出过结合两种空间的尝试，但《头号玩家》无疑是第一部敢于通篇尝试打通虚实两个世界的好莱坞大制作。</p><p>《头号玩家》与游戏IP电影的本质区别还在于，在该片中，游戏作为一种艺术形式的价值得到认可，并且开始直接和电影进行了两种艺术形式之间的平等嫁接。这种认可在片中的体现是多方面的，无论是空间构成、镜头运动，还是影片结构、角色设定，都是电影文化和游戏结合的产物，这与此前单纯的IP改编存在巨大差异。比如，三把钥匙的谜题设置，就既是典型的解谜游戏语汇，又有《公民凯恩》的痕迹。此外，由于大型游戏制作往往需要花费四五年甚至更长时间，这就使得游戏往往难以与现实世界直接对接，或多或少都会存在寓言性。因此，从某种程度上说，整部电影的乌托邦气质固然是由电影文本决定的，但也是影片的游戏属性所赋予的。</p><p>进一步看，影片的一个重要主题——不断强调要打破游戏“幻觉”，请求观众不要沉迷其中——贯穿始终，但游戏和电影却有着十分相似的幻觉机制，尤其是在3D电影时代，二者都旨在为受众营造巨大的虚构空间。电影依靠的是3D眼镜，游戏依靠的是VR。从这点看，《头号玩家》其实和米开朗琪罗·安东尼奥尼的经典作品《放大》有着异曲同工之妙，都有媒介自指的意味，而从这里我们大概也可以看出视觉艺术从摄影向电影再向游戏的嬗变。在这个序列当中，空间的连贯性在逐渐递增，直至形成一个完整的、全然沉浸的、身临其境的、令人无法自拔的体验。</p><p>可见，电影作为一门极具包容性的艺术门类，一直在吸纳其他艺术形式的精华，并为其所用，而从另一个角度看，这也对电影的制作者和研究者提出了更高要求。或许在不久的将来，如果不懂游戏，可能就会失去部分对电影的发言权。要知道，随着全球“80后”“90后”步入社会，那些曾经被定义甚至被贬低为“亚文化”的东西正在成为他们引以为豪且又频频引用的“圣经”。</p><p>斯皮尔伯格显然意识到了这个问题，作为可以驾驭亿万级制作的金牌导演和主流价值的捍卫者，他并没有一味地关注主流价值，拍摄更多的《拯救大兵瑞恩》或者《林肯》，而是自觉地将目光投向了全新的媒介和领域。事实上，不光是斯皮尔伯格，近两年美国主流电影工业正在一步步接纳并吸收亚文化元素，这种融合是通过两条路径实现的：一条是主流制作使用非主流话语，比如斯皮尔伯格拍《头号玩家》；另一条是非主流制作使用主流话语，比如B级片爱好者吉尔莫·德尔·托罗执导的《水形物语》。在两条路径的交汇之下，原先仅仅被一小部分人认可和欣赏的趣味变成了更加大众化的经典，也成为好莱坞电影内容的一大源泉。一个有趣的例证是，斯坦利·库布里克执导的恐怖片《闪灵》在《头号玩家》中被大量戏仿，而《闪灵》并非生来就是经典，它是在上世纪80年代获得大量恶评后才由亚文化路径进入主流视野的。</p><p>与此同时，对游戏有所了解的影迷会发现，片中很多元素并非来自于上世纪80年代，有些甚至出自于2000年之后的单机游戏，比如《生化奇兵》《无主之地》等，这表明《头号玩家》不仅是对上世纪末ACG(动画、漫画和游戏)亚文化的一次回顾，也是对整个数字时代以来网络文化的一次整合。</p><p>当然，这些所谓“彩蛋”的大规模爆发式出现也并非基于某种想象中的“彩蛋文化”，而是来自于切实的压力。事实上，片中对于上世纪50年代末电影衰落情况的种种描绘放到如今仍然成立，只不过那时对电影院线构成致命威胁的是电视，而现在则是电视和流媒体“联军”。如果放在这样一个背景下，《头号玩家》的诞生就更值得玩味。</p><p>斯皮尔伯格近日在接受媒体专访时表示，在流媒体平台上首播的电影，不应被允许参与奥斯卡的评选，“一旦电影变成了电视的规格，这就是电视电影，如果好看的话，可以得到艾美奖，但不应该得奥斯卡。”此番言论绝非偶然，从某种意义上讲，“院线+亚文化”与“电视+流媒体”的世纪之战即将展开。</p><p>《头号玩家》是一部恰好处在某个重要时间节点甚至时代十字路口上的片子。所以，不管我们称之为“后现代虚拟战争史诗”“80年代以来流行文化巡礼”，还是“近未来乌托邦科幻”，都不无道理，这些形容都可以为我们理解《头号玩家》的意义指出方向。而当我们剥开层层外壳后，影片的内核——对新技术背景下形成的权力关系的深切焦虑——也显露无遗。</p>\r\n\r\n<br></p>', '1524734574.jpg', '我要头盔我要头盔');
INSERT INTO `news` VALUES ('10', '歌姬', '2018-04-27', '<p>\r\n\r\n<div></div><div><h2><i></i>&nbsp;描述</h2><p>来源: © 百度云群组： 深海少女</p><p>我是受到百度云群组:深海少女  的群主：此生唯爱miku 的委托来分享已整理好的资源<br>资源贡献者:帝皇羽雪/残影/欲与莱士笑风生/wanan131690/旅行家77<br><strong>深海少女QQ群：512348719</strong><br>这份资源里也包含了2010到2017年初音未来的各种演唱会，而且不光有演唱会，还有一些其他的视频，比如说2014年大卫深夜秀里初音未来的短暂亮相<br>以下是合集资源的详情</p><p>初音未来鼓童特别演唱会<br>香蕉计划B.I.G嘉年华初音未来演唱会<br>初音未来2016上海演唱会<br>初音未来 SEGA Fes in Tokyo<br>初音未来2016魔法未来(含精华特典)<br>初音未来2016日本巡演<br>初音未来2015魔法未来(含精华特典)<br>初音未来2015上海演唱会<br>初音未来2014纽约演唱会<br>初音未来2014魔法未来(含精华特典)<br>初音未来2014印度尼西亚演唱会<br>初音未来2013非官方演唱会<br>初音未来2013魔法未来<br>初音未来2013关西演唱会(含精华特典)<br>初音未来2012台湾演唱会<br>初音未来2012感谢祭(含精华特典)<br>初音未来2011札幌演唱会(含精华特典)<br>初音未来2011洛杉矶演唱会<br>初音未来2011东京演唱会(含精华特典)<br>初音未来2010大感谢祭<br>大水怪-千本樱(2015音乐节).ts<br>『VOCALOID歌谣祭2013（春）～初音未来所衍生出的全新音乐世界～』[Vmoe字幕组].mp4<br>[大水怪] 初音ミク – ray .ts<br>[Vmoe字幕组]初音ミク 2014.39 MATSURI DA DIVA[LIVE部分].mp4<br>[Vmoe]「IA First Live Concert in Japan -PARTY A GO-GO-」[1080p].torrent<br>Vmoe 超party 2016 .mp4<br>Vmoe 超party 2015.mkv<br>Vmoe 超party 2014.mp4<br>Vmoe 超party 2013.mkv<br>Snow Fairy Story【SNOW MIKU 2015】.mp4<br>Sharing The World(2014大卫深夜秀) .mp4<br>BUMP OF CHICKEN feat. HATSUNE MIKU「ray」.mp4</p></div>\r\n\r\n<br></p>', '1524790943.jpg', '一起来聆听吧');

-- ----------------------------
-- Table structure for news_attribute
-- ----------------------------
DROP TABLE IF EXISTS `news_attribute`;
CREATE TABLE `news_attribute` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `news_id` int(11) NOT NULL COMMENT '资讯id',
  `like` int(11) NOT NULL DEFAULT '0' COMMENT '喜欢人数',
  `visitor` int(11) NOT NULL DEFAULT '0' COMMENT '访客人数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='资讯属性表';

-- ----------------------------
-- Records of news_attribute
-- ----------------------------
INSERT INTO `news_attribute` VALUES ('1', '3', '0', '1');
INSERT INTO `news_attribute` VALUES ('2', '4', '0', '27');
INSERT INTO `news_attribute` VALUES ('3', '5', '0', '0');
INSERT INTO `news_attribute` VALUES ('4', '7', '0', '4');
INSERT INTO `news_attribute` VALUES ('5', '10', '0', '17');
INSERT INTO `news_attribute` VALUES ('6', '6', '0', '22');

-- ----------------------------
-- Table structure for official_msg
-- ----------------------------
DROP TABLE IF EXISTS `official_msg`;
CREATE TABLE `official_msg` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL COMMENT '消息标题',
  `content` text NOT NULL COMMENT '内容',
  `date` datetime NOT NULL COMMENT '发布时间',
  `href` varchar(100) NOT NULL COMMENT '跳转路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='官方消息表';

-- ----------------------------
-- Records of official_msg
-- ----------------------------
INSERT INTO `official_msg` VALUES ('1', '2018年3月正式开启的项目', '电影网站开始正式施行中....', '2018-03-13 17:38:35', '1');
INSERT INTO `official_msg` VALUES ('2', '1周 计划启动', '电影网站的开发，已经过去一周了，项目的基本框架设计已经完成', '2018-03-20 17:39:31', '2');
INSERT INTO `official_msg` VALUES ('3', '2周 计划启动', '电影网站的初版页面已经完成ok', '2018-03-27 17:40:19', '3');
INSERT INTO `official_msg` VALUES ('4', '3周 计划启动', '电影网站的后台工程正式启动啦', '2018-04-03 17:40:51', '4');
INSERT INTO `official_msg` VALUES ('5', '4周 计划启动', '电影网站的后台功能已经完成了2/3了', '2018-04-10 17:41:25', '4');
INSERT INTO `official_msg` VALUES ('6', '5周 计划启动', '电影网站的个人模块即将实现了！！！！', '2018-04-17 17:41:57', '5');
INSERT INTO `official_msg` VALUES ('7', '6周 计划启动', '电影网站即将在下周四与大家见面了，movieEntral上线了!!!!', '2018-04-24 17:43:04', '6');
INSERT INTO `official_msg` VALUES ('8', '后台管理系统进度', '第7周了，后台管理系统的进度已经到了50%了，语句在今天能够完成项目，再有两天的时间收尾，就可以上线啦ヾ(๑╹◡╹)ﾉ\"', '2018-04-27 10:08:29', '1');
INSERT INTO `official_msg` VALUES ('9', '电影的上传功能实现', '更新了热血高校的资源，欢迎小伙伴前来观看', '2018-04-28 13:38:56', '1');
INSERT INTO `official_msg` VALUES ('10', '添加电影功能', '添加电影功能成功实现了，这意味着后台网页的全部功能已实现完成！', '2018-04-28 15:47:37', '1');
INSERT INTO `official_msg` VALUES ('11', '电影资源更新！', '增加了<<后来的我们>>的资源小伙伴们快来啊！', '2018-04-28 15:53:43', '1');
INSERT INTO `official_msg` VALUES ('12', '2018-4-28 19:09 这一刻', '电影网站的功能全部实现了，在五一之前完成，这一点真的很令我感动，算是给自己的学生时代划上一个完美的句号了！', '2018-04-28 19:11:11', '1');

-- ----------------------------
-- Table structure for per_to_film
-- ----------------------------
DROP TABLE IF EXISTS `per_to_film`;
CREATE TABLE `per_to_film` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `movie_id` int(11) NOT NULL COMMENT '电影id',
  `user_id` varchar(50) NOT NULL COMMENT '用户id',
  `score` int(11) DEFAULT '0' COMMENT '评分',
  `collect` enum('true','false') DEFAULT 'false' COMMENT '是否收藏',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8 COMMENT='个人对电影';

-- ----------------------------
-- Records of per_to_film
-- ----------------------------
INSERT INTO `per_to_film` VALUES ('4', '4', '141842291', '6', 'true');
INSERT INTO `per_to_film` VALUES ('5', '4', '141842261', '7', 'false');
INSERT INTO `per_to_film` VALUES ('6', '1', '141842261', '8', 'true');
INSERT INTO `per_to_film` VALUES ('7', '6', '141842261', '9', 'true');
INSERT INTO `per_to_film` VALUES ('8', '7', '141842261', '8', 'true');
INSERT INTO `per_to_film` VALUES ('10', '3', '141842261', '6', 'true');
INSERT INTO `per_to_film` VALUES ('11', '1', '141842291', '8', 'true');
INSERT INTO `per_to_film` VALUES ('12', '7', '141842291', '8', 'true');
INSERT INTO `per_to_film` VALUES ('13', '5', '141842291', '0', 'false');
INSERT INTO `per_to_film` VALUES ('14', '48', '141842291', '0', 'false');
INSERT INTO `per_to_film` VALUES ('15', '22', '141842261', '5', 'false');
INSERT INTO `per_to_film` VALUES ('16', '2', '141842291', '4', 'true');
INSERT INTO `per_to_film` VALUES ('17', '22', '141842291', '4', 'true');
INSERT INTO `per_to_film` VALUES ('18', '1', '958752538', '0', 'false');
INSERT INTO `per_to_film` VALUES ('19', '3', '958752538', '0', 'false');
INSERT INTO `per_to_film` VALUES ('20', '4', '958752538', '6', 'true');
INSERT INTO `per_to_film` VALUES ('21', '7', '958752538', '0', 'true');
INSERT INTO `per_to_film` VALUES ('22', '15', '958752538', '0', 'true');
INSERT INTO `per_to_film` VALUES ('23', '6', '958752538', '0', 'true');
INSERT INTO `per_to_film` VALUES ('24', '39', '141842291', '4', 'false');
INSERT INTO `per_to_film` VALUES ('25', '20', '141842291', '9', 'false');
INSERT INTO `per_to_film` VALUES ('26', '45', '141842291', '7', 'false');
INSERT INTO `per_to_film` VALUES ('27', '10', '141842291', '8', 'false');
INSERT INTO `per_to_film` VALUES ('28', '6', '141842291', '0', 'true');
INSERT INTO `per_to_film` VALUES ('29', '93', '958752538', '0', 'true');

-- ----------------------------
-- Table structure for per_to_review
-- ----------------------------
DROP TABLE IF EXISTS `per_to_review`;
CREATE TABLE `per_to_review` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `fr_id` int(11) NOT NULL COMMENT '影评id',
  `user_id` int(11) NOT NULL COMMENT '评论人id',
  `status` enum('true','false') DEFAULT NULL COMMENT '点击状态',
  `feel` enum('hate','like') DEFAULT NULL COMMENT '影评喜好',
  `date` datetime DEFAULT NULL COMMENT '点击时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='影评评论表';

-- ----------------------------
-- Records of per_to_review
-- ----------------------------
INSERT INTO `per_to_review` VALUES ('7', '14', '141842261', 'true', 'like', '2018-04-09 09:46:55');
INSERT INTO `per_to_review` VALUES ('10', '14', '141842291', 'true', 'like', '2018-04-09 10:26:26');
INSERT INTO `per_to_review` VALUES ('11', '13', '141842291', 'true', 'like', '2018-04-09 10:26:49');
INSERT INTO `per_to_review` VALUES ('12', '10', '141842291', 'true', 'hate', '2018-04-09 10:36:51');
INSERT INTO `per_to_review` VALUES ('13', '15', '958752538', 'true', 'like', '2018-04-20 14:42:58');
INSERT INTO `per_to_review` VALUES ('14', '6', '958752538', 'true', 'like', '2018-04-20 14:43:43');
INSERT INTO `per_to_review` VALUES ('15', '14', '958752538', 'true', 'like', '2018-04-24 14:04:29');
INSERT INTO `per_to_review` VALUES ('16', '16', '141842291', 'true', 'like', '2018-04-27 16:36:05');

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS `template`;
CREATE TABLE `template` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role` enum('official','persion') DEFAULT 'official' COMMENT '模板角色',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `content` text COMMENT '模板',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='模板表';

-- ----------------------------
-- Records of template
-- ----------------------------
INSERT INTO `template` VALUES ('1', 'official', '消息回复', '尊敬的会员你好，你的content消息已被查看！谢谢支持本网站');

-- ----------------------------
-- Table structure for type
-- ----------------------------
DROP TABLE IF EXISTS `type`;
CREATE TABLE `type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(11) NOT NULL COMMENT '类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='类似表';

-- ----------------------------
-- Records of type
-- ----------------------------
INSERT INTO `type` VALUES ('1', '爱情');
INSERT INTO `type` VALUES ('2', '动作');
INSERT INTO `type` VALUES ('3', '喜剧');
INSERT INTO `type` VALUES ('4', '科幻');
INSERT INTO `type` VALUES ('5', '灾难');
INSERT INTO `type` VALUES ('6', '悬疑');
INSERT INTO `type` VALUES ('7', '动画');
INSERT INTO `type` VALUES ('8', '剧情');
INSERT INTO `type` VALUES ('9', '记录');
INSERT INTO `type` VALUES ('10', '战争');
INSERT INTO `type` VALUES ('11', '运动');
INSERT INTO `type` VALUES ('12', '励志');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '不是uid 是注册顺序',
  `uid` varchar(50) NOT NULL COMMENT '账户',
  `birthday` date DEFAULT NULL COMMENT '出生日期',
  `person_sign` varchar(150) DEFAULT NULL COMMENT '个人签名',
  `person_notice` text COMMENT '个人公告',
  `name` varchar(50) DEFAULT NULL COMMENT '昵称',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `question` enum('name','hobby','lover') DEFAULT 'name' COMMENT '获取密码的问题',
  `answer` varchar(50) DEFAULT NULL COMMENT '密码的答案',
  `register_time` date DEFAULT NULL,
  `img_head` varchar(100) DEFAULT 'default.jpg' COMMENT '头像',
  `integral` int(11) DEFAULT NULL COMMENT '积分（暂定）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=141842269 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('141842261', '141842261', '2018-12-12', '黄海之水天上来！', '4897', 'momo桑', '123456', 'hobby', '美食', '2018-04-01', '141842261.jpg', null);
INSERT INTO `user` VALUES ('141842262', '141842291', '2018-12-12', '与光同行', '最近好忙呀！', 'momo', '7474740', 'hobby', '钟贵峰', '2018-04-02', '141842291.png', null);
INSERT INTO `user` VALUES ('141842263', '141842260', null, null, null, null, '123456', 'lover', '我自己', '2018-04-05', null, null);
INSERT INTO `user` VALUES ('141842264', '141842299', null, null, null, null, '123456', 'name', '龙傲天', '2018-04-14', null, null);
INSERT INTO `user` VALUES ('141842265', '15556695185', null, null, null, null, 'a6361666', 'lover', 'xtt', '2018-04-15', null, null);
INSERT INTO `user` VALUES ('141842266', '111111', null, null, null, null, '111111', 'hobby', '111111', '2018-04-15', null, null);
INSERT INTO `user` VALUES ('141842267', 'lilyzou', null, null, null, null, 'lilyzou520', 'name', '邹彦芳', '2018-04-23', 'default.jpg', null);
INSERT INTO `user` VALUES ('141842268', '958752538', '1996-02-12', 'bug', '我是一个小萌新哈哈\n', '我就是我', 'a6361666', 'lover', 'myself', '2018-04-24', '958752538.png', null);
