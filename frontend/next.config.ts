import process from "process";

const backend = process.env.BACKEND?.replace(/\/$/, '');
const prefix = process.env.PREFIX || '';

console.log('BACKEND:', backend);
console.log('PREFIX:', prefix);

const nextConfig = {
  async rewrites() {
    return [
      {
        source: `${process.env.PREFIX}/:path*`,
        destination: `${backend}${prefix}/:path*`,
      }
    ]
  },
  output: 'standalone',
};

export default nextConfig;
