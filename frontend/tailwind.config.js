/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './public/**/*.html',
    './src/**/*.{vue,js,ts}',
  ],
  theme: {
    extend: {
      screens: {
        "xl": "1380px",
        "lg": "1024px",
        "md": "768px",
        "sm": "640px",
        "xs": "480px",
        "2xs": "320px",
      },
      colors: {
        "ink": {
          DEFAULT: "#25180f",
        },
        "gold": {
          DEFAULT: "#d7ad63",
          "light": "#ebdec4",
          "dark": "#624b25",
        },
        "ember": {
          DEFAULT: "#8f3c24",
          "light": "#c7654d",
          "dark": "#391b13",
        },
      }
    }
  },
  plugins: [],
}
