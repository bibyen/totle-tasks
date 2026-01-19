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
    goal_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(user_id),
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    -- Visibility rules: 1: Private, 2: Friends
    visibility SMALLINT DEFAULT 1, 
    -- Goals are never deleted; is_active = false means 'archived'
    is_active BOOLEAN DEFAULT TRUE, 
    create_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bingo_cards (
    bingo_card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(user_id),
    title TEXT NOT NULL,
    columns INT NOT NULL DEFAULT 5,
    rows INT NOT NULL DEFAULT 5,
    
    year INT,
    month INT,
    card_type VARCHAR(50) DEFAULT 'STANDARD',
    
    is_active BOOLEAN DEFAULT TRUE,
    create_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bingo_slots (
    bingo_card_id UUID REFERENCES bingo_cards(bingo_card_id) ON DELETE CASCADE,
    -- NO ACTION ensures goal records cannot be deleted if referenced here
    goal_id UUID REFERENCES goals(goal_id) ON DELETE NO ACTION, 
    row_index INT NOT NULL,
    column_index INT NOT NULL,
    PRIMARY KEY (bingo_card_id, row_index, column_index),
    UNIQUE(bingo_card_id, goal_id)
);