export const rules = {
    required: (value: string) => !!value || 'Required.',
    waterlooId: (value: string) => {
        const waterlooIdRegex = /^[a-z]{2}\d{8}$/i;
        return waterlooIdRegex.test(value) || 'Invalid Waterloo ID.';
    },
    
    email: (value: string) => {
        const emailRegex = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/;
        return emailRegex.test(value) || 'Invalid e-mail.';
    }
};