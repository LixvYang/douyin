CREATE DATABASE IF NOT EXISTS `douyin`;

use `douyin`;

DROP TABLE IF EXISTS `video`;

CREATE TABLE `video` (
	`id` int UNSIGNED NOT NULL AUTO_INCREMENT,
	`author` int UNSIGNED NOT NULL COMMENT '作者id',
	`title` varchar(255) not null COMMENT '视频标题',
	`play_url` varchar(255) not null COMMENT '播放地址',
	`cover_url` varchar(255) not null COMMENT '封面地址',
	`favorite_count` int not null default 0 COMMENT '喜欢数',
	`comment_count` int not null default 0 COMMENT '评论数',
	-- `is_favorite` tinyint not null default 0 COMMENT '是否喜欢 0:否 1:是',
	`create_time` int null default null COMMENT '创建时间',
	primary key (`id`)
) ENGINE = InnoDB default CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '视频列表' ROW_FORMAT = DYNAMIC;

INSERT INTO `video` VALUES (1, 1, 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 1, 0, 1257894000);
INSERT INTO `video` VALUES (2, 1, 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 1, 0, 1257894000);

DROP TABLE IF EXISTS `user`;
create table `user` (
	`id` int UNSIGNED not null auto_increment COMMENT '用户ID',
	`username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci not null COMMENT '用户名称',
	`password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci not null COMMENT '用户密码',
	`follower_count` int not null default '0' COMMENT '用户的粉丝总数',
	`follow_count` int not null default '0' COMMENT '用户的关注数量',
	primary key (`id`),
	UNIQUE INDEX `id` (`id`) USING BTREE,
	UNIQUE INDEX `username` (`username`) USING BTREE
) ENGINE = InnoDB default CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表';

INSERT INTO `user` VALUES (1, '罹心', '123456', 1, 1);
INSERT INTO `user` VALUES (2, '无虞', '654321', 1, 1);


DROP TABLE IF EXISTS `user_follow`;

create table `user_follow` (
	`id` int UNSIGNED not null auto_increment COMMENT '自增ID',
	`user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
	`follow_id` int UNSIGNED not null default '0' COMMENT '关注的用户ID',
	`is_follow` tinyint(1) default '0' COMMENT '是否关注 1:关注 0:取消关注',
	`create_time` varchar(20) default null COMMENT '创建时间',
	-- `update_time` timestamp default null COMMENT '更新时间',
	primary key(`id`),
	UNIQUE INDEX `user_follow_id` (`user_id`, `follow_id`) USING BTREE,
	FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY (`follow_id`) REFERENCES `user`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB default CHARACTER SET = utf8mb4 COMMENT = '用户关注表';

INSERT INTO `user_follow` VALUES (1, 1, 2, 0, '2022-05-11 08:10:02');
INSERT INTO `user_follow` VALUES (2, 1, 3, 0, '2022-05-11 08:10:02');
INSERT INTO `user_follow` VALUES (3, 1, 4, 0, '2022-05-11 08:10:02');
INSERT INTO `user_follow` VALUES (4, 1, 5, 0, '2022-05-11 08:10:02');
INSERT INTO `user_follow` VALUES (5, 3, 1, 0, '2022-05-11 08:10:02');

-- 查看用户1的粉丝列表
-- select user_id from user_follow where follow_id = 1;
-- 查找用户1的关注列表
-- select follow_id from user_follow where user_id = 1;
-- 查看别人的粉丝或者关注列表
DROP TABLE IF EXISTS `comment`;
create table `comment` (
	`id` bigint not null auto_increment COMMENT '自增ID',
	-- `father_comment_id` bigint null default null COMMENT '父评论ID',
	-- `to_user_id` int UNSIGNED NULL default NULL COMMENT '被评论者ID',
	`video_id` int UNSIGNED not null COMMENT '视频ID',
	`from_user_id` int not null COMMENT '留言者ID',
	`comment` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci not null COMMENT '评论内容',
	`create_date` varchar(20) null default null COMMENT '评论时间',
	-- `delete_date` timestamp null default null COMMENT '删除时间',
	primary key (`id`)
) ENGINE = InnoDB default CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '评论列表' ROW_FORMAT = DYNAMIC;
INSERT INTO `comment` VALUES (1, 1, 1, '熊大熊大，俺要吃蜂蜜', '05-11');



DROP TABLE IF EXISTS `users_like_video`;
create table `users_like_video` (
	`id` int NOT NULL AUTO_INCREMENT,
	`user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
	`video_id` int UNSIGNED NOT NULL COMMENT '视频ID',
	`is_like` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否点赞 0:取消点赞 1:点赞',
	primary key (`id`),
	UNIQUE INDEX `user_video_rel` (`user_id`, `video_id`) USING BTREE
) ENGINE = InnoDB default CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户喜欢的视频' ROW_FORMAT = DYNAMIC;
INSERT INTO `users_like_video` VALUES (1, 1, 1, 1);
INSERT INTO `users_like_video` VALUES (2, 3, 1, 1);

-- 添加评论video_id 对应video_id的外键索引
-- alter table comment add FOREIGN KEY comment_video_fk_1 (video_id) REFERENCES video (id);
-- 添加用户喜欢的外键索引
-- alter table users_like_video add FOREIGN KEY userslike_user_fk_id (user_id) REFERENCES user (id);
-- alter table users_like_video add FOREIGN KEY userslike_video_fk_id (video_id) REFERENCES video (id);
-- 添加用户关注索引
-- 在表里建立
-- alter table user_follow add FOREIGN KEY userfollow_user_fk_id_1 (user_id) REFERENCES user (id);
-- alter table user_follow add FOREIGN KEY userfollow_user_fk_id2 (follow_id) REFERENCES user (id);

-- video 到user的外键索引
-- alter table video add FOREIGN KEY video_user_fk_id (author) REFERENCES user (id);
