import sitemap from '@astrojs/sitemap'
import vercel from '@astrojs/vercel'
import { defineConfig } from 'astro/config';

import tailwind from '@astrojs/tailwind';
import mdx from '@astrojs/mdx';

export default defineConfig({
  site: 'https://aacevski.com',
  adapter: vercel({
    webAnalytics: {
      enabled: true,
    },
  }),
  output: 'server',
  integrations: [sitemap(), tailwind(), mdx()]
});