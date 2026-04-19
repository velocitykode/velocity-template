import {
  AppLogo,
  Badge,
  IconButton,
  NavItem,
  StatCard,
  ThemeToggle,
  UserAvatar,
} from '@/components/ui';
import { usePageTitle } from '@/hooks/use-page-title';
import { Link, usePage } from '@inertiajs/react';
import {
  Activity,
  ArrowRight,
  BookOpen,
  Check,
  Database,
  Gauge,
  LayoutGrid,
  LogOut,
  Menu,
  Route,
  Server,
  Settings,
  Shield,
  Terminal,
  X,
} from 'lucide-react';
import { useState } from 'react';

const navigation = [
  { name: 'Dashboard', href: '/dashboard', icon: LayoutGrid, active: true },
  { name: 'Routes', href: '#', icon: Route },
  { name: 'Middleware', href: '#', icon: Shield },
  { name: 'Database', href: '#', icon: Database },
];

const resources = [
  {
    name: 'Documentation',
    href: 'https://velocity.velocitykode.com/',
    icon: BookOpen,
    external: true,
  },
  { name: 'Settings', href: '#', icon: Settings },
];

const stats = [
  { label: 'Routes', value: '12', icon: Route, change: '+2' },
  { label: 'Middleware', value: '6', icon: Shield },
  { label: 'Cache hit', value: '94%', icon: Gauge },
  { label: 'Uptime', value: '99.9%', icon: Server },
];

const tasks = [
  { task: 'Configure routes', done: true },
  { task: 'Set up middleware', done: true },
  { task: 'Connect database', done: false },
  { task: 'Configure cache', done: false },
  { task: 'Deploy to production', done: false },
];

const activities = [
  { action: 'Migration completed', time: '2m', icon: Database },
  { action: 'Cache invalidated', time: '15m', icon: Server },
  { action: 'Route registered', time: '1h', icon: Route },
  { action: 'Middleware updated', time: '3h', icon: Shield },
  { action: 'Server restarted', time: '5h', icon: Terminal },
];

