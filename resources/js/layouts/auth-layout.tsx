import { type PropsWithChildren } from 'react';

interface AuthLayoutProps {
    title: string;
    description: string;
}

export default function AuthLayout({
    title,
    description,
    children,
}: PropsWithChildren<AuthLayoutProps>) {
    return (
        <div className="flex h-screen items-center justify-center bg-muted/30 p-4">
            <div className="w-full max-w-md">
                <div className="mb-8 text-center">
                    <div className="mx-auto h-24 w-48 mb-4">
                        <img
                            src="/logo-light.png"
                            alt="Velocity"
                            className="w-full object-cover object-top dark:hidden"
                        />
                        <img
                            src="/logo-dark.png"
                            alt="Velocity"
                            className="w-full object-cover object-top hidden dark:block"
                        />
                    </div>
                    <h1 className="text-2xl font-semibold">{title}</h1>
                    <p className="mt-2 text-sm text-muted-foreground">
                        {description}
                    </p>
                </div>
                <div className="rounded-lg border bg-card p-8 shadow-sm">
                    {children}
                </div>
            </div>
        </div>
    );
}
