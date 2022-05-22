SELECT
    instance_info.instance_id,
    mst_game.game_name AS game_type_name,
    instance_info.play_style,
    instance_info.max_join,
    instance_info.recruit_title,
    instance_info.recruit_place,
    instance_info.recruit_comment,
    to_char(instance_info.recruit_due, 'YYYY/MM/DD hh24:mi') AS recruit_due,
    to_char(instance_info.meeting_time, 'YYYY/MM/DD hh24:mi') AS meeting_time,
    (
        case
            WHEN instance_info.join_pass IS NULL THEN FALSE
            WHEN instance_info.join_pass = '' THEN FALSE
            ELSE TRUE
        END
    ) AS join_pass_exist,
    instance_info.recruit_type,
    mst_recruit.recruit_type_name,
    user_info.user_name AS create_user_name
FROM
    instance_info
    INNER JOIN mst_game ON mst_game.game_type = instance_info.game_type
    INNER JOIN user_info ON user_info.user_id = instance_info.create_user
    INNER JOIN mst_recruit ON mst_recruit.recruit_type = instance_info.recruit_type
ORDER BY
    instance_info.recruit_due ASC