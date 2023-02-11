import anime from 'animejs';

export const alertUP = (element: HTMLElement): void => {
    anime({
        targets: element,
        delay: 0,
        // translateY: 50,
        opacity: 1,
    });
};

export const alertDOWN = (element: HTMLElement): void => {
    anime({
        targets: element,
        delay: 0,
        // translateY: -50,
        opacity: 0,
        
    });
}
