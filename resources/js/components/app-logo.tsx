import { useSidebar } from '@/components/ui/sidebar';

export default function AppLogo() {
    const { state } = useSidebar();
    const isCollapsed = state === 'collapsed';

    if (isCollapsed) {
        return (
            <img
                src="/icon.png"
                alt="Velocity"
                className="h-8"
            />
        );
    }

    return (
        <>
            <img
                src="/logo-light.png"
                alt="Velocity"
                className="h-24 dark:hidden"
            />
            <img
                src="/logo-dark.png"
                alt="Velocity"
                className="hidden h-24 dark:block"
            />
        </>
    );
}
