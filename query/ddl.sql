create table genre (
  genre_id serial primary key ,
  genre_name varchar(255) NOT NULL ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ 

);

create table games_detail (
  game_detail_id SERIAL primary key ,
  description text NOT NULL ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ 

);

create table developers(
  developer_id SERIAL primary key ,
  developer_name varchar(255) NOT NULL ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ 

);

create table users (
  user_id SERIAL primary key ,
  email varchar(255)  NOT NULL,
  password_hash varchar(255) NOT NULL ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ 

);

create table games (
  game_id serial primary key ,
  developer_id integer NOT NULL REFERENCES developers(developer_id) ON DELETE CASCADE,
  genre_id integer NOT NULL REFERENCES genre(genre_id) ON DELETE CASCADE,
  game_detail_id integer NOT NULL REFERENCES games_detail(game_detail_id) ON DELETE CASCADE,
  title varchar(255) NOT NULL,
  price float NOT NULL,
  release_date Date NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ 

);

create table genre_game (
  genre_game_id serial primary key ,
  genre_id integer NOT NULL REFERENCES genre(genre_id) ON DELETE CASCADE,
  game_id integer NOT NULL REFERENCES games(game_id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ 
);

create table orders (
  order_id SERIAL primary key ,
  user_id integer NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  game_id integer NOT NULL REFERENCES games(game_id) ON DELETE CASCADE,
  order_date date NOT NULL,
  game_quantity int NOT NULL,
  total_price float NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ 

);