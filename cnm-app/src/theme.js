
import { createMuiTheme, responsiveFontSizes } from '@material-ui/core/styles';

export default responsiveFontSizes(createMuiTheme({
    overrides: {
        MuiNativeSelect: {
            select:{
                color: '#707070',
                fontSize:'0.8rem'
            }
        },
        MuiTypography: {
            root: {
                color: '#3c3b3b',
                fontSize: '1rem'
            },
            subtitle1: {
                fontSize: '0.8rem',
                textTransform: "uppercase",
                fontWeight: '600'
            },
            h6: {
                fontSize: '1rem !important',
                color: '#3c3b3b',
            }
        },
        MuiMenuItem: {
            root: {
                fontSize: '0.7rem',
                paddingTop: '5px',
                paddingBottom: 0,
                paddingLeft: 0
            }
        },
        MuiListItemText: {
            primary: {
                color: '#3c3b3b',
                textTransform: 'uppercase',
                fontSize: '0.8em',
                fontWeight: '600'
            }
        },
        MuiIconButton: {
            root: {
                color: '#707070'
            },
            label: {
                color: '#fff'
            }
        },
        MuiSvgIcon: {
            root: {
                color: '#707070'
            }
        },
        MuiButton: {
            // Name of the rule
            text: {

                color: '#fff',
            },
        },
        MuiFab: {
            primary: {
                color: '#707070'
            },
            secondary: {

            },
            sizeSmall: {
                height: '30px !important'
            }


        },
        MuiBadge: {
            colorPrimary: {
                color: '#fff',
            }
        }
    },
    palette: {
        color: '#dddd',
        primary: {
            // light: will be calculated from palette.primary.main,
            main: '#4caf50',
            light: '#81c784',
            color: '#707070'
            // dark: will be calculated from palette.primary.main,
            // contrastText: will be calculated to contrast with palette.primary.main
        },
        secondary: {
            light: '#eeeeee',
            main: '#dddddd',
            // dark: will be calculated from palette.secondary.main,
            contrastText: '#4caf50',
        },
        // Used by `getContrastText()` to maximize the contrast between
        // the background and the text.
        contrastThreshold: 3,
        // Used by the functions below to shift a color's luminance by approximately
        // two indexes within its tonal palette.
        // E.g., shift from Red 500 to Red 300 or Red 700.
        tonalOffset: 0.2,

    }
}));
