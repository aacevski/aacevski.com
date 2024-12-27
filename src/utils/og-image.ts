interface OGImageProps {
  title: string;
  description: string;
  avatarImage: string;
}

export function createOGImage({ title, description, avatarImage }: OGImageProps) {
  return {
    type: 'div',
    props: {
      style: {
        height: '100%',
        width: '100%',
        display: 'flex',
        alignItems: 'center',
        backgroundColor: '#0a0a0a',
        padding: 80,
        backgroundImage: 'radial-gradient(circle at top, rgba(29, 78, 216, 0.15), transparent 80%)',
      },
      children: [
        {
          type: 'div',
          props: {
            style: {
              height: 160,
              width: 160,
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              borderRadius: 48,
              marginRight: 56,
              background: 'linear-gradient(to top right, #2563eb, #60a5fa)',
              padding: 4,
            },
            children: [
              {
                type: 'img',
                props: {
                  src: avatarImage,
                  style: {
                    height: '100%',
                    width: '100%',
                    objectFit: 'cover',
                    borderRadius: 44,
                  },
                  width: 160,
                  height: 160,
                },
              },
            ],
          },
        },
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
  };
}
