-- feimall.billing_address definition

CREATE TABLE `billing_address` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户 UUID',
  `first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '持卡人名字',
  `last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '持卡人姓氏',
  `holder_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '持卡人全名',
  `line1` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '地址行1',
  `line2` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地址行2',
  `city` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '城市',
  `state` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '州/省份',
  `country` varchar(2) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '国家代码，ISO 3166-1 标准',
  `postal_code` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮政编码',
  `district` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '行政区',
  `card_number` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '卡号',
  `card_type` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '卡类型：D=借记，C=信用卡',
  `bank_code` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '银行编码',
  `card_brand` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '信用卡品牌，例如Visa、MasterCard',
  `card_expiration_year` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '卡年有效期',
  `card_expiration_month` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '卡月有效期',
  `cvv` int DEFAULT NULL COMMENT '卡安全码',
  `phone_number` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号，格式“+区号-手机号”',
  `email` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱地址',
  `llpay_token` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '连连支付Token',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `index_user_id` (`user_id`),
  KEY `index_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='信用卡注册地址表';


-- feimall.cart definition

CREATE TABLE `cart` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '购物车ID',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `product_id` varchar(64) NOT NULL COMMENT '商品ID',
  `quantity` int unsigned NOT NULL DEFAULT '1' COMMENT '商品数量',
  `price` decimal(12,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `old_price` decimal(12,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `price_off` decimal(3,0) unsigned NOT NULL DEFAULT '100' COMMENT '商品折扣',
  `stock` int unsigned NOT NULL DEFAULT '0' COMMENT '商品库存',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `main_img` varchar(255) NOT NULL DEFAULT '' COMMENT '商品主图',
  `size_title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品尺寸标题',
  `color` varchar(16) NOT NULL DEFAULT '' COMMENT '商品颜色',
  `size` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品尺寸',
  `status` int unsigned NOT NULL DEFAULT '1' COMMENT '商品状态 0:删除购物车, 1: 正常',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='购物车表';


-- feimall.casbin_rule definition

CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `v0` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `v1` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `v2` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `v3` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='api 权限';


-- feimall.category_info definition

CREATE TABLE `category_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `tag_id` bigint unsigned NOT NULL COMMENT '商品标签ID',
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品ID',
  `title` varchar(255) DEFAULT NULL COMMENT '商品标签名称',
  `value` varchar(128) DEFAULT NULL COMMENT '商品标签值',
  `level` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '分类级别',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_level` (`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品信息表';


-- feimall.client_address definition

CREATE TABLE `client_address` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `line1` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `postal_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `country_name` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_default` tinyint(1) DEFAULT '0',
  `mark` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `country` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '国家',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '收货人电话',
  `line2` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间', 
  PRIMARY KEY (`id`),
  KEY `index_uuid` (`user_id`),
  KEY `index_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员地址表';


-- feimall.client_info definition

CREATE TABLE `client_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'userID',
  `username` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户登录密码',
  `email` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
  `phone` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户手机号',
  `nick_name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户昵称',
  `header_img` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  `uuid` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户UUID',
  `level` bigint unsigned DEFAULT NULL COMMENT '用户等级',
  `introducer` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '介绍人 id',
  `permission` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户权限',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_client_info_deleted_at` (`deleted_at`),
  KEY `idx_client_info_uuid` (`uuid`),
  KEY `idx_client_info_user_id` (`user_id`),
  KEY `idx_client_info_introducer` (`introducer`),
  KEY `idx_client_info_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员信息表';


-- feimall.document definition

CREATE TABLE `document` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '文档ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文档标题',
  `content` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文档内容',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 0:deleted, 1:active',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_document_id` (`id`),
  KEY `idx_title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文档表';


-- feimall.jwt_blacklists definition

CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text COLLATE utf8mb4_unicode_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='jwt';


-- feimall.media_asset definition

CREATE TABLE `media_asset` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '唯一ID',
  `uuid` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '全局唯一标识（用于业务关联）',
  `oss_key` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'OSS存储路径（如 products/2023/abc.jpg）',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '完整访问URL',
  `file_hash` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件哈希（用于去重和校验）',
  `file_size` int unsigned NOT NULL COMMENT '文件大小（字节）',
  `mime_type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件类型（如 image/jpeg）',
  `uploader_id` bigint unsigned DEFAULT NULL COMMENT '上传者ID（用户/管理员）',
  `asset_type` tinyint unsigned NOT NULL COMMENT '资源类型：1=商品展示图, 2=供应商样图, 3=公用说明图, 4=用户评论图, 5=其他',
  `ref_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '关联业务ID（如商品ID、评论ID）',
  `is_linked` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被业务关联（0=未关联可清理, 1=已关联）',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_oss_key` (`oss_key`),
  UNIQUE KEY `uniq_uuid` (`uuid`),
  KEY `idx_asset_type_ref` (`asset_type`,`ref_id`),
  KEY `idx_file_hash` (`file_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='统一媒体资源表';


-- feimall.orders definition

CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `order_id` varchar(64) NOT NULL COMMENT '订单号',
  `total_price` decimal(12,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '订单总价格',
  `discount` decimal(12,2) unsigned DEFAULT '0.00' COMMENT '折扣金额',
  `payment_method` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '支付方式 (例如: PayPal, Credit Card, etc.)',
  `payment_status` varchar(32) NOT NULL DEFAULT 'pending' COMMENT '支付状态 (pending, paid, failed)',
  `order_status` varchar(32) NOT NULL DEFAULT 'pending' COMMENT '订单状态 (pending, processing, shipped, completed, cancelled)',
  `shipping_company_id` bigint unsigned DEFAULT '0' COMMENT '快递公司ID，关联到shipping_company表',
  `shipping_method` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '配送方式',
  `shipping_price` decimal(12,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '配送费用',
  `shipping_address_id` bigint unsigned NOT NULL COMMENT '配送地址ID，关联到client_address表',
  `tracking_number` varchar(64) DEFAULT NULL COMMENT '快递单号',
  `mark` varchar(255) DEFAULT NULL COMMENT '客户备注',
  `note` varchar(255) DEFAULT NULL COMMENT '商家内部备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '订单创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '订单删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_id` (`order_id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_order_status` (`order_status`),
  KEY `idx_payment_status` (`payment_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单表';


-- feimall.orders_product definition

CREATE TABLE `orders_product` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单商品ID',
  `order_id` varchar(64) NOT NULL COMMENT '订单号，关联到orders表',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `product_id` varchar(64) NOT NULL COMMENT '商品ID，关联到product表',
  `quantity` int unsigned NOT NULL DEFAULT '1' COMMENT '商品数量',
  `price` decimal(12,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `old_price` decimal(12,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `price_off` decimal(3,0) unsigned NOT NULL DEFAULT '100' COMMENT '商品折扣',
  `stock` int unsigned NOT NULL DEFAULT '0' COMMENT '商品库存',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `main_img` varchar(255) NOT NULL DEFAULT '' COMMENT '商品主图',
  `size_title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品尺寸标题',
  `color` varchar(16) NOT NULL DEFAULT '' COMMENT '商品颜色',
  `size` varchar(16) NOT NULL DEFAULT '' COMMENT '商品尺寸',
  `mark` varchar(255) NOT NULL DEFAULT '' COMMENT '商品备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单商品详情表';


-- feimall.payment_llpay definition

CREATE TABLE `payment_llpay` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `order_id` varchar(64) NOT NULL COMMENT '订单号，关联到orders表',
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户 ID',
  `ll_transaction_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '第三方支付系统生成的交易 ID',
  `merchant_transaction_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商户生成的交易 ID',
  `payment_currency_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付币种，例如 "USD"',
  `payment_amount` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付金额，保留两位小数',
  `exchange_rate` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '汇率，精度为 8 位',
  `payment_status` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付状态（例如 "WP" - 待支付，"PS" - 支付成功）',
  `settlement_currency_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '结算币种，通常与支付币种相同',
  `settlement_amount` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '结算金额，保留两位小数',
  `installments` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分期付款数量，默认为 1，表示一次性支付',
  `payment_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付 URL，用于支付过程中的重定向链接',
  `trace_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '追踪 ID，用于唯一标识支付流程，便于追踪',
  `key_value` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付相关的 key，通常用于防止重复支付等用途',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间，默认为当前时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  `payment_time` varchar(32) DEFAULT NULL COMMENT 'payment_time',
  PRIMARY KEY (`id`),
  KEY `idx_merchant_transaction_id` (`merchant_transaction_id`) COMMENT '商户交易 ID 索引，用于查询特定商户的支付记录',
  KEY `idx_ll_transaction_id` (`ll_transaction_id`) COMMENT '第三方支付系统交易 ID 索引，用于查询特定支付系统的支付记录',
  KEY `idx_payment_status` (`payment_status`) COMMENT '支付状态索引，用于根据支付状态查询支付记录',
  KEY `idx_order_id` (`order_id`) COMMENT 'order_id 索引，用于查询特定订单的支付记录',
  KEY `idx_user_id` (`user_id`) COMMENT '用户 ID 索引，用于查询特定用户的支付记录',
  KEY `idx_trace_id` (`trace_id`) COMMENT '追踪 ID 索引，用于查询特定支付流程',
  KEY `idx_created_at` (`created_at`) COMMENT '创建时间索引，用于查询支付记录的时间范围'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='连连支付交易记录表';


-- feimall.payment_logs definition

CREATE TABLE `payment_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `payment_transaction_id` bigint unsigned NOT NULL COMMENT '关联支付交易ID，关联到payment_transactions表',
  `log_type` enum('info','warning','error') NOT NULL COMMENT '日志类型',
  `log_message` text NOT NULL COMMENT '日志信息',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '日志创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_payment_transaction_id` (`payment_transaction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='支付日志表，记录每次支付的详细状态及变更';


-- feimall.payment_providers definition

CREATE TABLE `payment_providers` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '支付平台ID',
  `provider_name` varchar(64) NOT NULL COMMENT '支付平台名称 (如: PayPal, Stripe, dLocal)',
  `api_key` varchar(255) NOT NULL COMMENT '支付平台的API密钥',
  `api_secret` varchar(255) NOT NULL COMMENT '支付平台的API密钥',
  `is_active` tinyint(1) NOT NULL DEFAULT '1' COMMENT '支付平台是否启用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_provider_name` (`provider_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='支付平台配置表';


-- feimall.payment_transactions definition

CREATE TABLE `payment_transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '支付交易ID',
  `order_id` varchar(64) NOT NULL COMMENT '关联订单号，关联到orders表',
  `provider_id` int unsigned NOT NULL COMMENT '支付平台ID，关联到payment_providers表',
  `transaction_id` varchar(255) DEFAULT NULL COMMENT '支付平台返回的交易ID',
  `amount` decimal(12,2) NOT NULL COMMENT '支付金额',
  `currency` varchar(3) NOT NULL COMMENT '支付货币 (USD, CAD, EUR等)',
  `status` enum('pending','completed','failed','refunded') NOT NULL DEFAULT 'pending' COMMENT '支付状态',
  `payment_date` timestamp NULL DEFAULT NULL COMMENT '支付完成时间',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '交易更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_transaction_id` (`transaction_id`),
  KEY `idx_provider_id` (`provider_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='支付交易表，记录每笔订单的支付信息';


-- feimall.product definition

CREATE TABLE `product` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品名称',
  `desction` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品描述',
  `seo_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'SEO关键词',
  `seo_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'SEO描述',
  `price` double DEFAULT NULL COMMENT '商品价格',
  `price_off` double DEFAULT NULL COMMENT '商品折扣',
  `main_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品图片',
  `stock` bigint unsigned DEFAULT NULL COMMENT '商品库存',
  `sales` bigint unsigned DEFAULT NULL COMMENT '商品销量',
  `status` bigint DEFAULT NULL COMMENT '商品状态 0:未上架, 1: 上架',
  `is_delete` bigint DEFAULT NULL COMMENT '商品删除状态 0:未删除, 1: 已下架',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_status` (`status`),
  KEY `idx_is_delete` (`is_delete`),
  KEY `idx_title` (`title`),
  KEY `idx_product_deleted_at` (`deleted_at`),
  KEY `idx_product_product_id` (`product_id`),
  KEY `idx_product_title` (`title`),
  KEY `idx_product_status` (`status`),
  KEY `idx_product_is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品信息表';


-- feimall.product_brand definition

CREATE TABLE `product_brand` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品ID',
  `brand_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '品牌 ID',
  `brand_title` varchar(128) NOT NULL DEFAULT '' COMMENT '品牌名称',
  `shape_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '类型 ID',
  `shape_title` varchar(128) NOT NULL DEFAULT '' COMMENT '类型名称',
  `tag_id` bigint unsigned DEFAULT NULL COMMENT '商品SKU标签ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品与品牌关联表';


-- feimall.product_comment definition

CREATE TABLE `product_comment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` varchar(64) NOT NULL COMMENT '商品ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `content` text NOT NULL COMMENT '评论内容',
  `is_img` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否图片评论',
  `img_url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片地址',
  `star` int unsigned NOT NULL DEFAULT '5' COMMENT '评论星级',
  `shop_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '商家回复内容',
  `title` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标题',
  `date` varchar(100) DEFAULT NULL COMMENT '数据中的日期, 临时数据',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品评论信息表';


-- feimall.product_detail definition

CREATE TABLE `product_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` varchar(64) NOT NULL COMMENT '商品ID',
  `content` text NOT NULL COMMENT '商品详情内容',
  `lang` varchar(16) NOT NULL COMMENT '语言',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品详情页面信息表';


-- feimall.product_img definition

CREATE TABLE `product_img` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品ID',
  `img_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品图片地址',
  `sort_id` int unsigned NOT NULL DEFAULT '0' COMMENT '图片排序序号, 0表示主图, 1..n表示其他图片',
  `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '图片类型 1:商品展示图, 2:商品详情图, 3: 商品轮播图, 9: 其他图片',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '商品状态 0:删除, 1: 上架',
  `name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '""' COMMENT '图片名称',
  `tag` varchar(32) NOT NULL DEFAULT '""' COMMENT '图片标签',
  `uuid` varchar(64) NOT NULL DEFAULT '' COMMENT '图片UUID',
  `is_main` int NOT NULL DEFAULT '0' COMMENT '是否为主图, 默认是 0, 1 为主图',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_sort_id` (`sort_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品图片信息表';


-- feimall.product_reviews definition

CREATE TABLE `product_reviews` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` varchar(64) NOT NULL COMMENT '商品ID',
  `star1` int unsigned NOT NULL DEFAULT '0' COMMENT '1星级得分总和',
  `star2` int unsigned NOT NULL DEFAULT '0' COMMENT '2星级得分总和',
  `star3` int unsigned NOT NULL DEFAULT '0' COMMENT '3星级得分总和',
  `star4` int unsigned NOT NULL DEFAULT '0' COMMENT '4星级得分总和',
  `star5` int unsigned NOT NULL DEFAULT '0' COMMENT '5星级得分总和',
  `total` int unsigned NOT NULL DEFAULT '0' COMMENT '总得分',
  `reviews` int unsigned NOT NULL DEFAULT '0' COMMENT '评论总数',
  `average` double NOT NULL DEFAULT '0' COMMENT '平均分',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品评论信息表';


-- feimall.product_sku definition

CREATE TABLE `product_sku` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` varchar(64) NOT NULL DEFAULT '' COMMENT '商品ID',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父级SKUID',
  `tag_id` bigint unsigned NOT NULL COMMENT '商品SKU标签ID',
  `title` varchar(255) NOT NULL COMMENT '商品名称',
  `stock` int unsigned NOT NULL DEFAULT '0' COMMENT '商品库存',
  `price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `price_off` decimal(3,0) unsigned NOT NULL DEFAULT '100' COMMENT '商品折扣',
  `main_img` varchar(255) NOT NULL DEFAULT '' COMMENT '商品主图',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品SKU信息表';


-- feimall.supplier definition

CREATE TABLE `supplier` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '供货商ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '供货商名称',
  `person_cn` varchar(255) NOT NULL DEFAULT '' COMMENT '联系人姓名',
  `person_en` varchar(255) NOT NULL DEFAULT '' COMMENT '联系人姓名',
  `wechat` varchar(255) NOT NULL DEFAULT '' COMMENT 'wechat',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '联系人电子邮件',
  `phone1` varchar(20) NOT NULL DEFAULT '' COMMENT '联系人电话',
  `phone2` varchar(20) NOT NULL DEFAULT '' COMMENT '联系人电话',
  `line1` varchar(255) NOT NULL DEFAULT '' COMMENT '地址行1',
  `line2` varchar(255) DEFAULT NULL COMMENT '地址行2',
  `city` varchar(100) NOT NULL DEFAULT '' COMMENT '城市',
  `state` varchar(100) NOT NULL DEFAULT '' COMMENT '州/省',
  `country` varchar(100) NOT NULL DEFAULT '中国' COMMENT '国家',
  `zip` varchar(20) NOT NULL DEFAULT '' COMMENT '邮政编码',
  `nice_name` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_supplier_id` (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='供货商表';


-- feimall.supplier_product definition

CREATE TABLE `supplier_product` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '供货商货品ID',
  `supplier_id` bigint unsigned DEFAULT NULL COMMENT '供货商ID',
  `product_id` varchar(191) DEFAULT NULL COMMENT '商品ID',
  `supplier_product_code` varchar(191) DEFAULT NULL COMMENT '供货商货品代码',
  `cost_price` double DEFAULT NULL COMMENT '供货商提供的成本价格',
  `stock` bigint unsigned DEFAULT NULL COMMENT '库存数量',
  `lead_time` bigint DEFAULT NULL COMMENT '交货时间(天) 0天是有现货',
  `size1` varchar(191) DEFAULT NULL COMMENT '货品尺寸1',
  `size2` varchar(191) DEFAULT NULL COMMENT '货品尺寸2',
  `size3` varchar(191) DEFAULT NULL COMMENT '货品尺寸3',
  `size4` varchar(191) DEFAULT NULL COMMENT '货品尺寸4',
  `size5` varchar(191) DEFAULT NULL COMMENT '货品尺寸5',
  `num1` bigint DEFAULT NULL COMMENT '货品数量1',
  `num2` bigint DEFAULT NULL COMMENT '货品数量2',
  `num3` bigint DEFAULT NULL COMMENT '货品数量3',
  `num4` bigint DEFAULT NULL COMMENT '货品数量4',
  `num5` bigint DEFAULT NULL COMMENT '货品数量5',
  `main_img` varchar(191) DEFAULT NULL COMMENT '商品主图',
  `status` bigint DEFAULT NULL COMMENT '货品状态 0:不可用, 1:可用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_supplier_id` (`supplier_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_supplier_product_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='供货商货品表';


-- feimall.tag_info definition

CREATE TABLE `tag_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品标签ID',
  `tag_id` bigint unsigned NOT NULL COMMENT '商品标签ID',
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '商品ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品标签名称',
  `value` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品标签值',
  `level` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '商品标签级别, 级别越高, 优先级越高, 1-9',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品标签信息表';


-- feimall.tags definition

CREATE TABLE `tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品标签ID',
  `type` tinyint unsigned NOT NULL COMMENT '商品标签类型, 0:分类标签, 1:表示商品标签, 2:表示商品sku类型',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父级标签ID',
  `title` varchar(255) NOT NULL COMMENT '商品标签名称',
  `title_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品EN 标题',
  `value` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品标签值',
  `value_cm` varchar(64) DEFAULT NULL COMMENT '厘米',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_tag_id` (`id`),
  KEY `idx_tag_parent_id` (`parent_id`),
  KEY `idx_tag_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品标签表';


-- feimall.web_jwtblacklists definition

CREATE TABLE `web_jwtblacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `jwt` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `index_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='jwt-blacklists';