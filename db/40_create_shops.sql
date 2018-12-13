USE eve-navi;

create table Shops (
  id int PRIMARY KEY AUTO_INCREMENT,

  name varchar(25),
  genre_id int,
  address varchar(25),
  longtitude int NOT NULL,
  latitude int NOT NULL,
  
  image_id int,

  FOREIGN KEY (genre_id) REFERENCES ShopGenres(id)
);