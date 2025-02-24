# 数据库设计

[[toc]]


## 商品数据库


### 商品信息

```sql
CREATE TABLE product (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	puid varchar(16)  NOT NULL COMMENT '商品ID',
	title VARCHAR(255) NOT NULL COMMENT '商品名称',
	price DECIMAL(12,2) UNSIGNED NOT NULL default 0.00 COMMENT '商品价格',
	price_off decimal(3,2) UNSIGNED NOT NULL DEFAULT 1.00 COMMENT '商品折扣',
	desction text COMMENT '商品描述',
	main_img VARCHAR(255) NOT NULL COMMENT '商品图片',
	stock INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品库存',
	sales INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品销量',
	category_id bigint UNSIGNED NOT NULL COMMENT '商品分类ID',
	status tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品状态 0:未上架, 1: 上架',
	delete tinyint unsigned NOT NULL DEFAULT 0 COMMENT '商品删除状态 0:未删除, 1: 已删除',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY pk_product (id),
	KEY idex_uuid (puid),
	KEY idx_status (status),
	KEY idx_title (title)
) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT '商品信息表';

```




### 商品分类


```sql
CREATE TABLE product_category (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '分类ID',
	category_name VARCHAR(255) NOT NULL COMMENT '分类名称',
	category_parent_id bigint UNSIGNED NOT NULL default 0 COMMENT '父级分类ID',
	category_level tinyint UNSIGNED NOT NULL COMMENT '分类级别',
	category_sort tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '分类排序',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY pk_product_category (id)
)ENGINE=innodb COMMENT '商品分类表';

```



### 商品详情


```sql
CREATE TABLE product_detail (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	product_id bigint UNSIGNED NOT NULL COMMENT '商品ID',
	product_category_id bigint UNSIGNED NOT NULL COMMENT '商品分类ID',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY pk_product (id, product_id)
)ENGINE=innodb COMMENT '商品详情信息表';

```


### 商品图片

```sql
CREATE TABLE product_img (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	product_id varchar(16)  NOT NULL COMMENT '商品ID',
	product_imgurl VARCHAR(255) NOT NULL COMMENT '商品图片地址',
	product_mainurl VARCHAR(255) NOT NULL COMMENT '商品主图地址',
	product_order int UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片排序',
	product_type tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '图片类型 1:商品展示图, 2:商品详情图, 3: 商品轮播图, 9: 其他图片',
	product_status tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '商品状态 0:删除, 1: 上架',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY pk_product (id),
	KEY idex_product_id (product_id),
	KEY idx_product_status (product_status),
	KEY idx_product_main_img (product_main_img),
	KEY idx_product_other_img (product_other_img),
	key idx_product_order (product_order),
) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT '商品图片';
```


### 商品SKU


```sql
CREATE TABLE product_sku (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	product_name VARCHAR(255) NOT NULL COMMENT '商品名称',
	product_price DECIMAL(12,2) UNSIGNED NOT NULL COMMENT '商品价格',
	product_off decimal(3,2) UNSIGNED NOT NULL DEFAULT 1.00 COMMENT '商品折扣',
	product_status tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '商品状态',
	product_desc text COMMENT '商品描述',
	product_img VARCHAR(250) NOT NULL COMMENT '商品图片',
	product_stock INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品库存',
	product_sales INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品销量',
	product_category_id bigint UNSIGNED NOT NULL COMMENT '商品分类ID',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY pk_product (id)
)ENGINE=innodb COMMENT 'SKU信息表';

```

### 商品供应商


```sql
CREATE TABLE product_supplier (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	supplier_name VARCHAR(255) NOT NULL COMMENT '供应商名称',
	supplier_phone VARCHAR(255) NOT NULL COMMENT '供应商电话',
	supplier_address VARCHAR(255) NOT NULL COMMENT '供应商地址',
	supplier_address VARCHAR(255) NOT NULL COMMENT '供应商地址',
	supplier_address VARCHAR(255) NOT NULL COMMENT '供应商地址',
	supplier_desc text COMMENT '供应商描述',
	supplier_img VARCHAR(250) NOT NULL COMMENT '供应商图片',
	supplier_status tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '供应商状态',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY pk_product (id)
)ENGINE=innodb COMMENT '供应商信息表';

```



