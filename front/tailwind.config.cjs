/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        "primary": "#3b82f6",
        "warg-blue": "#081b2e",
        "warg-gray": "#f9fafb",
        "warg-accent": "#FCE2D3",
      },
    },
  },
  plugins: [],
}
