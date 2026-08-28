import inertia from '@inertiajs/vite';
import tailwindcss from '@tailwindcss/vite';
import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite-plus';
import velocity from '@velocitykode/velocity-vite-plugin';

export default defineConfig({
    fmt: {
        singleQuote: true,
        tabWidth: 4,
    },
    check: {
        fmt: false,
    },
    plugins: [
        velocity('resources/js/app.tsx'),
        inertia({
            pages: 'resources/js/pages',
        }),
        react(),
        tailwindcss(),
    ],
    server: {
        port: 5173,
        strictPort: true,
        host: 'localhost',
    },
});
