create table users (
    id serial primary key,
    provider varchar(50) not null,
    provider_id varchar(255) not null,
    email varchar(255) not null,
    name varchar(255),
    profile_picture_url varchar(255),
    access_token varchar(255),
    refresh_token varchar(255),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);


insert into users (provider, provider_id, email, name, profile_picture_url, access_token, refresh_token)
values 
('google', 'google-id-123', 'user1@gmail.com', 'User One', 'http://example.com/user1.jpg', 'access-token-123', 'refresh-token-123'),
('facebook', 'facebook-id-456', 'user2@facebook.com', 'User Two', 'http://example.com/user2.jpg', 'access-token-456', 'refresh-token-456'),
('apple', 'apple-id-789', 'user3@apple.com', 'User Three', 'http://example.com/user3.jpg', 'access-token-789', 'refresh-token-789');