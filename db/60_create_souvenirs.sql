USE eve-navi;

create table Souvenirs (
  id int PRIMARY KEY AUTO_INCREMENT,

  name varchar(25),
  genre_id int,
  price int,

  image_id int,

  FOREIGN KEY (genre_id) REFERENCES ShopGenres(id)
);