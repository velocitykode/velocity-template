import inertia from '@inertiajs/vite';
import tailwindcss from '@tailwindcss/vite';
import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';
import velocity from '@velocitykode/velocity-vite-plugin';

// Vite+ (vp) options; not part of vite's UserConfig type.
const vitePlus = {
    fmt: {
        singleQuote: true,
        tabWidth: 4,
    },
    check: {
        fmt: false,
    },
};

export default {
    ...defineConfig({
        plugins: [
            velocity('resources/js/app.tsx'),
            inertia(),
            react(),
            tailwindcss(),
        ],
        server: {
            port: 5173,
            strictPort: true,
            host: 'localhost',
        },
    }),
    ...vitePlus,
};
