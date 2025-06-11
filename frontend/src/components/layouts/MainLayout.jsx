import React from 'react';
import Header from './Header';
import Footer from './Footer';

const MainLayout = ({ children }) => {
    return (
        <div>
            <Header />
            <main className="min-h-screen bg-foreground text-surface-light">
                {children}
            </main>
            <Footer />
        </div>
    );
}

export default MainLayout;
