DROP FUNCTION IF EXISTS current_tenant_id;

CREATE FUNCTION current_tenant_id()
RETURNS BIGINT
DETERMINISTIC
RETURN @current_tenant_id;

CREATE OR REPLACE
VIEW users_view
AS
SELECT *
  FROM users
 WHERE tenant_id = current_tenant_id()
WITH CHECK OPTION;