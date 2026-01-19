CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    display_name VARCHAR(100),
    avatar_url TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_identities (
    identity_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    provider_type VARCHAR(20) NOT NULL, -- 'local', 'google'?
    provider_key VARCHAR(255) NOT NULL,  -- email or oauth_id
    password_hash TEXT,                  
    UNIQUE(provider_type, provider_key)
);

CREATE TABLE friendships (
    user_id_1 UUID NOT NULL REFERENCES users(user_id),
    user_id_2 UUID NOT NULL REFERENCES users(user_id),
    initiator_id UUID NOT NULL REFERENCES users(user_id),
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING', -- 'PENDING', 'ACCEPTED', 'BLOCKED'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id_1, user_id_2),
    CONSTRAINT pair_order CHECK (user_id_1 < user_id_2),
    CONSTRAINT initiator_is_part_of_pair CHECK (initiator_id IN (user_id_1, user_id_2))
);

CREATE TABLE goals (
    goal_id UUID PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    visibility INT DEFAULT 1,
    is_active BOOLEAN DEFAULT TRUE,
    is_assigned BOOLEAN DEFAULT FALSE, -- <--- Add this column
    create_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bingo_cards (
    bingo_card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(user_id),
    title TEXT NOT NULL,
    columns INT NOT NULL DEFAULT 5 CHECK (columns = 5),
    rows INT NOT NULL DEFAULT 5 CHECK (rows = 5),
    year INT,
    month INT,
    is_active BOOLEAN DEFAULT TRUE,
    create_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bingo_slots (
    bingo_card_id UUID REFERENCES bingo_cards(bingo_card_id) ON DELETE CASCADE,
    goal_id UUID REFERENCES goals(goal_id) ON DELETE NO ACTION, 
    row_index INT NOT NULL CHECK (row_index BETWEEN 0 AND 4),
    column_index INT NOT NULL CHECK (column_index BETWEEN 0 AND 4),
    PRIMARY KEY (bingo_card_id, row_index, column_index),
    UNIQUE(bingo_card_id, goal_id)
);
