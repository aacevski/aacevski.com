import rss from '@astrojs/rss'
import { getCollection } from 'astro:content'
import { SITE_TITLE, SITE_DESCRIPTION } from '../consts'

export async function GET(context) {
  const posts = await getCollection('blog')
  return rss({
    title: "Andrej Acevski - Software Developer",
    description: "Andrej Acevski's personal website and blog",
    site: context.site,
    items: posts.map((post) => ({
      ...post.data,
      link: `/blog/${post.slug}/`,
    })),
  })
}