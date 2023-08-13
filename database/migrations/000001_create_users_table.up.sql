CREATE TABLE `users` (
    `id` BIGINT unsigned PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `created_at` DATETIME(3),
    `updated_at` DATETIME(3),
    `deleted_at` DATETIME(3)
);