-- Вставка 10 тестових котів (Spy Cats)
INSERT INTO cats (name, breed, experience, salary, created_at, updated_at) VALUES
('Whisker Shadow', 'Persian', 5, 5000.00, NOW(), NOW()),
('Midnight Prowler', 'Siamese', 3, 4200.00, NOW(), NOW()),
('Snow Assassin', 'Maine Coon', 7, 6000.00, NOW(), NOW()),
('Silent Paw', 'Bengal', 4, 4800.00, NOW(), NOW()),
('Ghost Whisper', 'Russian Blue', 6, 5500.00, NOW(), NOW()),
('Night Claw', 'Scottish Fold', 2, 3900.00, NOW(), NOW()),
('Dark Phantom', 'British Shorthair', 8, 7000.00, NOW(), NOW()),
('Stealth Hunter', 'Sphinx', 5, 5100.00, NOW(), NOW()),
('Thunder Fang', 'Abyssinian', 3, 4300.00, NOW(), NOW()),
('Crimson Stalker', 'Norwegian Forest', 6, 5800.00, NOW(), NOW());

-- Вставка 10 тестових місій (Missions)
INSERT INTO missions (cat_id, completed, created_at, updated_at) VALUES
(1, FALSE, NOW(), NOW()),
(2, FALSE, NOW(), NOW()),
(3, FALSE, NOW(), NOW()),
(4, FALSE, NOW(), NOW()),
(5, FALSE, NOW(), NOW()),
(6, FALSE, NOW(), NOW()),
(7, FALSE, NOW(), NOW()),
(8, FALSE, NOW(), NOW()),
(9, FALSE, NOW(), NOW()),
(10, FALSE, NOW(), NOW());

-- Вставка 10 тестових цілей (Targets)
INSERT INTO targets (mission_id, name, country, notes, completed, created_at, updated_at) VALUES
(1, 'Operation Nightfall', 'USA', '[{"entry": "Initial reconnaissance complete."}]'::JSONB, FALSE, NOW(), NOW()),
(2, 'Deep Cover', 'UK', '[{"entry": "Agent embedded successfully."}]'::JSONB, FALSE, NOW(), NOW()),
(3, 'Shadow Pursuit', 'Russia', '[{"entry": "Enemy movement detected."}]'::JSONB, FALSE, NOW(), NOW()),
(4, 'Silent Infiltration', 'France', '[{"entry": "Surveillance cameras bypassed."}]'::JSONB, FALSE, NOW(), NOW()),
(5, 'Ghost Recon', 'Germany', '[{"entry": "Target facility mapped."}]'::JSONB, FALSE, NOW(), NOW()),
(6, 'Stealth Strike', 'China', '[{"entry": "High-security area penetrated."}]'::JSONB, FALSE, NOW(), NOW()),
(7, 'Dark Horizon', 'Japan', '[{"entry": "Undercover asset activated."}]'::JSONB, FALSE, NOW(), NOW()),
(8, 'Silent Thunder', 'Canada', '[{"entry": "Covert extraction planned."}]'::JSONB, FALSE, NOW(), NOW()),
(9, 'Phantom Shadow', 'Italy', '[{"entry": "Disguise successfully used."}]'::JSONB, FALSE, NOW(), NOW()),
(10, 'Crimson Claw', 'Brazil', '[{"entry": "Enemy network infiltrated."}]'::JSONB, FALSE, NOW(), NOW());
