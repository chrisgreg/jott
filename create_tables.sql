CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL ,
  `pass` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE(`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1

CREATE TABLE `blogs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `title` VARCHAR(50) NOT NULL,
  `subtitle` VARCHAR(50) NOT NULL,
  `created` DATETIME null,
  `private` BOOLEAN NOT NULL,
  `read_count` int(10) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `tb_fk` (`user_id`),
  CONSTRAINT `tb_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1

CREATE TABLE `jotts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `blog_id` int(11) NOT NULL,
  `content` VARCHAR(3000) NOT NULL,
  `created` DATETIME null,
  PRIMARY KEY (`id`),
  KEY `jott_fk` (`blog_id`),
  KEY `ujott_fk` (`user_id`),
  CONSTRAINT `jott_fk` FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`),
  CONSTRAINT `ujott_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) 

 --GET Blogs with Jotts & Content & Authors
--  SELECT
--  b.title,
--  b.subtitle,
--  b.read_count,
--  b.created blog_created,
--  ub.name blog_author,
--  j.created jott_created,
--  uj.name jott_author,
--  j.content
--  from blogs b
--  INNER JOIN jotts j on b.id=j.blog_id
--  INNER JOIN users uj on j.user_id=uj.id
--  INNER JOIN users ub on ub.id=b.id
--  ORDER BY jott_created