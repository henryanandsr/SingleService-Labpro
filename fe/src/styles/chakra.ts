// chakra theme
import { baseTheme, extendTheme, withDefaultColorScheme } from '@chakra-ui/react';


const theme = extendTheme(
  {
    colors: {
      primary: baseTheme.colors.telegram, // Change this to change the primary color
    },    
  },
  withDefaultColorScheme({ colorScheme: 'primary' })
);

export default theme;