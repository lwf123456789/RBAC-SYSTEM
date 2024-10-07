/** @type {import('next').NextConfig} */
const nextConfig = {
    images: {
        unoptimized: true,
    },
    reactStrictMode: false,
    // async rewrites() {
    //     return [
    //         {
    //             source: "/v1/:path*",
    //             destination: "http://192.168.110.2:8082/v1/:path*",
    //         },
    //     ];
    // },
    // webpack(config) {
    //     return config;
    // }
};

export default nextConfig;
