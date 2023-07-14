import  { generateClasses } from '@formkit/themes'
import myTailwindTheme from './tailwind-theme.js' // change to your theme's path

export default {
  config: {
    classes: generateClasses(myTailwindTheme),
  },
}