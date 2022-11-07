package query_consts

const (
	INSERT_INTO_TBL_VALUES_VAL        = " INSERT INTO %s VALUES %s ON CONFLICT DO NOTHING"
	SELECT_COL_FROM_TBL_WHERE_CND     = " SELECT %s FROM %s WHERE %s"
	SELECT_COL_FROM_TBL               = " SELECT %s FROM %s"
	UPDATE_TBL_SET_VAL_WHERE_CND      = " UPDATE %s SET %s WHERE %s"
	DELETE_FROM_TBL_WHERE_CND         = " DELETE FROM %s WHERE %s"
	DELETE_CASCADE_FROM_TBL_WHERE_CND = " DELETE CASCADE FROM %s WHERE %s"

	SELECT_COL1_FROM_TBL1_WHERE_CND1_SELECT_COL2_FROM_TBL2_CND2 = ` 
	SELECT %s FROM %s
	WHERE %s=(` + SELECT_COL_FROM_TBL_WHERE_CND + `)
	`

	PAGINATION_COL_TBL_CND_PAG_TBL_PAG_LIM = `
	SELECT %s
	FROM %s
	WHERE %s > (CASE WHEN %d > 0 THEN
				   (SELECT max(id)
					FROM (SELECT id
						  FROM %s
						  LIMIT %d)
					as foo)
				ELSE
					-1
				END)
	LIMIT %d;
	`

	SELECT_VIDEO_PATH_BETWEEN = `
	SELECT CONCAT("pathRecord", '/', "fileName")
	FROM "VideoRecord"
	WHERE "pathStream"='%s' 
	AND "recordTime" BETWEEN 
		(
			SELECT "recordTime"
			FROM "VideoRecord"
			WHERE "pathStream"='%s'
			AND "recordTime" <= '%s'::timestamp
			ORDER BY "recordTime"  DESC
			LIMIT 1
		) 
			AND 
		(
			SELECT "recordTime" 
			FROM "VideoRecord"
			WHERE "pathStream"='%s' 
			AND "recordTime" >= '%s'::timestamp + (interval '%dm')
			ORDER BY "recordTime"  ASC
			LIMIT 1
		)
	`
)
