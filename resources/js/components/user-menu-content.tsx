import { router } from '@inertiajs/react';
import { LogOut } from 'lucide-react';
import { User } from '@/types';

interface UserMenuContentProps {
    user?: User;
}

export default function UserMenuContent({ user }: UserMenuContentProps) {
    const handleLogout = () => {
        router.post('/logout');
    };

    return (
        <>
            <button
                onClick={handleLogout}
                className="flex w-full items-center gap-2 px-2 py-1.5 text-sm hover:bg-accent rounded-md text-left"
            >
                <LogOut className="h-4 w-4" />
                Logout
            </button>
        </>
    );
}
