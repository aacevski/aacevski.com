export const startAnimation = {
  show: {
    y: 0,
    opacity: 1,
    transition: {
      type: 'tween',
      duration: 0.5,
      ease: 'easeInOut',
    },
  },
  hide: {
    y: 20,
    opacity: 0,
    transition: {
      type: 'tween',
      duration: 0.5,
      ease: 'easeInOut',
    },
  },
};
