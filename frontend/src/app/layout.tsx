import type { Metadata } from "next";
import "./globals.css";
import { Sidebar } from "@/components/sidebar";

export const metadata: Metadata = {
    title: "BucketWise",
    description: "Automated financial insights and budgeting",
};

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode;
}) {
    return (
        <html lang="en">
        <body className="min-h-screen flex bg-background text-foreground">
        <Sidebar />
        <main className="flex-1 p-10">{children}</main>
        </body>
        </html>
    );
}