## 订单数据库



```sql
CREATE TABLE product_order (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	product_id varchar(16)  NOT NULL COMMENT '商品ID',
	product_name VARCHAR(250) NOT NULL COMMENT '商品名称',
	product_price DECIMAL(12,2) UNSIGNED NOT NULL default 0.00 COMMENT '商品价格',
	product_off decimal(3,2) UNSIGNED NOT NULL DEFAULT 1.00 COMMENT '商品折扣',
	product_desc text COMMENT '商品描述',
	product_img VARCHAR(250) NOT NULL default '' COMMENT '商品图片',
	product_stock INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品库存',
	product_sales INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品销量',
	product_category_id bigint UNSIGNED NULL COMMENT '商品分类ID',
	product_status tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品状态 0:未上架, 1: 上架',
	product_delete tingint unsigned NOT NULL DEFAULT 0 COMMENT '商品删除状态 0:未删除, 1: 已删除',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY pk_product (id, product_id)
)ENGINE=innodb COMMENT '订单信息表';

```


## 会员数据库


### 会员信息

```sql
CREATE TABLE client_info (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	uuid VARCHAR(16) NOT NULL UNIQUE,
	username VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  phone VARCHAR(20),
  header_img varchar(128) COMMENT '头像',
  permission varchar(16) NOT NULL DEFAULT 'client' COMMENT '会员权限, client',
  level int UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员等级',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY pk_product (id),
	index index_id(id)
	index idex_uuid (uuid),
)ENGINE=innodb COMMENT '会员信息表';

```


### 地址信息

```sql
CREATE TABLE client_address (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	uuid VARCHAR(16) NOT NULL,
	first_name VARCHAR(20) NOT NULL, 
	last_name  VARCHAR(20) NOT NULL,
	street1 VARCHAR(128) NOT NULL,
	street2 VARCHAR(128),
	ctiy VARCHAR(50) NOT NULL,
	province  VARCHAR(50) NOT NULL,
	zip_code  VARCHAR(10) NOT NULL,
	country_code VARCHAR(2) NOT NULL,
	is_default BOOLEAN DEFAULT FALSE,
	mark VARCHAR(50),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY client_address (id),
	INDEX index_uuid(uuid),
	index index_id(id)
)ENGINE=innodb COMMENT '会员地址表';

```




#### JWT-blacklists

```sql
CREATE TABLE web_jwtblacklists (
	id bigint UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
	jwt text NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY web_jwtblacklists (id),
	index index_id(id)
)ENGINE=innodb COMMENT 'jwt-blacklists';
```





#### 权限管理 数据库


客户的权限号是: client

插入一条记录
```sql
INSERT INTO casbin_rule (ptype, v0, v1, v2)
VALUES ('p', 'client', "/api/web/myself/info", "POST");
VALUES ('p', 'client', "/api/web/myself/update", "POST");
VALUES ('p', 'client', "/api/web/myself/order", "POST");
VALUES ('p', 'client', "/api/web/myself/order", "GET");
VALUES ('p', 'client', "/api/web/myself/address", "POST");
VALUES ('p', 'client', "/api/web/myself/address", "GET");
VALUES ('p', 'client', "/api/web/myself/address", "PUT");
VALUES ('p', 'client', "/api/web/myself/address", "DELETE");
```


插入多条记录
```sql
INSERT INTO casbin_rule (ptype, v0, v1, v2)
VALUES ('p', 'client', "/api/web/myself/info", "POST"),
('p', 'client', "/api/web/myself/update", "POST"),
('p', 'client', "/api/web/myself/order", "POST"),
('p', 'client', "/api/web/myself/order", "GET"),
('p', 'client', "/api/web/myself/address", "POST"),
('p', 'client', "/api/web/myself/address", "GET"),
('p', 'client', "/api/web/myself/address", "PUT"),
('p', 'client', "/api/web/myself/address", "DELETE")
```