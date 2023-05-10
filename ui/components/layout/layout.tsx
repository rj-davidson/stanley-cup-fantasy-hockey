import React, { ReactNode } from 'react';
import Navigation from "@/components/layout/navigation";
import Header from "@/components/layout/header";
import Body from "@/components/layout/body";
import Footer from "@/components/layout/footer";

interface LayoutProps {
    children: ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
    return (
        <>
            <Navigation />
            <Header />
            <Body>
                {children}
            </Body>
            <Footer />
        </>
    );
};

export default Layout;
