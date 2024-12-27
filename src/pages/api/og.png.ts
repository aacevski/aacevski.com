import type { APIRoute } from 'astro';
import satori from 'satori';
import sharp from 'sharp';
import { createOGImage } from '../../utils/og-image';
import { loadFonts } from '../../utils/fonts';

export const GET: APIRoute = async ({ url }) => {
  const { searchParams } = url;
  const title = searchParams.get('title') || 'Andrej Acevski';
  const description = searchParams.get('description') || 'Software Developer';

  const svg = await satori(
    createOGImage({ title, description }),
    {
      width: 1200,
      height: 630,
      fonts: loadFonts(),
    }
  );

  const png = await sharp(Buffer.from(svg))
    .png()
    .toBuffer();

  return new Response(png, {
    headers: {
      'Content-Type': 'image/png',
      'Cache-Control': 'public, max-age=31536000, immutable',
    },
  });
};