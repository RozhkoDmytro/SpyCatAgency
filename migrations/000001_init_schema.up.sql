-- migrations/000001_init_schema.up.sql
CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    breed VARCHAR(50) NOT NULL,
    experience INT NOT NULL CHECK (experience >= 0),
    salary DECIMAL(10,2) NOT NULL CHECK (salary >= 0),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE missions (
    id SERIAL PRIMARY KEY,
    cat_id INT REFERENCES cats(id) ON DELETE SET NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);
;

CREATE TABLE targets (
    id SERIAL PRIMARY KEY,
    mission_id INT REFERENCES missions(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    notes JSONB DEFAULT '[]'::JSONB, 
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    UNIQUE (mission_id, name)
);

CREATE INDEX idx_cats_deleted_at ON cats(deleted_at);
CREATE INDEX idx_missions_deleted_at ON missions(deleted_at);
CREATE INDEX idx_targets_deleted_at ON targets(deleted_at);
CREATE INDEX idx_missions_cat_id ON missions(cat_id);
CREATE INDEX idx_targets_mission_id ON targets(mission_id);