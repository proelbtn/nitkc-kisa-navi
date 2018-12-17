USE eve-navi;

CREATE TABLE Shops (
  id INT PRIMARY KEY AUTO_INCREMENT,

  name VARCHAR(25) NOT NULL,
  genre_id INT NOT NULL,
  address VARCHAR(25) NOT NULL,
  longitude INT NOT NULL,
  latitude INT NOT NULL,
  open INT,
  close INT,
  
  image_id INT,

  FOREIGN KEY (genre_id) REFERENCES ShopGenres(id)
);