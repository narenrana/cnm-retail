function storeAuth(auth) {
  localStorage.setItem("auth", JSON.stringify(auth));
}

function getToken() {
  const auth = localStorage.getItem("auth") || "{}";
  const authJosn = JSON.parse(auth);
  return authJosn.token;
}

function refreshToken() {
  const token = getToken();
  console.log({ token });
}
function getAuthState() {
  const auth = localStorage.getItem("auth") || "{}";
  const authJosn = JSON.parse(auth);
  return authJosn;
}
function clearToken() {
  localStorage.removeItem("auth");
}

refreshToken();

export { storeAuth, getToken, refreshToken, getAuthState, clearToken };
