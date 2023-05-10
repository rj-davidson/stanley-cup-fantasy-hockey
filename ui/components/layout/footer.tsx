import React from 'react';

const Footer: React.FC = () => {
    const currentYear = new Date().getFullYear();
    return (
        <footer>
            <p>&copy; {currentYear} Fantasy Hockey. All rights reserved.</p>
        </footer>
    );
};

export default Footer;
