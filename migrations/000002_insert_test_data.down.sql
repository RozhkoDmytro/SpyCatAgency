-- Видалення тестових цілей
DELETE FROM targets WHERE name IN (
    'Operation Nightfall', 'Deep Cover', 'Shadow Pursuit', 'Silent Infiltration', 
    'Ghost Recon', 'Stealth Strike', 'Dark Horizon', 'Silent Thunder', 
    'Phantom Shadow', 'Crimson Claw'
);

-- Видалення тестових місій
DELETE FROM missions WHERE id BETWEEN 1 AND 10;

-- Видалення тестових котів
DELETE FROM cats WHERE name IN (
    'Whisker Shadow', 'Midnight Prowler', 'Snow Assassin', 'Silent Paw', 
    'Ghost Whisper', 'Night Claw', 'Dark Phantom', 'Stealth Hunter', 
    'Thunder Fang', 'Crimson Stalker'
);
