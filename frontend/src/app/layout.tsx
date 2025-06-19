import Navbar from "@/components/layout/Navbar";
import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
    title: "TourVista",
    description: "Plan and explore epic tours with ease",
};

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <html lang="en">
            <body>
                <Navbar /> {/* Global navigation bar */}
                <main>{children}</main>
            </body>
        </html>
    );
}
