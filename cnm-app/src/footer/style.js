import { makeStyles } from '@material-ui/core/styles';
export default makeStyles((theme) => ({
    footer: {
        backgroundColor: theme.palette.background.paper,
        // marginTop: theme.spacing(8),
        padding: theme.spacing(6, 0),
    },
    footerBanner: {
        objectFit: 'cover',
        width: '100%'
    },
    socialLinks: {
        padding: '8px'
    },
    subscribeRoot: {
       display: 'flex',
       position: 'relative'
    },
    subscribeMenu: {
        position: 'absolute',
        width: '400px',
        left: 'calc(50% - 200px)',
        top: '20%'
    },
    subscribeInputContainer:{
        display: 'flex',
        borderRadius: '26px',
        width: '100%',
        justifyContent: 'space-around',
        color: '#fff'
    },
    subscribeTypehead: {
        color: '#fff',
        fonWeight: '500',
        fontSize: '1.1rem',
        padding: '5px'
    },
    subscribeInput: {
        backgroundColor: '#fff',
        width: '100%',
        borderRadius: '20px',
        padding: '2px 15px',
        fontSize: '0.8rem',
    },
    subscribeButton: {
        margin: '0 10px'
    },
    subscribeButtonText: {
        color: '#fff',
        textTransform: 'capitalize',
    },
    
}));