/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './**/*.{html,js}',
    ],
    theme: {
        extend: {
            colors: {
                'aura-just': "#61ffca",
                'just': "#10b981",
                'just-light': "#34d399",
                'just-lighter': "#6ee7b7",
                'just-dark': "#059669",
                'just-darker': "rgb(4, 120, 87)",
            },
        }
    },
    plugins: [],
}

