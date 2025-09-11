-- Задание: 1 Найдите номер модели, скорость и размер жесткого диска для всех ПК стоимостью менее 500 дол. Вывести: model, speed и hd
SELECT model, speed, hd FROM PC WHERE price < 500

-- Задание: 2 Найдите производителей принтеров. Вывести: maker
SELECT DISTINCT maker FROM product WHERE type='Printer'

-- Задание: 3 Найдите номер модели, объем памяти и размеры экранов ПК-блокнотов, цена которых превышает 1000 дол.
SELECT model, ram, screen FROM Laptop where price > 1000

-- Задание: 4 Найдите все записи таблицы Printer для цветных принтеров.
SELECT * FROM Printer WHERE color = 'y'

-- Задание: 5 Найдите номер модели, скорость и размер жесткого диска ПК, имеющих 12x или 24x CD и цену менее 600 дол.
SELECT model, speed, hd FROM PC WHERE (CD = '12x' OR CD = '24x') AND price < 600

-- Задание: 6 Для каждого производителя, выпускающего ПК-блокноты c объёмом жесткого диска не менее 10 Гбайт, найти скорости таких ПК-блокнотов. Вывод: производитель, скорость.
SELECT DISTINCT Product.maker,Laptop.speed
FROM Product 
    JOIN Laptop ON Product.model = Laptop.model
WHERE Laptop.hd >= 10

-- Задание: 7 Найдите номера моделей и цены всех имеющихся в продаже продуктов (любого типа) производителя B (латинская буква).
SELECT DISTINCT pr.model,
    ISNULL(pc.price, 0) + ISNULL(l.price, 0) + ISNULL(prt.price, 0) price
FROM (((Product pr 
        LEFT JOIN PC pc ON pr.model = pc.model) 
            LEFT JOIN Laptop l ON pr.model = l.model) 
                LEFT JOIN Printer prt ON pr.model = prt.model)
WHERE pr.maker = 'B'

-- Задание: 8 Найдите производителя, выпускающего ПК, но не ПК-блокноты.
SELECT DISTINCT p.maker
FROM Product p
  LEFT JOIN Product p1 ON p.maker=p1.maker AND p1.type = 'Laptop'
WHERE p.type = 'PC' and p1.maker IS NULL

-- Задание: 9 Найдите производителей ПК с процессором не менее 450 Мгц. Вывести: Maker
SELECT DISTINCT Maker
FROM Product p
  JOIN PC pc ON p.Model = pc.Model
WHERE p.type = 'PC' AND pc.speed > '450'

-- Задание: 10 Найдите модели принтеров, имеющих самую высокую цену. Вывести: model, price
SELECT model, price
FROM Printer
WHERE price = (SELECT MAX(price) FROM Printer)
