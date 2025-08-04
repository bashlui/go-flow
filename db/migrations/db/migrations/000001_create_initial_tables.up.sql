CREATE TABLE stocks (
    symbol VARCHAR(10) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    last_price NUMERIC(10, 2),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE news_articles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stock_symbol VARCHAR(10) REFERENCES stocks(symbol),
    headline TEXT NOT NULL,
    source VARCHAR(255),
    published_at TIMESTAMP WITH TIME ZONE NOT NULL,
    sentiment TEXT
);