CREATE TYPE payment_method AS ENUM ('uzs', 'usd', 'card');
CREATE TYPE transaction_type AS ENUM ('income', 'expense');

CREATE TABLE clients
(
    id         UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    full_name  VARCHAR(60) NOT NULL,
    address    VARCHAR(50),
    phone      VARCHAR(13),
    created_at DATE DEFAULT now()
);

CREATE TABLE product_categories
(
    id         UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name       VARCHAR(50) NOT NULL,
    created_at DATE DEFAULT now()
);

CREATE TABLE products
(
    id          UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    category_id UUID REFERENCES product_categories (id) NOT NULL,
    name        VARCHAR(50)                             NOT NULL,
    bill_format VARCHAR(5)                              NOT NULL, -- Возможно, лучше использовать ENUM для форматов
    base_price  DECIMAL(10, 2)                          NOT NULL,
    total_count INT  DEFAULT 0,
    created_at  DATE DEFAULT now()
);

CREATE TABLE purchases
(
    id             UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    product_id     UUID REFERENCES products (id) NOT NULL,
    salesperson_id UUID REFERENCES clients (id)  NOT NULL,
    quantity       INT  DEFAULT 1,
    price          DECIMAL(10, 2)                NOT NULL,
    total_price    DECIMAL(10, 2)                NOT NULL,
    description    TEXT,
    created_at     DATE DEFAULT now()
);

CREATE TABLE sales
(
    id               UUID           DEFAULT gen_random_uuid() PRIMARY KEY,
    product_id       UUID REFERENCES products (id) NOT NULL,
    client_id        UUID REFERENCES clients (id)  NOT NULL,
    sale_price       DECIMAL(10, 2)                NOT NULL,
    quantity         INT            DEFAULT 0,
    total_sale_price DECIMAL(10, 2)                NOT NULL,
    payment_method   payment_method DEFAULT 'uzs',
    created_at       DATE           DEFAULT now()
);

CREATE TABLE debts
(
    id            UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    client_id     UUID REFERENCES clients (id) NOT NULL,
    amount_paid   DECIMAL(10, 2)               NOT NULL,
    amount_unpaid DECIMAL(10, 2)               NOT NULL,
    total_debt    DECIMAL(10, 2)               NOT NULL,
    next_payment  DATE,
    last_paid_day DATE DEFAULT now(),
    is_fully_paid BOOL DEFAULT FALSE,
    created_at    DATE DEFAULT now()
);

CREATE TABLE cash_category
(
    id   UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE cash_flow
(
    id               UUID           DEFAULT gen_random_uuid() PRIMARY KEY,
    transaction_date DATE           DEFAULT now(),
    amount           DECIMAL(10, 2)                     NOT NULL,
    transaction_type transaction_type                   NOT NULL,
    category_id      UUID REFERENCES cash_category (id) NOT NULL,
    description      VARCHAR(255),
    payment_method   payment_method DEFAULT 'uzs'
);