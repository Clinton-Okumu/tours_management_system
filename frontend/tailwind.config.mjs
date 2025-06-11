// tailwind.config.js or .mjs
export default {
  theme: {
    extend: {
      colors: {
        // Semantic color roles
        primary: {
          DEFAULT: "#2563eb", // clint-blue 600
          light: "#60a5fa", // clint-blue 400
          dark: "#1d4ed8", // clint-blue 700
        },
        surface: {
          light: "#f1f5f9", // clint-dark 100
          DEFAULT: "#cbd5e1", // clint-dark 300
          dark: "#334155", // clint-dark 700
        },
        muted: "#94a3b8", // clint-dark 400
        base: "#1e293b", // clint-dark 800
        foreground: "#0f172a", // clint-dark 900
      },
    },
  },
  content: ["./src/**/*.{js,ts,jsx,tsx}"],
};
