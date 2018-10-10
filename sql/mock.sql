use renrenbuy;

insert into classes (name, description) values ("水果", "这是水果");
insert into classes (name, description) values ("蔬菜", "这是蔬菜");

insert into subclasses (class_id, name, description) values (1, "时令水果", "basic");
insert into subclasses (class_id, name, description) values (1, "进口水果", "advanced");
insert into subclasses (class_id, name, description) values (2, "时令蔬菜", "basic");
insert into subclasses (class_id, name, description) values (2, "精品蔬菜", "advanced");

insert into products (subclass_id, name, description) values (1, "苹果", "pingguo");
insert into products (subclass_id, name, description) values (2, "橘子", "juzi");
insert into products (subclass_id, name, description) values (3, "白菜", "baicai");
insert into products (subclass_id, name, description) values (4, "萝卜", "luobu");

insert into product_properties (product_id, display_url) values (1, "/display/product/1.jpg");
insert into product_properties (product_id, display_url) values (2, "/display/product/2.jpg");
insert into product_properties (product_id, display_url) values (3, "/display/product/3.jpg");
insert into product_properties (product_id, display_url) values (4, "/display/product/4.jpg");

insert into regions (name, longitude, latitude) values ("北京", 30, 50);

insert into shops (region_id, name, address, description, display_url) values (1, "第一超市", "北京", "只是一家超市", "/display/shop/1.jpg")

insert into sales (product_id, shop_id, price, title, description, create_time) values (1, 1, 10, "便宜苹果", "这是最便宜的苹果", 1539180514);
insert into sales (product_id, shop_id, price, title, description, create_time) values (2, 1, 10, "便宜橘子", "这是最便宜的橘子", 1539180514);
insert into sales (product_id, shop_id, price, title, description, create_time) values (3, 1, 10, "便宜白菜", "这是最便宜的白菜", 1539180514);
insert into sales (product_id, shop_id, price, title, description, create_time) values (4, 1, 10, "便宜萝卜", "这是最便宜的萝卜", 1539180514);

