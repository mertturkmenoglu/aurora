const FormKitVariants = require('@formkit/themes/tailwindcss')
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [    './tailwind-theme.js',
],
  theme: {
    extend: {},
  },
  plugins: [FormKitVariants],
}

