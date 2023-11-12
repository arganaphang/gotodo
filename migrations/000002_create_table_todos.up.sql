CREATE TABLE IF NOT EXISTS todos(
    id uuid PRIMARY KEY default uuid_generate_v4(),
    title VARCHAR not null,
    is_done BOOLEAN not NULL DEFAULT false,
    created_at TIMESTAMP not NULL default CURRENT_TIMESTAMP
);