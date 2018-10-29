USE renrenbuy;
CREATE TABLE clients (
 id INT AUTO_INCREMENT,
 open_id INT NOT NULL,
 name VARCHAR(32) NOT NULL,
 sex BOOLEAN NOT NULL,
 phone INT NOT NULL,
 PRIMARY KEY (id)
);
CREATE INDEX clients_open_id_index ON clients (open_id);

CREATE TABLE sellers (
 id INT AUTO_INCREMENT,
 name VARCHAR(32) NOT NULL,
 open_id INT NOT NULL,
 phone INT NOT NULL,
 address VARCHAR(128) NOT NULL,
 PRIMARY KEY (id)
);
CREATE INDEX sellers_open_id_index ON sellers (open_id);

CREATE TABLE orders (
 id INT AUTO_INCREMENT,
 client_id INT NOT NULL,
 sale_id INT NOT NULL,
 num INT NOT NULL,
 create_time INT NOT NULL,
 PRIMARY KEY (id)
);
CREATE INDEX order_client_id_index ON orders (client_id);
