import inertia from '@inertiajs/vite';
import tailwindcss from '@tailwindcss/vite';
import react from '@vitejs/plugin-react';
import { defineConfig, lazyPlugins } from 'vite-plus';
import velocity from '@velocitykode/velocity-vite-plugin';

export default defineConfig({
    plugins: lazyPlugins(() => [
        velocity('resources/js/app.tsx'),
        inertia(),
        react(),
        tailwindcss(),
    ]),
    server: {
        port: 5173,
        strictPort: true,
        host: 'localhost',
    },
    fmt: {
        singleQuote: true,
        tabWidth: 4,
        sortTailwindcss: {
            functions: ['clsx', 'cn', 'cva'],
            entryPoint: 'resources/css/app.css',
        },
    },
    check: {
        fmt: false,
    },
});
