-- Создаем ENUM типы для способов оплаты и типов транзакций
CREATE TYPE payment_method AS ENUM ('uzs', 'usd', 'card');
CREATE TYPE transaction_type AS ENUM ('income', 'expense');

-- Таблица пользователей
CREATE TABLE users
(
    user_id      UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name   VARCHAR(50) NOT NULL,
    last_name    VARCHAR(50) NOT NULL,
    email        VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(15),
    role         VARCHAR(20) NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW()
);

-- Таблица клиентов
CREATE TABLE clients
(
    id         UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    full_name  VARCHAR(60) NOT NULL,
    address    VARCHAR(50),
    phone      VARCHAR(13),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица категорий продуктов
CREATE TABLE product_categories
(
    id         UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name       VARCHAR(50) NOT NULL,
    created_by UUID REFERENCES users (user_id) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица продуктов
CREATE TABLE products
(
    id             UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    category_id    UUID REFERENCES product_categories (id) NOT NULL,
    name           VARCHAR(50) NOT NULL,
    bill_format    VARCHAR(5) NOT NULL, -- можно заменить на ENUM, если есть ограниченное количество форматов
    incoming_price DECIMAL(10, 2) NOT NULL,
    standard_price DECIMAL(10, 2) NOT NULL,
    total_count    INT DEFAULT 0,
    created_by     UUID REFERENCES users (user_id) NOT NULL,
    created_at     TIMESTAMP DEFAULT NOW()
);

-- Таблица заказов (продаж)
CREATE TABLE orders
(
    id               UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    client_id        UUID REFERENCES clients (id) NOT NULL,
    sold_by          UUID REFERENCES users (user_id) NOT NULL,
    total_sale_price DECIMAL(10, 2) NOT NULL, -- общая сумма заказа
    payment_method   payment_method DEFAULT 'uzs',
    created_at       TIMESTAMP DEFAULT NOW()
);

-- Таблица позиций заказа (каждый товар в заказе)
CREATE TABLE order_items
(
    id             UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    order_id       UUID REFERENCES orders (id) NOT NULL,
    product_id     UUID REFERENCES products (id) NOT NULL,
    quantity       INT DEFAULT 1 NOT NULL,
    sale_price     DECIMAL(10, 2) NOT NULL,
    total_price    DECIMAL(10, 2) NOT NULL -- общая цена за конкретный товар в заказе
);

-- Таблица категорий денежных потоков
CREATE TABLE cash_category
(
    id   UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Таблица денежных потоков (доходы и расходы)
CREATE TABLE cash_flow
(
    id               UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id          UUID REFERENCES users (user_id) NOT NULL,
    transaction_date TIMESTAMP DEFAULT NOW(),
    amount           DECIMAL(10, 2) NOT NULL,
    transaction_type transaction_type NOT NULL,
    category_id      UUID REFERENCES cash_category (id) NOT NULL,
    description      VARCHAR(255),
    payment_method   payment_method DEFAULT 'uzs'
);

-- Таблица задолженностей (по каждому заказу)
CREATE TABLE debts
(
    id            UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    order_id      UUID REFERENCES orders (id) NOT NULL, -- Привязка долга к заказу
    amount_paid   DECIMAL(10, 2) NOT NULL, -- Сумма, уже оплаченная
    amount_unpaid DECIMAL(10, 2) NOT NULL, -- Сумма, остающаяся к оплате
    total_debt    DECIMAL(10, 2) NOT NULL, -- Общая сумма долга
    next_payment  DATE,
    last_paid_day TIMESTAMP DEFAULT NOW(),
    is_fully_paid BOOLEAN DEFAULT FALSE,
    recipient_id  UUID REFERENCES users (user_id) NOT NULL, -- Кто принял платёж
    created_at    TIMESTAMP DEFAULT NOW()
);

-- Таблица для частичных выплат по задолженности
CREATE TABLE debt_payments
(
    id           UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    debt_id      UUID REFERENCES debts (id) NOT NULL, -- Привязка к задолженности
    payment_date TIMESTAMP DEFAULT NOW(),
    amount       DECIMAL(10, 2) NOT NULL, -- Сумма частичного платежа
    paid_by      UUID REFERENCES users (user_id) NOT NULL -- Кто внес платёж
);
