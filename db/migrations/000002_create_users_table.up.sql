create table users (
    `id` int unsigned auto_increment primary key,
    `username` varchar(255) unique,
    `password` varchar(255),
    `created_at` datetime(3),
    `updated_at` datetime(3),
    `deleted_at` datetime(3)
)