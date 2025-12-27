import InputError from '@/components/input-error';
import TextLink from '@/components/text-link';
import { Button } from '@/components/ui/button';
import { Checkbox } from '@/components/ui/checkbox';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { usePageTitle } from '@/hooks/use-page-title';
import AuthLayout from '@/layouts/auth-layout';
import { useForm } from '@inertiajs/react';
import { Eye, EyeOff, LoaderCircle } from 'lucide-react';
import { useState } from 'react';

interface LoginProps {
    status?: string;
    canResetPassword?: boolean;
    errors?: {
        email?: string;
        password?: string;
    };
    old?: {
        email?: string;
    };
}

export default function Login({
    status,
    canResetPassword = false,
    errors = {},
    old = {},
}: LoginProps) {
    usePageTitle('Log in');
    const [showPassword, setShowPassword] = useState(false);

    const { data, setData, post, processing } = useForm({
        email: old.email || '',
        password: '',
        remember: false,
    });

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        post('/login');
    };

    return (
        <AuthLayout
            title="Welcome back"
            description="Enter your email below to login to your account"
        >
            {status && (
                <div className="mb-4 rounded-md border border-green-200 bg-green-50 p-3 text-sm text-green-800 dark:border-green-800 dark:bg-green-900/20 dark:text-green-400">
                    {status}
                </div>
            )}

            <form onSubmit={handleSubmit} className="space-y-5">
                <div>
                    <Label htmlFor="email">Email</Label>
                    <Input
                        id="email"
                        type="email"
                        value={data.email}
                        onChange={(e) => setData('email', e.target.value)}
                        placeholder="you@example.com"
                        autoComplete="email"
                        autoFocus
                        required
                        className={errors.email ? 'border-destructive' : ''}
                    />
                    <InputError message={errors.email} />
                </div>

                <div>
                    <Label htmlFor="password">Password</Label>
                    <div className="relative">
                        <Input
                            id="password"
                            type={showPassword ? 'text' : 'password'}
                            value={data.password}
                            onChange={(e) => setData('password', e.target.value)}
                            autoComplete="current-password"
                            required
                            className={`pr-10 ${errors.password ? 'border-destructive' : ''}`}
                        />
                        <button
                            type="button"
                            onClick={() => setShowPassword(!showPassword)}
                            className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground transition-colors"
                            tabIndex={-1}
                        >
                            {showPassword ? (
                                <EyeOff className="h-4 w-4" />
                            ) : (
                                <Eye className="h-4 w-4" />
                            )}
                        </button>
                    </div>
                    <InputError message={errors.password} />
                </div>

                <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2">
                        <Checkbox
                            id="remember"
                            checked={data.remember}
                            onCheckedChange={(checked) => setData('remember', checked === true)}
                        />
                        <Label htmlFor="remember" className="text-sm font-normal cursor-pointer">
                            Remember me
                        </Label>
                    </div>
                    {canResetPassword && (
                        <TextLink href="/password/request" className="text-sm">
                            Forgot password?
                        </TextLink>
                    )}
                </div>

                <Button type="submit" className="w-full" disabled={processing}>
                    {processing && (
                        <LoaderCircle className="mr-2 h-4 w-4 animate-spin" />
                    )}
                    Sign in
                </Button>

                <div className="text-center text-sm text-muted-foreground">
                    Don't have an account?{' '}
                    <TextLink href="/register">Sign up</TextLink>
                </div>
            </form>
        </AuthLayout>
    );
}
