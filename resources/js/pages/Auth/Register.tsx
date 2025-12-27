import InputError from '@/components/input-error';
import TextLink from '@/components/text-link';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { usePageTitle } from '@/hooks/use-page-title';
import AuthLayout from '@/layouts/auth-layout';
import { useForm } from '@inertiajs/react';
import { LoaderCircle } from 'lucide-react';

interface RegisterProps {
    errors?: {
        name?: string;
        email?: string;
        password?: string;
        password_confirmation?: string;
    };
    old?: {
        name?: string;
        email?: string;
    };
}

export default function Register({ errors = {}, old = {} }: RegisterProps) {
    usePageTitle('Register');

    const { data, setData, post, processing } = useForm({
        name: old.name || '',
        email: old.email || '',
        password: '',
        password_confirmation: '',
    });

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        post('/register');
    };

    return (
        <AuthLayout
            title="Create account"
            description="Get started with your account"
        >

            <form onSubmit={handleSubmit} className="space-y-4">
                <div className="space-y-2">
                    <Label htmlFor="name">Name</Label>
                    <Input
                        id="name"
                        type="text"
                        value={data.name}
                        onChange={(e) => setData('name', e.target.value)}
                        placeholder="Full name"
                        autoComplete="name"
                        autoFocus
                        required
                    />
                    <InputError message={errors.name} />
                </div>

                <div className="space-y-2">
                    <Label htmlFor="email">Email</Label>
                    <Input
                        id="email"
                        type="email"
                        value={data.email}
                        onChange={(e) => setData('email', e.target.value)}
                        placeholder="name@example.com"
                        autoComplete="email"
                        required
                    />
                    <InputError message={errors.email} />
                </div>

                <div className="space-y-2">
                    <Label htmlFor="password">Password</Label>
                    <Input
                        id="password"
                        type="password"
                        value={data.password}
                        onChange={(e) => setData('password', e.target.value)}
                        autoComplete="new-password"
                        required
                    />
                    <InputError message={errors.password} />
                </div>

                <div className="space-y-2">
                    <Label htmlFor="password_confirmation">
                        Confirm password
                    </Label>
                    <Input
                        id="password_confirmation"
                        type="password"
                        value={data.password_confirmation}
                        onChange={(e) => setData('password_confirmation', e.target.value)}
                        autoComplete="new-password"
                        required
                    />
                    <InputError message={errors.password_confirmation} />
                </div>

                <Button type="submit" className="w-full" disabled={processing}>
                    {processing && (
                        <LoaderCircle className="mr-2 h-4 w-4 animate-spin" />
                    )}
                    Create account
                </Button>

                <p className="text-center text-sm text-muted-foreground">
                    Already have an account?{' '}
                    <TextLink href="/login">Sign in</TextLink>
                </p>
            </form>
        </AuthLayout>
    );
}
