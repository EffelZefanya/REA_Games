-- Seed data for genre table
INSERT INTO genre (genre_name) VALUES
('Action'),
('Adventure'),
('RPG'),
('Strategy'),
('Simulation'),
('Sports');

-- Seed data for games_detail table
INSERT INTO games_detail (description) VALUES
('An action-packed game with thrilling combat sequences.'),
('A story-driven game with open-world exploration.'),
('An immersive role-playing game with deep character development.'),
('A tactical strategy game with multiplayer modes.'),
('A simulation game with realistic management mechanics.'),
('A sports simulation game with real-life teams and players.');

-- Seed data for developers table
INSERT INTO developers (developer_name) VALUES
('Naughty Dog'),
('Rockstar Games'),
('Bethesda'),
('Blizzard Entertainment'),
('EA Sports'),
('Ubisoft');

-- Seed data for users table
INSERT INTO users (username, email, password_hash) VALUES
('john_doe', 'john.doe@example.com', 'hashedpassword1'),
('jane_smith', 'jane.smith@example.com', 'hashedpassword2'),
('alex_williams', 'alex.williams@example.com', 'hashedpassword3'),
('lily_brown', 'lily.brown@example.com', 'hashedpassword4'),
('james_jones', 'james.jones@example.com', 'hashedpassword5');

-- Seed data for games table
INSERT INTO games (developer_id, genre_id, game_detail_id, title, price, release_date) VALUES
(1, 1, 1, 'Uncharted 4', 59.99, '2016-05-10'),
(2, 2, 2, 'Red Dead Redemption 2', 69.99, '2018-10-26'),
(3, 3, 3, 'The Elder Scrolls V: Skyrim', 39.99, '2011-11-11'),
(4, 4, 4, 'Overwatch', 39.99, '2016-05-24'),
(5, 5, 5, 'The Sims 4', 49.99, '2014-09-02'),
(6, 6, 6, 'FIFA 21', 59.99, '2020-10-09');

-- Seed data for genre_game table (if a game can have multiple genres)
INSERT INTO genre_game (genre_id, game_id) VALUES
(1, 1), -- Uncharted 4 is an Action game
(2, 2), -- Red Dead Redemption 2 is an Adventure game
(3, 3), -- Skyrim is an RPG
(4, 4), -- Overwatch is a Strategy game
(5, 5), -- The Sims 4 is a Simulation game
(6, 6), -- FIFA 21 is a Sports game
(1, 6), -- FIFA 21 is also an Action game
(2, 1); -- Uncharted 4 is also an Adventure game

-- Seed data for orders table
INSERT INTO orders (users_id, game_id, order_date, game_quantity, total_price) VALUES
(1, 1, '2023-11-25', 1, 59.99),
(2, 2, '2023-11-26', 2, 139.98),
(3, 3, '2023-11-27', 1, 39.99),
(4, 4, '2023-11-28', 3, 119.97),
(5, 5, '2023-11-29', 1, 49.99);