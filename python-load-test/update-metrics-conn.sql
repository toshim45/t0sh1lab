INSERT INTO metrics.connections
SELECT 
CURRENT_TIMESTAMP AS id,
COUNT(*) FILTER(WHERE state='active') AS active,
COUNT(*) FILTER(WHERE state='idle') AS idle,
COUNT(*) FILTER(WHERE state='idle in transaction') AS idle_in_txn
FROM pg_stat_activity
WHERE datname IS NOT NULL
AND usename = 'svchasura';
