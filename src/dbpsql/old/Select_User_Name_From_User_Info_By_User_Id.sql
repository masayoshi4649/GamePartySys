SELECT
    user_info.user_name
FROM
    user_info
WHERE
    user_info.user_id = [$1]