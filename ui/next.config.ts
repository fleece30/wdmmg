import type { NextConfig } from "next"

const nextConfig: NextConfig = {
  // output: "standalone",
  images: {
    remotePatterns: [new URL("https://4kwallpapers.com/**")],
  },
}

export default nextConfig
