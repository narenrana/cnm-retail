import { fade,makeStyles } from '@material-ui/core/styles';
export default makeStyles((theme) => ({
    toolbar: {
      borderBottom: `1px solid ${theme.palette.divider}`,
      display: 'flex',
      justifyContent: 'flex-end'
    },
    toolbarTitle: {
      flex: 1,
    },
    toolbarSecondary: {
      justifyContent: 'flex-start',
      overflowX: 'auto',
      background: '#90909014'
    },
    toolbarLink: {
      padding: theme.spacing(1),
      flexShrink: 0,
    },
    rightMenuPanel: {
        flexDirection: 'start-end'
    },
    extendedIconText: {
      fontSize: '0.8em',
      textTransform: 'capitalize',
      margin: '0 10px'
    },
    storeButton:{
      margin: '0 20px',
      minWidth: '34px',
      borderRadius: '8px',
      boxShadow: 'none',
      background: '#e0e0e0'
    },
    search: {
      position: 'relative',
      borderRadius: theme.shape.borderRadius,
      backgroundColor: fade(theme.palette.common.white, 0.15),
      '&:hover': {
        backgroundColor: fade(theme.palette.common.white, 0.25),
      },
      marginRight: theme.spacing(2),
      marginLeft: 0,
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(3),
        width: 'auto',
      },
    },
    searchIcon: {
      padding: theme.spacing(0, 2),
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
    inputRoot: {
      color: 'inherit',
    },
    inputInput: {
      padding: theme.spacing(1, 1, 1, 0),
      // vertical padding + font size from searchIcon
      paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
      transition: theme.transitions.create('width'),
      width: '100%',
      [theme.breakpoints.up('md')]: {
        width: '20ch',
      }
    },
  }));