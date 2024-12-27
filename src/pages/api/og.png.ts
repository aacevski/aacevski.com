import type { APIRoute } from 'astro';
import satori from 'satori';
import sharp from 'sharp';

export const GET: APIRoute = async ({ url }) => {
  const { searchParams } = url;
  const title = searchParams.get('title') || 'Andrej Acevski';
  const description = searchParams.get('description') || 'Software Developer';

  const svg = await satori(
    {
      type: 'div',
      props: {
        style: {
          height: '100%',
          width: '100%',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'flex-start',
          justifyContent: 'center',
          backgroundColor: '#0a0a0a',
          padding: 80,
          backgroundImage: 'radial-gradient(circle at top, rgba(29, 78, 216, 0.15), transparent 80%)',
        },
        children: [
          {
            type: 'div',
            props: {
              style: {
                display: 'flex',
                flexDirection: 'column',
                gap: 24,
              },
              children: [
                {
                  type: 'span',
                  props: {
                    style: {
                      fontSize: 72,
                      fontWeight: 700,
                      fontFamily: 'JetBrains Mono',
                      background: 'linear-gradient(to right, white, rgb(107, 114, 128))',
                      backgroundClip: 'text',
                      color: 'transparent',
                      letterSpacing: '-0.05em',
                      lineHeight: 1.2,
                    },
                    children: title,
                  },
                },
                {
                  type: 'span',
                  props: {
                    style: {
                      fontSize: 32,
                      color: 'rgb(156, 163, 175)',
                      fontFamily: 'Inter',
                    },
                    children: description,
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      width: 1200,
      height: 630,
      fonts: [
        {
          name: 'JetBrains Mono',
          data: await fetch(
            'https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@700&display=swap'
          ).then((res) => res.arrayBuffer()),
          weight: 700,
          style: 'normal',
        },
        {
          name: 'Inter',
          data: await fetch(
            'https://fonts.googleapis.com/css2?family=Inter&display=swap'
          ).then((res) => res.arrayBuffer()),
          weight: 400,
          style: 'normal',
        },
      ],
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