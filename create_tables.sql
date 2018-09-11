CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL ,
  `pass` varchar(100) NOT NULL,
  `github_profile` varchar(80) NOT NULL,
  `twitter_profile`  varchar(80) NOT NULL,
  `facebook_profile`  varchar(80) NOT NULL,
  `website`  varchar(80) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE(`email`),
  UNIQUE(`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1

CREATE TABLE `blogs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `title` VARCHAR(50) NOT NULL,
  `subtitle` VARCHAR(50) NOT NULL,
  `created` DATETIME null,
  `private` BOOLEAN NOT NULL,
  `read_count` int(10) NOT NULL DEFAULT 0,
  `theme` ENUM('black-white', 'jott-yellow', 'hot-coral', 'teal-orange', 'blue-lagoon', 'forest-green', 'calm-pink', 'beige'),
  PRIMARY KEY (`id`),
  KEY `tb_fk` (`user_id`),
  CONSTRAINT `tb_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1