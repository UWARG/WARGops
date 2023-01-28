const converType = (type: string) => {
    switch (type) {
        case '0':
            return 'Deposit';
        case '1':
            return 'Rembursment';
        case '2':
            return 'Procurement';
        default:
            return 'Unkonwn';
    }
};