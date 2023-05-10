import React, { useState } from 'react';
import { useRouter } from 'next/router';
import { AppBar, Toolbar, IconButton, Drawer, List, ListItem, ListItemText, Box } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import { theme } from '@/styles/theme';
import { ThemeProvider } from "@mui/material/styles";
import Link from 'next/link';

const drawerWidth = 240;

interface NavigationItem {
    label: string;
    path: string;
}

const navigationItems: NavigationItem[] = [
    { label: 'Home', path: '/' },
    { label: 'Leagues', path: '/league' },
    // Add more items here as needed
];

export default function Navigation() {
    const router = useRouter();
    const [drawerOpen, setDrawerOpen] = useState(false);

    const handleDrawerToggle = () => {
        setDrawerOpen(!drawerOpen);
    };

    const handleNavItemClick = (path: string) => {
        router.push(path);
        setDrawerOpen(false);
    };

    return (
        <ThemeProvider theme={theme}>
            <Box sx={{ display: 'flex' }}>
                <AppBar position="fixed">
                    <Toolbar>
                        <IconButton edge="start" color="inherit" aria-label="menu" aria-haspopup="true" onClick={handleDrawerToggle}>
                            <MenuIcon />
                        </IconButton>
                        <span>{router.pathname}</span>
                    </Toolbar>
                </AppBar>
                <Drawer
                    sx={{
                        flexShrink: { md: 0 },
                        width: { md: drawerWidth },
                    }}
                    variant="temporary"
                    anchor="left"
                    open={drawerOpen}
                    onClose={handleDrawerToggle}
                >
                    <Box sx={{ overflow: 'auto', width: drawerWidth }}>
                        <List>
                            {navigationItems.map(({ label, path }) => (
                                <ListItem button key={label} onClick={() => handleNavItemClick(path)}>
                                    <ListItemText primary={label} />
                                </ListItem>
                            ))}
                        </List>
                    </Box>
                </Drawer>
            </Box>
        </ThemeProvider>
    );
}
