import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Input from "@material-ui/core/Input";
import InputLabel from "@material-ui/core/InputLabel";
import InputAdornment from "@material-ui/core/InputAdornment";
import FormControl from "@material-ui/core/FormControl";
import { useDispatch, useSelector } from "react-redux";
import Grid from "@material-ui/core/Grid";
import { useHistory } from "react-router-dom";
import IconButton from "@material-ui/core/IconButton";

import clsx from "clsx";
import Visibility from "@material-ui/icons/Visibility";
import VisibilityOff from "@material-ui/icons/VisibilityOff";
import Button from "@material-ui/core/Button";
import { withRouter } from "react-router-dom";
import { login } from "../../common/redux";

const useStyles = makeStyles((theme) => ({
  root: {
    marginTop: "10%",
    display: "flex",
    width: "100%",
    justifyContent: "center",
  },
  margin: {
    margin: theme.spacing(1),
  },
  container: {
    display: "flex",
    maxWidth: 500,
    margin: "100px",
    background: "rgb(76 175 80 / 5%)",
    padding: "100px",
  },
}));

export default withRouter(function Login(props) {
  const classes = useStyles();
  let history = useHistory();
  const dispatch = useDispatch();
  const { auth } = useSelector((state) => state.commonStore);
  const [values, setValues] = React.useState({
    email: "",
    password: "",
    rememberMe: true,
    showPassword: false,
  });

  const handleChange = (prop) => (event) => {
    setValues({ ...values, [prop]: event.target.value });
  };

  const handleClickShowPassword = () => {
    setValues({ ...values, showPassword: !values.showPassword });
  };

  const handleMouseDownPassword = (event) => {
    event.preventDefault();
  };

  const onSignIn = () => {
    const { email, password, rememberMe = true } = values;
    dispatch(login({ email, password, rememberMe }));
    //dispatch(loadCoupon("ywuqywuq"));
  };

  const onSignUpClick = () => {
    history.push("/signup");
  };

  useEffect(() => {
    if (auth.isLogin) {
      history.push("/home");
    } else {
      //history.push("/login");
    }
  }, [history, auth.isLogin]);

  return (
    <div className={classes.root}>
      <div className={classes.container}>
        <Grid container>
          <Grid item xs={12} md={12} xl={6}>
            <FormControl
              className={clsx(classes.margin, classes.textField)}
              fullWidth
            >
              <InputLabel htmlFor="standard-adornment-password">
                User Email
              </InputLabel>
              <Input
                id="standard-adornment-password"
                type={"text"}
                value={values.email}
                onChange={handleChange("email")}
              />
            </FormControl>
            <FormControl
              className={clsx(classes.margin, classes.textField)}
              fullWidth
            >
              <InputLabel htmlFor="standard-adornment-password">
                Password
              </InputLabel>
              <Input
                id="standard-adornment-password"
                type={values.showPassword ? "text" : "password"}
                value={values.password}
                onChange={handleChange("password")}
                endAdornment={
                  <InputAdornment position="end">
                    <IconButton
                      aria-label="toggle password visibility"
                      onClick={handleClickShowPassword}
                      onMouseDown={handleMouseDownPassword}
                    >
                      {values.showPassword ? <Visibility /> : <VisibilityOff />}
                    </IconButton>
                  </InputAdornment>
                }
              />
            </FormControl>

            <FormControl
              className={clsx(classes.margin, classes.textField)}
              fullWidth
            >
              <Button
                variant="outlined"
                color="primary"
                className={classes.checkout}
                onClick={onSignIn}
              >
                Sign In
              </Button>
            </FormControl>
            <FormControl
              className={clsx(classes.margin, classes.textField)}
              fullWidth
            >
              <Button
                variant="outlined"
                color="primary"
                className={classes.checkout}
                onClick={onSignUpClick}
              >
                Sign Up
              </Button>
            </FormControl>
          </Grid>
        </Grid>
      </div>
    </div>
  );
});
