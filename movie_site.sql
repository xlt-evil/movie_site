-- 表1
CREATE TABLE `movie`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`movie_name` varchar(50)  DEFAULT NULL COMMENT '电影名字',
	`area` int(11)   DEFAULT NULL COMMENT '地区',
	`language` varchar(50) DEFAULT NULL COMMENT '语言',
	`type` int(11) DEFAULT NULL COMMENT '电影类型',
	`director` varchar(50) DEFAULT NULL COMMENT '导演',
	`release_date` date DEFAULT NULL COMMENT '上映日期',
	`update_date` date DEFAULT NULL COMMENT '更新日期',
	`introuduce` varchar(50) DEFAULT NULL COMMENT '电影简介',
	`year` varchar(50) DEFAULT NULL COMMENT '年份',
	`plot` text(1024) DEFAULT NULL COMMENT '剧情',
	`source` enum('exist','disappear') DEFAULT NULL COMMENT '是否有片源 1 有 0 无',
	`source_path` varchar(255) DEFAULT NULL COMMENT '资源路径',
	`source_from` int(11) DEFAULT NULL COMMENT '资源出处 1官方0个人 （暂定）',
	`integral` int(11) DEFAULT NULL COMMENT '积分（暂定）',
	`movie_img` varchar(50) DEFAULT 'default.jpg' COMMENT '电影图片'
	PRIMARY KEY (`id`)
)ENGINE=INNODB AUTO_INCREMENT=8 DEFAULT CHARSET = utf8 COMMENT='电影表'


-- 表2
CREATE TABLE `movie_attribute` (
	`id` INT (11) NOT NULL AUTO_INCREMENT,
	`movie_id` INT (11) NOT NULL COMMENT '电影id',
	`popularity` INT (11) DEFAULT 0 COMMENT '人气',
	`collect_num` INT (11) DEFAULT 0 COMMENT '收藏人数',
	`score`  DECIMAL(2,1) DEFAULT 0 COMMENT '评分',
	`year` VARCHAR (50) DEFAULT NULL COMMENT '年份',
	PRIMARY KEY (`id`)
)ENGINE=INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '电影属性表'

