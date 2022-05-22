INSERT INTO
    d_token(
        d_id,
        access_token,
        expires_in,
        refresh_token,
        scope,
        token_type,
        update_time
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        CURRENT_TIMESTAMP
    ) ON CONFLICT ON CONSTRAINT d_token_pkc DO
UPDATE
SET
    access_token = $2,
    expires_in = $3,
    refresh_token = $4,
    scope = $5,
    token_type = $6,
    update_time = CURRENT_TIMESTAMP;