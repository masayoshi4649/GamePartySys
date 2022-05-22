SELECT
    instance_info.instance_id,
    instance_info.recruit_type,
    mst_recruit.recruit_type_name,
    instance_info.recruit_title,
    instance_info.recruit_place,
    instance_info.max_join,
    instance_info.min_join,
    instance_info.recruit_due,
    (
        CASE
            WHEN instance_info.join_pass = '' THEN FALSE
            WHEN instance_info.join_pass IS NULL THEN FALSE
            ELSE TRUE
        END
    ) AS join_pass_exist,
    instance_info.recruit_comment_must,
    instance_info.recruit_comment,
    instance_info.play_style,
    d_info.d_username || '#' || (
        CASE
            WHEN user_config.d_discriminator_display = TRUE THEN d_info.d_discriminator
            ELSE '????'
        END
    ) AS display_user_name,
    'https://cdn.discordapp.com/avatars/' || d_info.d_id || '/' || d_info.avatar AS display_user_icon,
    instance_info.create_time
FROM
    instance_info
    INNER JOIN mst_recruit ON mst_recruit.recruit_type = instance_info.recruit_type
    INNER JOIN d_info ON d_info.d_id = instance_info.d_id
    LEFT OUTER JOIN user_config ON user_config.d_id = instance_info.d_id
WHERE
    instance_info.recruit_due > CURRENT_TIMESTAMP
ORDER BY
    instance_info.recruit_due ASC,
    instance_info.create_time ASC