-- 3
CREATE TABLE `per_to_film`(
	`id` int(11) NOT NULL AUTO_INCREMENT ,
	`movie_id` int(11) NOT NULL COMMENT '电影id',
	`user_id` varchar(50) NOT NULL COMMENT '用户id',
	`score` int(11) DEFAULT 0 COMMENT '评分' ,
	`collect` enum('true','false') DEFAULT 'false' COMMENT '是否收藏',
	PRIMARY KEY(`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET=utf8 COMMENT = '个人对电影'

-- 表四
CREATE TABLE `movie_commentary`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`movie_id` int(11) NOT NULL COMMENT '电影id',
	`user_id` varchar(50) NOT NULL COMMENT '用户id',
	`date` datetime NOT NULL COMMENT '短评时间',
	`content` varchar(100) NOT NULL COMMENT '评论内容',
	PRIMARY KEY(`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '电影短评表'

-- 表五
CREATE TABLE `film_review`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`movie_id` int(11) NOT NULL COMMENT  '电影id',
	`user_id` varchar(50) NOT NULL COMMENT  '用户id',
	`title` varchar(100) NOT NULL COMMENT '影评标题',
	`content` text(2048) NOT NULL COMMENT  '内容',
	`date` datetime NOT NULL COMMENT  '发布时间',
	PRIMARY KEY (`id`)
)ENGINE= INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '电影影评'

-- 表6
CREATE TABLE `fr_attribute`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`fr_id` int(11) NOT NULL COMMENT '影评id',
	`movie_id` int(11) NOT NULL COMMENT '电影id',
	`like` int(11) DEFAULT 0 COMMENT '喜欢人数',
	`hate` int(11) DEFAULT 0 COMMENT '讨厌人数',
	`talk_num` int(11) DEFAULT 0 COMMENT '谈论次数',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '影评属性表'

-- 表7
CREATE TABLE `fr_talk`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`fr_id` int(11) NOT NULL COMMENT '影评id',
	`movie_id` int(11) NOT NULL COMMENT '电影id',
	`user_id` varchar(50) NOT NULL COMMENT '评论人id',
	`content` text(1024) COMMENT '评论内容',
	`date`  datetime COMMENT '发布时间',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '影评评论表'

-- 表8
CREATE TABLE `per_to_review`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`fr_id` int(11) NOT NULL COMMENT '影评id',
	`user_id` varchar(50) NOT NULL COMMENT '评论人id',
	`status` enum('true','false') COMMENT '点击状态',
	`feel`	enum('hate','like') COMMENT '影评喜好',
	`date`  datetime COMMENT '点击时间',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '影评喜好表'

-- 表9
CREATE TABLE `news`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`title` varchar(100) NOT NULL COMMENT '新闻标题',
	`date` datetime COMMENT '发布日期',
	`content` text(4096) COMMENT '内容',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '资讯表'

-- 表10
CREATE TABLE `news_attribute`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`news_id` int(11) NOT NULL COMMENT '资讯id',
	`like` int(11) NOT NULL DEFAULT 0 COMMENT '喜欢人数',
	`visitor` int(11) NOT NULL DEFAULT 0 COMMENT '访客人数',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '资讯属性表'

-- 表11
CREATE TABLE `official_msg`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`title` varchar(100) NOT NULL COMMENT '消息标题',
	`content` text(1024) NOT NULL COMMENT '内容',
	`date` datetime NOT NULL COMMENT '发布时间',
	`href` varchar(100) NOT NULL COMMENT '跳转路径',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '官方消息表'

-- 表12
CREATE TABLE `template`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`role` enum('official','persion') DEFAULT 'official ' COMMENT '模板角色',
	`title` varchar(50) NOT NULL COMMENT '标题',
	`content` text(1024) COMMENT '模板',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '模板表'

-- 表13
CREATE TABLE `msg_record`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`user_id` varchar(50) NOT NULL COMMENT '用户id',
	`title` varchar(50) NOT NULL COMMENT '标题',
	`content` text(1024) COMMENT '内容',
	`status` enum('false','true') DEFAULT 'false' COMMENT '查看状态',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '消息记录表'

-- 表14
CREATE TABLE `user`(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '不是uid 是注册顺序',
	`user_id` varchar(50) NOT NULL COMMENT '账户',
	`birthday` date COMMENT '出生日期',
	`person_sign` varchar(150) COMMENT '个人签名',
	`person_notice` text(1024) COMMENT '个人公告',
	`name` varchar(50) DEFAULT NULL COMMENT '昵称',
	`password` varchar(50) NOT NULL COMMENT '密码',
	`question` enum('name','hobby','lover') DEFAULT 'name' COMMENT '获取密码的问题',
	`answer` varchar(50) COMMENT '密码的答案',
	`regsiter_time` date COMMENT '注册时间',
	`img_head` varchar(100) DEFAULT 'default.jpg' COMMENT '头像',
	`integral` int(11) DEFAULT NULL COMMENT '积分（暂定）',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 141842261 CHARSET = utf8 COMMENT = '用户表'

-- 表15
CREATE TABLE `person_space`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`user_id` varchar(50) NOT NULL COMMENT '账户',
	`visitor` int(11) DEFAULT 0 COMMENT '访问次数',
	`fun` int(11) DEFAULT 0 COMMENT '粉丝（暂定）',
	`follow` int(11) DEFAULT 0 COMMENT '关注（暂定）',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 141842261 CHARSET = utf8 COMMENT = '个人空间表'

-- 表16
CREATE TABLE `history_record`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`user_id` varchar(50) int(11) NOT NULL COMMENT '账户',
	`movie_id` int(11) NOT NULL COMMENT '电影id',
	`date` datetime COMMENT '观看时间',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '看电影历史记录表'

-- 表17
CREATE TABLE `advise`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`user_id` varchar(50) int(11) NOT NULL COMMENT '账户',
	`role` enum('resource','build') DEFAULT 'resource' COMMENT '建议类型',
	`title` varchar(50) COMMENT '标题',
	`content`	text(1024) COMMENT '内容',
	`status` enum('wait','pass','fair')  COMMENT '建议状态',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 8 CHARSET = utf8 COMMENT = '建议表'

-- 表18
CREATE TABLE `language`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(11) NOT NULL COMMENT '地区',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 1 CHARSET = utf8 COMMENT = '地区表'

-- 表19
CREATE TABLE `language`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(11) NOT NULL COMMENT '语言',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 1 CHARSET = utf8 COMMENT = '语言表'

-- 表20
CREATE TABLE `type`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(11) NOT NULL COMMENT '类型',
	PRIMARY KEY (`id`)
)ENGINE = INNODB AUTO_INCREMENT = 1 CHARSET = utf8 COMMENT = '类型表'




-- 插入数据
insert INTO area (name) values('大陆')
insert INTO area (name) values('香港')
insert INTO area (name) values('台湾')
insert INTO area (name) values('日本')
insert INTO area (name) values('韩国')
insert INTO area (name) values('英国')
insert INTO area (name) values('西班牙')
insert INTO area (name) values('法国')
insert INTO area (name) values('美国')
insert INTO area (name) values('德国')
insert INTO area (name) values('印度')
insert INTO area (name) values('其他')

-- 
insert INTO language (name) values('国语')
insert INTO language (name) values('港语')
insert INTO language (name) values('英语')
insert INTO language (name) values('韩语')
insert INTO language (name) values('日语')
insert INTO language (name) values('西班牙语')
insert INTO language (name) values('德语')
insert INTO language (name) values('法语')
insert INTO language (name) values('其他')

-- 
insert INTO type (name) values('爱情')
insert INTO type (name) values('动作')
insert INTO type (name) values('喜剧')
insert INTO type (name) values('科幻')
insert INTO type (name) values('灾难')
insert INTO type (name) values('悬疑')
insert INTO type (name) values('动画')
insert INTO type (name) values('剧情')
insert INTO type (name) values('记录')
insert INTO type (name) values('战争')
insert INTO type (name) values('运动')
insert INTO type (name) values('励志')
