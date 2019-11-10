package data

const dbCreateTableIfDoesNotExistQuery string = `
CREATE TABLE IF NOT EXISTS user_sessions (
	id SERIAL PRIMARY KEY NOT NULL,
	user_id BIGINT NOT NULL,
	session_key CHAR(64) NOT NULL
);`

const dbInsertSessionQuery string = `
INSERT INTO user_sessions (user_id, session_key)
VALUES ($1, $2)`

const dbDeleteSessionQuery string = `
DELETE FROM user_sessions
WHERE user_id = $1 AND session_key = $2`

const dbFindSessionQuery string = `
SELECT EXISTS(
	SELECT user_id, session_key
	FROM user_sessions
	WHERE user_id = $1 AND session_key = $2
);`
