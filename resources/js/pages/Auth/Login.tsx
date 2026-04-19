import { Button, FormField, Input } from '@/components/ui';
import { usePageTitle } from '@/hooks/use-page-title';
import AuthLayout from '@/layouts/auth-layout';
import { Link, useForm } from '@inertiajs/react';
import { ArrowRight, Eye, EyeOff } from 'lucide-react';
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
  usePageTitle('Sign in');
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
      title="Sign in"
      description="Welcome back - manage your application and continue where you left off."
      switchPrompt="New to Velocity?"
      switchHref="/register"
      switchLabel="Create an account"
    >
      {status && (
        <div className="mb-6 border border-emerald-600/40 bg-emerald-600/10 px-4 py-3">
          <p className="caption-mono text-emerald-700 dark:text-emerald-400">{status}</p>
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-4">
        <FormField label="Email" htmlFor="email" error={errors.email}>
          <Input
            id="email"
            type="email"
            value={data.email}
            onChange={(e) => setData('email', e.target.value)}
            placeholder="you@example.com"
            autoComplete="email"
            autoFocus
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
              autoComplete="current-password"
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

        <div className="flex items-center justify-between">
          <label className="flex cursor-pointer items-center gap-2.5">
            <input
              type="checkbox"
              checked={data.remember}
              onChange={(e) => setData('remember', e.target.checked)}
              className="h-4 w-4 rounded-none border-border-strong bg-surface accent-fg"
            />
            <span className="caption-mono text-muted">Remember me</span>
          </label>

          {canResetPassword && (
            <Link
              href="/password/request"
              className="link-underline caption-mono text-muted"
            >
              Trouble logging in?
            </Link>
          )}
        </div>

        <Button
          type="submit"
          fullWidth
          size="lg"
          loading={processing}
          iconRight={ArrowRight}
          className="mt-2"
        >
          Sign in
        </Button>
      </form>
    </AuthLayout>
  );
}
