import type { APIRoute } from "astro";
import satori from "satori";
import sharp from "sharp";
import { loadFonts } from "../../utils/fonts";
import { loadAvatarImage } from "../../utils/image";

export const GET: APIRoute = async ({ url }) => {
  const { searchParams } = url;
  const title = searchParams.get("title") || "Andrej Acevski";
  const description = searchParams.get("description") || "Software Developer";
  const avatarImage = loadAvatarImage();

  const svg = await satori(createInstagramStoryImage({ title, description, avatarImage }), {
    width: 1080,
    height: 1920,
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

function createInstagramStoryImage({
  title,
  description,
  avatarImage,
}: {
  title: string;
  description: string;
  avatarImage: string;
}) {
  return {
    type: "div",
    props: {
      style: {
        height: "100%",
        width: "100%",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        backgroundColor: "#0a0a0a",
        padding: 80,
        backgroundImage:
          "radial-gradient(circle at center, rgba(29, 78, 216, 0.15), transparent 70%)",
      },
      children: [
        {
          type: "div",
          props: {
            style: {
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              gap: 60,
              maxWidth: 900,
            },
            children: [
              {
                type: "div",
                props: {
                  style: {
                    height: 200,
                    width: 200,
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                    borderRadius: 60,
                    background: "linear-gradient(to top right, #2563eb, #60a5fa)",
                    padding: 5,
                  },
                  children: [
                    {
                      type: "img",
                      props: {
                        src: avatarImage,
                        style: {
                          height: "100%",
                          width: "100%",
                          objectFit: "cover",
                          borderRadius: 55,
                        },
                        width: 200,
                        height: 200,
                      },
                    },
                  ],
                },
              },
              {
                type: "div",
                props: {
                  style: {
                    display: "flex",
                    flexDirection: "column",
                    alignItems: "center",
                    gap: 24,
                    textAlign: "center",
                  },
                  children: [
                    {
                      type: "span",
                      props: {
                        style: {
                          fontSize: 72,
                          fontWeight: 700,
                          fontFamily: "JetBrains Mono",
                          background:
                            "linear-gradient(to right, white, rgb(107, 114, 128))",
                          backgroundClip: "text",
                          color: "transparent",
                          letterSpacing: "-0.05em",
                          lineHeight: 1.2,
                          textAlign: "center",
                          maxWidth: "100%",
                          display: "block",
                        },
                        children: title,
                      },
                    },
                    {
                      type: "span",
                      props: {
                        style: {
                          fontSize: 36,
                          color: "rgb(156, 163, 175)",
                          fontFamily: "Inter",
                          textAlign: "center",
                          maxWidth: "90%",
                          display: "block",
                        },
                        children: description,
                      },
                    },
                  ],
                },
              },
              {
                type: "div",
                props: {
                  style: {
                    marginTop: 80,
                    padding: "16px 32px",
                    borderRadius: 40,
                    backgroundColor: "rgba(255, 255, 255, 0.1)",
                    border: "1px solid rgba(255, 255, 255, 0.2)",
                    display: "flex",
                    justifyContent: "center",
                  },
                  children: [
                    {
                      type: "span",
                      props: {
                        style: {
                          fontSize: 24,
                          color: "rgb(156, 163, 175)",
                          fontFamily: "Inter",
                          display: "block",
                        },
                        children: "aacevski.com",
                      },
                    },
                  ],
                },
              },
            ],
          },
        },
      ],
    },
  };
}