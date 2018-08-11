CREATE TABLE `User` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `pass` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1

CREATE TABLE `Blog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `title` VARCHAR(50) NOT NULL,
  `subtitle` VARCHAR(50) NOT NULL,
  `created` DATETIME null,
  PRIMARY KEY (`id`),
  KEY `tb_fk` (`user_id`),
  CONSTRAINT `tb_fk` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1

CREATE TABLE `Jotts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `blog_id` int(11) NOT NULL,
  `content` VARCHAR(3000) NOT NULL,
  `created` DATETIME null,
  PRIMARY KEY (`id`),
  KEY `jott_fk` (`blog_id`),
  CONSTRAINT `jott_fk` FOREIGN KEY (`blog_id`) REFERENCES `Blog` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1

