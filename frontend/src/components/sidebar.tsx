"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

const routes = [
    { label: "Dashboard", href: "/" },
    { label: "Categories", href: "/categories" },
    { label: "Transactions", href: "/transactions" },
];

export function Sidebar() {
    const pathname = usePathname();

    return (
        <aside className="w-64 min-h-screen p-6 bg-sidebar text-sidebar-foreground border-r border-sidebar-border">
            <h1 className="text-2xl font-bold mb-8">BucketWise</h1>

            <nav className="space-y-2">
                {routes.map((route) => {
                    const active = pathname === route.href;

                    return (
                        <Link
                            key={route.href}
                            href={route.href}
                            className={`
                block px-3 py-2 rounded-md text-sm font-medium
                ${
                                active
                                    ? "bg-sidebar-accent text-sidebar-accent-foreground"
                                    : "hover:bg-sidebar-accent"
                            }
              `}
                        >
                            {route.label}
                        </Link>
                    );
                })}
            </nav>
        </aside>
    );
}
