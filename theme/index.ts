import { extendTheme } from '@chakra-ui/react';
import { mode } from '@chakra-ui/theme-tools';

const theme = extendTheme({
  fonts: {
    heading: 'Inter',
    body: 'Inter',
  },
  styles: {
    global: (props) => ({
      body: {
        bg: mode('#F6F6F6', '#0B0B0C')(props),
      },
    }),
  },
});

export default theme;
