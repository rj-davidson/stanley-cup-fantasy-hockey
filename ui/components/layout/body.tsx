import { ReactNode } from 'react';
import { makeStyles } from '@mui/styles';
import { Container } from '@mui/material';

const useStyles = makeStyles(() => ({
    root: {
        flexGrow: 1,
        paddingTop: '20px',
        paddingBottom: '20px',
    },
}));

interface Props {
    children: ReactNode;
}

export default function Body({ children }: Props) {
    const classes = useStyles();

    return (
        <Container maxWidth="lg" className={classes.root}>
            {children}
        </Container>
    );
}
