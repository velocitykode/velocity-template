import { Button, FormField, Input } from '@/components/ui';
import { usePageTitle } from '@/hooks/use-page-title';
import AuthLayout from '@/layouts/auth-layout';
import { useForm } from '@inertiajs/react';
import { ArrowRight, Eye, EyeOff } from 'lucide-react';
import { useState } from 'react';

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
  usePageTitle('Create account');
  const [showPassword, setShowPassword] = useState(false);

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
      description="Get started with Velocity in seconds - no credit card, no ceremony."
      switchPrompt="Already have an account?"
      switchHref="/login"
      switchLabel="Sign in"
    >
      <form onSubmit={handleSubmit} className="space-y-4">
        <FormField label="Name" htmlFor="name" error={errors.name}>
          <Input
            id="name"
            type="text"
            value={data.name}
            onChange={(e) => setData('name', e.target.value)}
            placeholder="Your name"
            autoComplete="name"
            autoFocus
            required
            error={!!errors.name}
          />
        </FormField>

        <FormField label="Email" htmlFor="email" error={errors.email}>
          <Input
            id="email"
            type="email"
            value={data.email}
            onChange={(e) => setData('email', e.target.value)}
            placeholder="you@example.com"
            autoComplete="email"
            required
            error={!!errors.email}
          />
        </FormField>

        <FormField label="Password" htmlFor="password" error={errors.password}>
          <div className="relative">
            <Input
              id="password"
              type={showPassword ? 'text' : 'password'}
              value={data.password}
              onChange={(e) => setData('password', e.target.value)}
              autoComplete="new-password"
              required
              error={!!errors.password}
              className="pr-12"
            />
            <button
              type="button"
              onClick={() => setShowPassword(!showPassword)}
              className="absolute right-3 top-1/2 -translate-y-1/2 text-muted hover:text-fg transition-colors"
              tabIndex={-1}
              aria-label={showPassword ? 'Hide password' : 'Show password'}
            >
              {showPassword ? <EyeOff className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
            </button>
          </div>
        </FormField>

        <FormField
          label="Confirm password"
          htmlFor="password_confirmation"
          error={errors.password_confirmation}
        >
          <Input
            id="password_confirmation"
            type={showPassword ? 'text' : 'password'}
            value={data.password_confirmation}
            onChange={(e) => setData('password_confirmation', e.target.value)}
            autoComplete="new-password"
            required
            error={!!errors.password_confirmation}
          />
        </FormField>

        <Button
          type="submit"
          fullWidth
          size="lg"
          loading={processing}
          iconRight={ArrowRight}
          className="mt-2"
        >
          Create account
        </Button>
      </form>
    </AuthLayout>
  );
}
