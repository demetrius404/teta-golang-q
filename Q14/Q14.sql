--- Нечетко сформулированная задача / вопрос: Напишите sql запросы для создания индексов по этим таблицам: ...
--- Какая целевая база данных и соответсвенно диалект SQL?
--- Какие колонки? — одна или несколько (зависит от условий использования)?
--- Тип индекса? — зависит от условий использования / запроса?
--- и т.д. - оставляю на свое усмотрение
--- PostgreSQL
--- CREATE INDEX ...
--- CREATE INDEX CONCURRENTLY ...
--- CREATE INDEX IF NOT EXISTS ...
--- ref: https://www.postgresql.org/docs/current/sql-createindex.html
CREATE INDEX idx_orders_id ON orders USING btree(id);
CREATE INDEX idx_orders_created_at ON orders WITH (fillfactor = 70);
CREATE INDEX idx_users_id ON users USING btree(id);
CREATE INDEX idx_users_email ON users USING gist(email);
--- Function (language: plpgsql) ---
--- CREATE FUNCTION OR REPLACE ...
DROP FUNCTION IF EXISTS create_btree_index(TEXT, TEXT, TEXT);
CREATE FUNCTION create_btree_index(_index TEXT, _table TEXT, _column TEXT) RETURNS void AS
$$
DECLARE 
idx_count INTEGER;
BEGIN
SELECT count(*) INTO idx_count
FROM pg_indexes
WHERE schemaname = 'public'
AND tablename = lower(_table)
AND indexname = lower(_index);
IF idx_count = 0 THEN
EXECUTE format('CREATE INDEX %I ON %I USING btree(%s)', _index, _table, _column);
END IF;
END;
$$
LANGUAGE plpgsql;
--- DML (Data Manipulation Language) --- EXECUTE FUNCTION ---
SELECT create_btree_index('idx_carts_id', 'carts', 'id');
