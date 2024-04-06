-- name: ListCharacters :many
SELECT * FROM characters;

-- name: UpdateCharactersRank :exec
WITH RankedCharacters AS (
    SELECT character_id, ROW_NUMBER() OVER (ORDER BY fame DESC) as calculated_rank
    FROM characters c
    WHERE c.job_id = $1 and c.job_grow_id = $2
)
UPDATE characters c
SET rank = rc.calculated_rank
FROM RankedCharacters rc
WHERE c.character_id = rc.character_id;

-- name: UpdateCharacter :exec
INSERT INTO characters (character_id, server_id, character_name, job_id, job_grow_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(character_id)
DO UPDATE SET
    server_id = $2,
    character_name = $3,
    job_id = $4,
    job_grow_id = $5;