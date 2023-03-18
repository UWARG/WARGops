export const moneyToInt = (value: string): number => {
    const money: string = value.replace(/,/g, "").replace(/\./g, "");
    return parseInt(money);
};

export const intToMoney = (value: number): string => {

    if (value.toString().length == 1)
        return `${value < 0 ? "-" : ""}0.0${value}`;
    else if (value.toString().length == 2)
        return `${value < 0 ? "-" : ""}0.${value}`;
    else {
        const dollars = value >= 100 || value <= -100 ? value.toString().substring(0, value.toString().length - 2).replace(/\B(?=(\d{3})+(?!\d))/g, ",") : "0";
        const cents = value.toString().substring(value.toString().length - 2);
        return `${dollars}.${cents}`;

    }
};