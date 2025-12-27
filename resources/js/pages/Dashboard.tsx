import AppLayout from '@/layouts/app-layout';
import { type BreadcrumbItem } from '@/types';
import { usePageTitle } from '@/hooks/use-page-title';

const breadcrumbs: BreadcrumbItem[] = [
    {
        title: 'Dashboard',
        href: '/dashboard',
    },
];

export default function Dashboard() {
    usePageTitle('Dashboard');

    return (
        <AppLayout breadcrumbs={breadcrumbs}>
            <div className="flex flex-1 flex-col gap-4 p-4">
                <div className="grid gap-4 md:grid-cols-3">
                    <div className="rounded-lg border bg-card p-6 min-h-32"></div>
                    <div className="rounded-lg border bg-card p-6 min-h-32"></div>
                    <div className="rounded-lg border bg-card p-6 min-h-32"></div>
                </div>

                <div className="rounded-lg border bg-card p-8">
                    <h2 className="text-xl font-semibold">Welcome to Velocity</h2>
                    <p className="mt-2 text-sm text-muted-foreground">
                        Start building your application with Velocity. Check out the documentation to learn more.
                    </p>
                </div>
            </div>
        </AppLayout>
    );
}
