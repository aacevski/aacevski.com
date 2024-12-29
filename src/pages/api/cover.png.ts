import type { APIRoute } from "astro";
import satori from "satori";
import sharp from "sharp";
import { loadFonts } from "../../utils/fonts";
import { createCoverImage } from "../../utils/cover-image";

export const GET: APIRoute = async ({ url }) => {
  const { searchParams } = url;
  const width = Number(searchParams.get("width")) || 1200;
  const height = Number(searchParams.get("height")) || 630;

  const svg = await satori(createCoverImage(), {
    width,
    height,
    fonts: loadFonts(),
  });

  const png = await sharp(Buffer.from(svg)).png().toBuffer();

  return new Response(png, {
    headers: {
      "Content-Type": "image/png",
      "Cache-Control": "public, max-age=31536000, immutable",
    },
  });
};
