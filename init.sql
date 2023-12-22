-- init.sql

CREATE TABLE IF NOT EXISTS dish (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC NOT NULL,
    description TEXT,
    image_url TEXT,
    tags JSONB
);

INSERT INTO dish (name, price, description, image_url, tags)
VALUES 
    ('Цыплёнок Грин', 345, 'Сочное куриное филе,сыр моцарелла, сладкий перец, соус «Чесночный рэнч», чеснок и томаты.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/7731ed084014c6954824e7834a7aab24.webp', '["Все меню", "Пицца"]'),
    ('Пепперони Грин', 360, 'Пикантная пепперони, сладкий перец, Моцарелла с фирменным томатным соусом.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/59ae0f885ff8f5bc8810ca659f568196.webp', '["Все меню", "Пицца"]'),
    ('Мясное барбекю', 400, 'Мясная пицца с пикантной пепперони, альпийскими колбасками, Моцареллой, луком и соусом Барбекю.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/ed0545535948902e28d9b3bf0b824cdc.webp', '["Все меню", "Пицца"]'),
    ('Вестерн Барбекю', 699, 'Сезон барбекю! Ароматная говядина, сочные томаты, лук с сыром Моцарелла и соусом барбекю', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/8d1349fcfb8713f2611bac5850f6183b.webp', '["Все меню", "Пицца"]'),
    ('Evervess Манящая Гранада', 100, 'газированный напиток с классическим вкусом гренадина и содовой.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/d2e907615430c4804cb4194c0a24ebd1.webp', '["Все меню", "Напитки"]'),
    ('Evervess Биттер-Лимон', 110, 'безалкогольный газированный напиток с терпким цитрусовым вкусом лимона и легкой горчинкой.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/5c383c9f9ef2ad4a656a70a68fea948f.webp', '["Все меню", "Напитки"]'),
    ('Русский дар Лимонад', 105, 'безалкогольный газированный напиток с лимонным вкусом.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/771050f917573e7f5c45acc351afa6cb.webp', '["Все меню", "Напитки"]'),
    ('Узелки сладкие', 159, 'Мягкие булочки с корицей и сахарной пудрой.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/Sweet%20knots.jpg', '["Все меню", "Десерты"]'),
    ('Чизкейк Нью-Йорк', 249, 'Классический чизкейк из сливочного сыра, с добавлением натуральной ванили.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/d8472a931a6f6472cff33bc11b18a7ad.webp', '["Все меню", "Десерты"]'),
    ('Pаpа Gamе', 1399, 'Мощный заряд! Супер Папа 30см и Пепперони 30см, Evervess Апельсин 1 л.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/49ea43aa652dccf6073b46ba0be499c4.webp', '["Все меню", "Комбо"]'),
    ('Мяснoй сeт', 1028, 'Сытная тройка: Супер Папа и Мясная 23 см, Evervess Искрящийся Лемон-Лайм 1 л.', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/92bd249b8b0fbb7831e0ac06b0c808b9.webp', '["Все меню", "Комбо"]'),
    ('Сет Кидс с рогаликами', 699, 'Пицца улыбка, рогалики с колбасками, Аква минерали без газов', 'https://cdn.papajohns.ru//images/catalog/thumbs/cart/b322ee3a3871ff2c2815377cabd42c7a.webp', '["Все меню", "Комбо"]');

CREATE TABLE IF NOT EXISTS user_info (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    phone INT NOT NULL,
    address TEXT NOT NULL
);

INSERT INTO user_info (login, password, email, phone, address)
VALUES 
    ('example_user', 'password123', 'user@example.com', 1234567890, '123 Main St');
