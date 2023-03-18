export const rules = {
  required: (value: string) => !!value || "Required.",
  waterlooId: (value: string) => {
    const waterlooIdRegex = /^[a-z]{2}\d{8}$/i;
    return waterlooIdRegex.test(value) || "Invalid Waterloo ID.";
  },

  email: (value: string) => {
    const emailRegex =
      /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/;
    return emailRegex.test(value) || "Invalid e-mail.";
  },
  money: (value: string) => {
    const moneyRegex1 = /^(\d{1,3})(,\d{1,3})*(\.\d{1,})?$/;
    const moneyRegex2 = /^(\d{1,3})(,\d{1,3})*$/;
    const moneyRegex3 = /^\d+(\.\d{1,2})?$/;
    if (moneyRegex1.test(value) || moneyRegex2.test(value) || moneyRegex3.test(value))
      return true;
    else
      return "Invalid money format. (e.g. 1,000.00 or 1,000 )";
  },
};
