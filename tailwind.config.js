/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./internal/templates/**/*.html",
		"./static/**/*.js"
	],
	darkMode: 'class',
	theme: {
		extend: {
			colors: {
				background: 'var(--color-background)',
				text: 'var(--color-text)',
				card: 'var(--color-card)',
				primary: 'var(--color-primary)',
				secondary: 'var(--color-secondary)',
			},
		},
	},
	plugins: [],
}

