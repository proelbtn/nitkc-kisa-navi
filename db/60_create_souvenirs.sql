USE eve-navi;

CREATE TABLE Souvenirs (
  id INT PRIMARY KEY AUTO_INCREMENT,

  name VARCHAR(25),
  genre_id INT,
  price INT,

  image_id INT,

  FOREIGN KEY (genre_id) REFERENCES ShopGenres(id)
);