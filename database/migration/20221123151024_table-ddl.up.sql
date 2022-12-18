CREATE TABLE roles(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name LONGTEXT,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	PRIMARY KEY (id),
	KEY (deleted_at)
);

CREATE TABLE user_coins(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	amount BIGINT,
	PRIMARY KEY (id)
);

CREATE TABLE credits(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	amount DOUBLE,
	PRIMARY KEY (id)
);

CREATE TABLE users(
	id VARCHAR(36) NOT NULL,
	name LONGTEXT,
	email LONGTEXT,
	password LONGTEXT,
	mobile_number LONGTEXT,
	user_coin_id BIGINT UNSIGNED,
	credit_id BIGINT UNSIGNED,
	is_active TINYINT(1),
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	role_id BIGINT UNSIGNED,
	PRIMARY KEY (id),
	KEY (deleted_at,role_id,user_coin_id,credit_id),
	FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE CASCADE ON DELETE RESTRICT,
	FOREIGN KEY (user_coin_id) REFERENCES user_coins(id) ON UPDATE CASCADE ON DELETE RESTRICT,
	FOREIGN KEY (credit_id) REFERENCES credits(id) ON UPDATE CASCADE ON DELETE RESTRICT
);


CREATE TABLE categories(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name LONGTEXT,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	PRIMARY KEY (id),
	KEY (deleted_at)
);

CREATE TABLE rewards(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name LONGTEXT,
	description LONGTEXT,
	required_point BIGINT,
	valid_until DATETIME(3),
	category_id BIGINT UNSIGNED,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	PRIMARY KEY (id),
	KEY (deleted_at,category_id),
	FOREIGN KEY (category_id) REFERENCES categories(id)
	ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE products(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name LONGTEXT,
	description LONGTEXT,
	provider TEXT,
	active_period BIGINT,
	price DOUBLE,
	category_id BIGINT UNSIGNED,
	minimum_transaction BIGINT,
	coins BIGINT,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	PRIMARY KEY (id),
	KEY (deleted_at,category_id),
	FOREIGN KEY (category_id) REFERENCES categories(id)
	ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE otps(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	otp_code LONGTEXT,
	email LONGTEXT,
	created_at DATETIME(3),
	PRIMARY KEY (id)
);

CREATE TABLE redeems(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	reward_id BIGINT UNSIGNED,
	user_id VARCHAR(36),
	point_spent BIGINT,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	PRIMARY KEY (id),
	KEY (deleted_at,user_id,reward_id),
	FOREIGN KEY (user_id) REFERENCES users(id)
	ON UPDATE CASCADE ON DELETE SET NULL,
	FOREIGN KEY (reward_id) REFERENCES rewards(id)
	ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE forgot_passwords(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	email TEXT,
	token LONGTEXT,
	expired_at DATETIME,
	PRIMARY KEY (id)
);

CREATE TABLE faqs(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	question LONGTEXT,
	answer LONGTEXT,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	KEY (deleted_at),
	PRIMARY KEY (id)
);

CREATE TABLE transactions(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	user_id VARCHAR(36),
	status TEXT NOT NULL,
	amount DOUBLE,
	coins_earned BIGINT,
	product_id BIGINT UNSIGNED,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	KEY (user_id, product_id, deleted_at),
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id)
	ON UPDATE CASCADE ON DELETE SET NULL,
	FOREIGN KEY (product_id) REFERENCES products(id)
	ON UPDATE CASCADE ON DELETE SET NULL
);

-- CREATE TABLE payment_invoices(
-- 	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
-- 	transaction_id BIGINT UNSIGNED,
-- 	url TEXT,
-- 	amount BIGINT,
-- 	created_at DATETIME(3),
-- 	updated_at DATETIME(3),
-- 	deleted_at DATETIME(3),
-- 	KEY (deleted_at),
-- 	PRIMARY KEY (id)
-- );