CREATE DATABASE renrenbuy character set utf8;
GRANT ALL ON renrenbuy.* TO 'renrenbuytest'@'1.1.1.1' IDENTIFIED BY 'somepassword';
FLUSH PRIVILEGES;
USE renrenbuy;
CREATE TABLE products (
 id INT AUTO_INCREMENT,
 subclass_id INT NOT NULL,
 name VARCHAR(32) NOT NULL,
 description VARCHAR(128) NOT NULL,
 PRIMARY KEY (id)
);
CREATE INDEX products_subclass_id_index ON products (subclass_id);


CREATE TABLE classes (
 id INT AUTO_INCREMENT,
 name VARCHAR(32) NOT NULL,
 description VARCHAR(128) NOT NULL,
 PRIMARY KEY (id)
);

CREATE TABLE subclasses (
 id INT AUTO_INCREMENT,
 class_id INT NOT NULL,
 name VARCHAR(32) NOT NULL,
 description VARCHAR(128) NOT NULL,
 PRIMARY KEY (id)
);
CREATE INDEX subclasses_class_id_index ON subclasses (class_id);

CREATE TABLE product_properties (
 product_id INT NOT NULL,
 display_url VARCHAR(128) NOT NULL,
 PRIMARY KEY (product_id)
);

CREATE TABLE regions (
 id INT AUTO_INCREMENT,
 name VARCHAR(32) NOT NULL,
 longitude INT NOT NULL,
 latitude INT NOT NULL,
 PRIMARY KEY (id)
);

CREATE TABLE shops (
 id INT AUTO_INCREMENT,
 region_id INT NOT NULL,
 name VARCHAR(32) NOT NULL,
 address VARCHAR(128) NOT NULL,
 description VARCHAR(256) NOT NULL,
 display_url VARCHAE(128) NOT NULL,
 PRIMARY KEY (id)
);
CREATE INDEX shops_region_id_index ON shops (region_id);

CREATE TABLE sales (
 id INT AUTO_INCREMENT,
 product_id INT NOT NULL,
 shop_id INT NOT NULL,
 price INT NOT NULL,
 title VARCHAR(32),
 description VARCHAR(128),
 create_time INT NOT NULL,
 PRIMARY KEY (id)
);
CREATE INDEX sales_product_id_index ON sales (product_id);
CREATE INDEX sales_shop_id_index ON sales (shop_id);
