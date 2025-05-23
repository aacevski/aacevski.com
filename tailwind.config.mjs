/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}"],
	theme: {
		extend: {
			fontFamily: {
				mono: ["JetBrains Mono", "monospace"],
				sans: ["Inter", "sans-serif"],
			},
			typography: {
				invert: {
					css: {
						"--tw-prose-body": "rgb(209 213 219)",
						"--tw-prose-headings": "rgb(255 255 255)",
						"--tw-prose-links": "rgb(96 165 250)",
						"--tw-prose-bold": "rgb(255 255 255)",
						"--tw-prose-counters": "rgb(209 213 219)",
						"--tw-prose-bullets": "rgb(209 213 219)",
						"--tw-prose-hr": "rgb(75 85 99)",
						"--tw-prose-quotes": "rgb(255 255 255)",
						"--tw-prose-quote-borders": "rgb(75 85 99)",
						"--tw-prose-captions": "rgb(156 163 175)",
						"--tw-prose-code": "rgb(255 255 255)",
						"--tw-prose-pre-code": "rgb(209 213 219)",
						"--tw-prose-pre-bg": "rgb(17 24 39)",
						"--tw-prose-th-borders": "rgb(75 85 99)",
						"--tw-prose-td-borders": "rgb(75 85 99)",
					},
				},
			},
		},
	},
	plugins: [require("@tailwindcss/typography")],
};
