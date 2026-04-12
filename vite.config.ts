import inertia from '@inertiajs/vite';
import tailwindcss from '@tailwindcss/vite';
import react from '@vitejs/plugin-react';
import path from 'path';
import { defineConfig } from 'vite';

export default defineConfig({
    plugins: [
        inertia({
            pages: 'resources/js/pages',
            ssr: false,
        }),
        react(),
        tailwindcss(),
    ],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './resources/js'),
        },
    },
    build: {
        outDir: 'public/build',
        manifest: true,
        rollupOptions: {
            input: 'resources/js/app.tsx',
        },
    },
    server: {
        port: 5173,
        strictPort: true,
        host: 'localhost',
    },
});
