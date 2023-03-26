INSERT INTO
    accounts (
        id,
        waterloo_id,
        name,
        source,
        allocation_date,
        expiry_date,
        active,
        creator,
        point_of_contact,
        creation_date,
        term
    )
VALUES
    (
        "e8ccb501-4537-4d92-a770-fc7a8caa7f46",
        "johndoe",
        "WEEF",
        "waterloo",
        '1998-04-16',
        '1998-04-16',
        "true",
        "john doe",
        "admin@gmail.com",
        '1998-04-16',
        "Fall 2020"
    );

INSERT INTO
    transactions (
        id,
        name,
        account_id,
        creator,
        type,
        ref,
        status,
        amount,
        approval_date,
        approved_by,
        payment_date,
        creation_date,
        rejected_date,
        notes
    )
VALUES
    (
        "1",
        "New Kicks 👟",
        "e8ccb501-4537-4d92-a770-fc7a8caa7f46",
        "john doe",
        1,
        "ref",
        1,
        100,
        '1998-04-16',
        'john doe2',
        '1998-04-16',
        '1998-04-16',
        '1998-04-16',
        "notes"
    );

INSERT INTO
    transactions (
        id,
        name,
        account_id,
        creator,
        type,
        ref,
        status,
        amount,
        approval_date,
        approved_by,
        payment_date,
        creation_date,
        rejected_date,
        notes
    )
VALUES
    (
        "2",
        "Expensive Circuit Board",
        "e8ccb501-4537-4d92-a770-fc7a8caa7f46",
        "john doe 2",
        0,
        "ref",
        0,
        100,
        '1998-04-16',
        'john doe2',
        '1998-04-16',
        '1998-04-16',
        '1998-04-16',
        "notes"
    );