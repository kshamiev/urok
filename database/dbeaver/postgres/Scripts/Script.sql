	SELECT
		i."Id",
		i."Caption",
		COALESCE(i."SignPersonId", i."CreatePersonId") AS "SignPersonId",
		p."FirstName" AS "SignFirstName",
		p."LastName" AS "SignLastName",
		p."Patronymic" AS "SignPatronymic",
		p."Position" AS "SignPosition",
		pa."UrlKey" AS "SignSmallAvatarKey",
		i."CreateDate",
		ir."views" 
	FROM
		instruction AS i
	INNER JOIN person AS p ON
		p."Id" = COALESCE(i."SignPersonId", i."CreatePersonId")
	LEFT JOIN personsavatar pa ON
		pa."PersonId" = p."Id"
	RIGHT JOIN accesstoinstruction a ON
		a."InstructionId" = i."Id"
	RIGHT JOIN instructionrating ir ON
		ir.id = i."Id"
	WHERE
		a."PersonId" = 3289
		AND (a."Deleted" = FALSE OR a."Deleted" ISNULL)
		AND EXISTS(SELECT id FROM instructionrating WHERE id = i."Id")
	ORDER BY
		i."StartDate" DESC, ir."views" DESC
	LIMIT 6
