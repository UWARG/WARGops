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
        creation_date
    )
VALUES
    (
        "1",
        "johndoe",
        "WEEF",
        "waterloo",
        '1998-04-16',
        '1998-04-16',
        "true",
        "john doe",
        "admin@gmail.com",
        '1998-04-16'
    );

INSERT INTO
    transactions (
        id,
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
        "1",
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