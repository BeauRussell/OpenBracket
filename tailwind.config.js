/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./internal/templates/**/*.html",
		"./static/**/*.js"
	],
	theme: {
		extend: {
			height: {
				'128': '32rem',
			},
		},
	},
	plugins: [],
}

