USE eve-navi;

CREATE TABLE Shops (
  id INT PRIMARY KEY AUTO_INCREMENT,

  name VARCHAR(25),
  genre_id INT,
  address VARCHAR(25),
  longtitude INT NOT NULL,
  latitude INT NOT NULL,
  
  image_id INT,

  FOREIGN KEY (genre_id) REFERENCES ShopGenres(id)
);