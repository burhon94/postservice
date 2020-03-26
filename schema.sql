CREATE TABLE if not exists post
(
    id          BIGINT PRIMARY KEY,
    title       TEXT,
    dataandtime TIMESTAMP,
    poster      TEXT,
    category    TEXT,
    urlfile     TEXT
);



INSERT INTO post (title, dataandtime, poster, category, urlfile)
VALUES ('Steam Gamer', '2020-03-20 10:24:05.000000',
        'http://localhost:20000/web/media/eeb269f5-b58b-4341-a650-d096f09ffe87.jpg', 'game',
        'http://localhost:20000/web/media/8fab677a-7fb1-4dec-b4d4-caeed46e2e83.exe'),
       ('Kung Fu Children 1986', '2020-03-20 10:23:54.000000',
        'http://localhost:20000/web/media/93b5213e-42f1-4249-8f76-6ac419575fc5.jpg', 'movie',
        'http://localhost:20000/web/media/310c00e9-8dfe-499f-8388-1c18abadaaaf.mp4'),
       ('Zip Activator', '2020-03-20 10:24:02.000000',
        'http://localhost:20000/web/media/dc4e5a17-8dc8-4dd8-bd0b-33de5b3148b4.jpg', 'prog',
        'http://localhost:20000/web/media/e72f07d1-611e-444c-9f5f-4d04212f286e.exe'),
       ('Grand theft Auto San Andreas',
        '2020-03-20 10:23:58.000000', 'http://localhost:20000/web/media/02d60565-e2de-470e-ab67-a6f6b28387f6.jpg',
        'game', 'http://localhost:20000/web/media/51548f13-797d-4de7-ba8d-5ad28886687d.exe'),

       ('X Plane Tutorial',
        '2020-03-20 10:24:05.000000', 'http://localhost:20000/web/media/43874b4e-94d3-4b5a-8db3-fc2083df0692.jpg',
        'movie', 'http://localhost:20000/web/media/2f49dacc-bc7d-46d5-8390-829fe21b01e5.mp4'),
       ('Era - Ameno',
        '2020-03-20 10:24:05.000000', 'http://localhost:20000/web/media/e6d3681e-4f6c-4cdb-a88f-f4d2d6cd84d4.jpg',
        'music', 'http://localhost:20000/web/media/11331a84-25a1-4419-99c1-f1a48899f0ae.mp3'),
       ('Supertramp - The logical Song',
        '2020-03-20 10:24:05.000000', 'http://localhost:20000/web/media/46846f7d-a072-442c-b542-c879b093c128.jpg',
        'music', 'http://localhost:20000/web/media/ebb9d58a-9824-4113-9280-3830348ba16a.mp3'),

       ('Internet Download Master',
        '2020-03-20 10:24:05.000000', 'http://localhost:20000/web/media/24ff42b4-6d48-4d72-8c35-b7d5d99551d2.jpg',
        'prog', 'http://localhost:20000/web/media/9a6923a9-cd70-47c1-aef3-3338653c82cb.zip');

SELECT id, title, dataandtime, poster, category
FROM post;

SELECT id, title, dataandtime, poster, category
FROM post
WHERE id = ?;

SELECT id, title, dataandtime, poster, category
FROM post
WHERE category = 'TestCategoryMovie';