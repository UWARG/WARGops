export const moneyToInt = (value: string): number => {
    const money: string = value.replace(/,/g, "").replace(/\./g, "");
    return parseInt(money);
};

export const intToMoney = (value: number): string => {
    return value.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
};