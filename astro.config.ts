import sitemap from "@astrojs/sitemap";
import vercel from "@astrojs/vercel";
import { defineConfig } from "astro/config";

import mdx from "@astrojs/mdx";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
	site: "https://aacevski.com",
	adapter: vercel({
		webAnalytics: {
			enabled: true,
		},
	}),
	output: "server",
	vite: {
		plugins: [tailwindcss()],
	},
	integrations: [sitemap(), mdx()],
});
