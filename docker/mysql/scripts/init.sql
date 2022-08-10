CREATE DATABASE IF NOT EXISTS `golang_mvc_rest_api` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'golang_mvc_rest_api'@'%' IDENTIFIED WITH mysql_native_password BY 'golang_mvc_rest_api';
GRANT ALL ON *.* TO 'golang_mvc_rest_api'@'%';

USE golang_mvc_rest_api;