/** @type {import('tailwindcss').Config} */

import colors from "tailwindcss/colors";

module.exports = {
  darkMode: "class",
  content: ["./src/**/*.{html,ts}"],
  theme: {
    colors: {
      white: colors.white,
      black: colors.black,
      gray: colors.gray,
      yellow: colors.yellow,
      purple: colors.purple,
      green: colors.green,
      vs: "#242424",
    },
    extend: {
      colors: {
        // Light colors
        primary: colors.white,
        secondary: "#565656",

        // Dark colors
        "dark-primary": "#242424",
        "dark-secondary": "#888888",
      },
    },
  },
  plugins: [],
};
