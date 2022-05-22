INSERT INTO
    d_info(
        d_id,
        d_username,
        avatar,
        d_discriminator,
        public_flags,
        flags,
        email,
        verified,
        locale,
        mfa_enabled,
        create_time,
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
        $7,
        $8,
        $9,
        $10,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ) ON CONFLICT ON CONSTRAINT d_info_pkc DO
UPDATE
SET
    d_username = $2,
    avatar = $3,
    d_discriminator = $4,
    public_flags = $5,
    flags = $6,
    email = $7,
    verified = $8,
    locale = $9,
    mfa_enabled = $10,
    update_time = CURRENT_TIMESTAMP;