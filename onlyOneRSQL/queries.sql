-- Section1
-- your 1st query here
DELETE FROM users
WHERE family NOT LIKE '%m%'
    AND family NOT LIKE '%d%';
-- Section2
-- your 2nd query here
Delete from users
where family = 'mohammadi'
    OR salary IN (3801, 7414, 2885, 9701, 7356);
-- Section3
-- your 3rd query here
DELETE FROM users
WHERE family = 'booazar'
    OR birth_date BETWEEN '1995-01-01' AND '2000-12-31';