export default function Dashboard() {
  usePageTitle('Dashboard');
  const { props } = usePage<{ auth: { user: { name: string; email: string } } }>();
  const user = props.auth?.user;
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const completed = tasks.filter((t) => t.done).length;
  const progressPct = Math.round((completed / tasks.length) * 100);

  return (
    <div className="flex h-screen bg-bg text-fg">
      {/* Mobile overlay - only when drawer open */}
      {sidebarOpen && (
        <div
          className="fixed inset-0 z-40 bg-black/50 lg:hidden"
          onClick={() => setSidebarOpen(false)}
        />
      )}

      {/* Sidebar: mobile drawer <lg, static rail at lg+ */}
      <aside
        className={`fixed inset-y-0 left-0 z-50 flex w-72 flex-col border-r border-border bg-surface transition-transform duration-200 ease-in-out lg:static lg:inset-auto lg:w-64 lg:translate-x-0 lg:transition-none xl:w-72 ${
          sidebarOpen ? 'translate-x-0' : '-translate-x-full'
        }`}
      >
        {/* Header */}
        <div className="flex h-12 items-center justify-between border-b border-border px-4">
          <AppLogo variant="full" size="md" href="/" />
          <div className="flex items-center gap-2">
            <Badge variant="outline">v0.22</Badge>
            <IconButton
              icon={X}
              variant="ghost"
              aria-label="Close menu"
              onClick={() => setSidebarOpen(false)}
              className="lg:hidden"
            />
          </div>
        </div>

        {/* Nav */}
        <nav className="flex-1 overflow-y-auto px-2 py-4">
          <div className="caption-mono-upper mb-2 px-3 text-dim">&gt; Platform</div>
          <div className="space-y-0">
            {navigation.map((item) => (
              <NavItem key={item.name} {...item} onClick={() => setSidebarOpen(false)} />
            ))}
          </div>

          <div className="caption-mono-upper mt-6 mb-2 px-3 text-dim">&gt; Resources</div>
          <div className="space-y-0">
            {resources.map((item) => (
              <NavItem key={item.name} {...item} />
            ))}
          </div>
        </nav>

        {/* Footer: user */}
        <div className="border-t border-border p-3">
          <div className="flex items-center gap-3">
            <UserAvatar name={user?.name} size="md" />
            <div className="min-w-0 flex-1">
              <div className="truncate text-sm font-medium text-fg">
                {user?.name || 'User'}
              </div>
              <div className="truncate caption-mono text-dim">{user?.email}</div>
            </div>
            <Link href="/logout" method="post" as="button">
              <IconButton
                icon={LogOut}
                variant="danger"
                aria-label="Log out"
                as="button"
                type="button"
                size="sm"
              />
            </Link>
          </div>
        </div>
      </aside>

      {/* Main */}
      <div className="flex flex-1 flex-col overflow-hidden">
        {/* Header */}
        <header className="flex h-12 items-center justify-between border-b border-border bg-bg px-4 sm:px-5 lg:px-6">
          <div className="flex items-center gap-3">
            <IconButton
              icon={Menu}
              variant="ghost"
              size="sm"
              aria-label="Open menu"
              onClick={() => setSidebarOpen(true)}
              className="lg:hidden"
            />
            <h1 className="caption-mono-upper text-muted">
              &gt; dashboard
            </h1>
          </div>
          <div className="flex items-center gap-2">
            <ThemeToggle />
          </div>
        </header>

        {/* Content - capped at 1800px for XXL */}
        <main className="flex-1 overflow-y-auto">
          <div className="mx-auto w-full max-w-[1800px] px-4 py-5 sm:px-5 sm:py-6 lg:px-6 xl:px-8 xl:py-7">
            {/* Welcome */}
            <div className="mb-6 lg:mb-7">
              <div className="flex items-center gap-2 mb-3">
                <span className="relative flex h-2 w-2">
                  <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-60" />
                  <span className="relative inline-flex h-2 w-2 rounded-full bg-emerald-500" />
                </span>
                <span className="caption-mono-upper text-emerald-700 dark:text-emerald-400">
                  Live
                </span>
              </div>
              <h2 className="font-sans text-2xl lg:text-3xl xl:text-[36px] font-semibold tracking-tight">
                Welcome back, {user?.name?.split(' ')[0] || 'Developer'}.
              </h2>
              <p className="mt-2 text-[14px] text-muted">
                Your application overview - routes, middleware, cache, uptime.
              </p>
            </div>

            {/* Stats */}
            <div className="mb-6 lg:mb-7 grid gap-px bg-border grid-cols-2 lg:grid-cols-4 border border-border">
              {stats.map((stat) => (
                <StatCard key={stat.label} {...stat} className="border-0" />
              ))}
            </div>

            {/* Two Column */}
            <div className="grid gap-5 lg:gap-6 lg:grid-cols-2">
              {/* Getting Started */}
              <section className="border border-border bg-surface">
                <header className="flex items-center justify-between border-b border-border px-4 py-3 lg:px-5">
                  <div>
                    <div className="caption-mono-upper text-dim">
                      &gt; Getting started
                    </div>
                    <div className="mt-0.5 text-xs text-muted">
                      {completed} of {tasks.length} complete
                    </div>
                  </div>
                  <span className="font-sans text-2xl font-semibold tracking-tight">
                    {progressPct}%
                  </span>
                </header>
                <ul className="divide-y divide-border">
                  {tasks.map((item, i) => (
                    <li
                      key={i}
                      className="flex items-center gap-3 px-4 py-2.5 lg:px-5"
                    >
                      <div
                        className={`flex h-5 w-5 items-center justify-center border ${
                          item.done
                            ? 'border-fg bg-fg text-bg'
                            : 'border-border-strong'
                        }`}
                      >
                        {item.done && <Check className="h-3 w-3" strokeWidth={3} />}
                      </div>
                      <span
                        className={`flex-1 text-sm ${
                          item.done
                            ? 'text-dim line-through'
                            : 'text-fg'
                        }`}
                      >
                        {item.task}
                      </span>
                      <ArrowRight className="h-4 w-4 text-dim" />
                    </li>
                  ))}
                </ul>
                <footer className="border-t border-border p-3">
                  <a
                    href="https://github.com/velocitykode/velocity"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex w-full items-center justify-center gap-2 bg-accent py-2.5 text-sm font-medium tracking-wide text-accent-fg transition-opacity hover:opacity-90"
                  >
                    <BookOpen className="h-4 w-4" />
                    View documentation
                  </a>
                </footer>
              </section>

              {/* Activity */}
              <section className="border border-border bg-surface">
                <header className="flex items-center justify-between border-b border-border px-4 py-3 lg:px-5">
                  <div className="flex items-center gap-3">
                    <Activity className="h-4 w-4 text-muted" />
                    <div>
                      <div className="caption-mono-upper text-dim">
                        &gt; Recent activity
                      </div>
                      <div className="mt-0.5 text-xs text-muted">
                        Live events
                      </div>
                    </div>
                  </div>
                  <span className="flex items-center gap-2">
                    <span className="relative flex h-2 w-2">
                      <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-60" />
                      <span className="relative inline-flex h-2 w-2 rounded-full bg-emerald-500" />
                    </span>
                  </span>
                </header>
                <ul className="divide-y divide-border">
                  {activities.map((item, i) => (
                    <li
                      key={i}
                      className="flex items-center gap-3 px-4 py-2.5 lg:px-5"
                    >
                      <item.icon className="h-4 w-4 flex-shrink-0 text-muted" />
                      <span className="flex-1 truncate text-sm text-fg">
                        {item.action}
                      </span>
                      <span className="caption-mono text-dim">
                        {item.time}
                      </span>
                    </li>
                  ))}
                </ul>
                <footer className="border-t border-border p-3">
                  <button className="w-full text-center caption-mono-upper text-muted hover:text-fg transition-colors py-1">
                    &gt; View all activity
                  </button>
                </footer>
              </section>
            </div>
          </div>
        </main>
      </div>
    </div>
  );
}
