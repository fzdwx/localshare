/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './**/*.{html,js}',
    ],
    theme: {
        extend: {
            colors: {
                'aura-just': "rgb(97,255,202)",
                'just': "rgb(16,185,129)",
                'just-light': "rgb(52,211,153)",
                'just-lighter': "rgb(110,231,183)",
                'just-dark': "#059669",
                'just-darker': "rgb(4, 120, 87)",
            },
        }
    },
    plugins: [],
}

