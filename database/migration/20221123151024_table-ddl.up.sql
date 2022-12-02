CREATE TABLE roles(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name LONGTEXT,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	PRIMARY KEY (id),
	KEY (deleted_at)
);

CREATE TABLE users(
	id VARCHAR(36) NOT NULL,
	name LONGTEXT,
	email LONGTEXT,
	password LONGTEXT,
	mobile_number LONGTEXT,
	is_active TINYINT(1),
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	role_id BIGINT UNSIGNED,
	PRIMARY KEY (id),
	KEY (deleted_at,role_id),
	FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE CASCADE ON DELETE CASCADE
);


CREATE TABLE user_coins(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	user_id varchar(36),
	amount BIGINT UNSIGNED,
	PRIMARY KEY (id),
	KEY (user_id),
	FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE	
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
	required_point BIGINT UNSIGNED,
	valid_until DATETIME(3),
	category_id BIGINT UNSIGNED,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	PRIMARY KEY (id),
	KEY (deleted_at,category_id),
	FOREIGN KEY (category_id) REFERENCES categories(id)
	ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE products(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name LONGTEXT,
	category_id BIGINT UNSIGNED,
	minimum_transaction INT UNSIGNED,
	points BIGINT,
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3),
	PRIMARY KEY (id),
	KEY (deleted_at,category_id),
	FOREIGN KEY (category_id) REFERENCES categories(id)
	ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE otps(
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	otp_code LONGTEXT,
	email LONGTEXT,
	created_at DATETIME(3),
	PRIMARY KEY (id)
);