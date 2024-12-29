export function createCoverImage() {
  return {
    type: "div",
    props: {
      style: {
        height: "100%",
        width: "100%",
        display: "flex",
        alignItems: "center",
        backgroundColor: "#0a0a0a",
        padding: 80,
        backgroundImage:
          "radial-gradient(circle at top, rgba(29, 78, 216, 0.15), transparent 80%)",
      },
      children: [
        {
          type: "div",
          props: {
            style: {
              height: 160,
              width: 160,
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              borderRadius: 48,
              marginRight: 56,
              background: "linear-gradient(to top right, #2563eb, #60a5fa)",
              padding: 4,
            },
          },
        },
      ],
    },
  };
}
