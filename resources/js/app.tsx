import '../css/app.css';

import { createInertiaApp, http, router } from '@inertiajs/react';
import { initializeTheme } from './hooks/use-appearance';

let csrfToken: string | null = readInitialCsrfToken();

http.onRequest((config) => {
    if (csrfToken) {
        config.headers = { ...config.headers, 'X-CSRF-Token': csrfToken };
    }
    return config;
});

router.on('navigate', (event) => {
    const pageProps = event.detail.page.props as { csrf_token?: string };
    if (pageProps.csrf_token) {
        csrfToken = pageProps.csrf_token;
        const meta = document.querySelector('meta[name="csrf-token"]');
        if (meta) {
            meta.setAttribute('content', csrfToken);
        }
    }
});

createInertiaApp({
    progress: {
        color: '#4B5563',
    },
});

initializeTheme();

function readInitialCsrfToken(): string | null {
    const el = document.getElementById('app');
    const raw = el?.dataset.page;
    if (!raw) return null;
    try {
        const page = JSON.parse(raw) as { props?: { csrf_token?: string } };
        return page.props?.csrf_token ?? null;
    } catch {
        return null;
    }
}
