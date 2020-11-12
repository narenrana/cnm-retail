import * as _ from "lodash";
import { getToken, clearToken } from "../../core";

async function httpClient(URL, options) {
  const token = getToken();
  if (token) {
    const commonOptions = {
      headers: { Authorization: token },
      mode: "cors",
      withCredentials: true,
    };
    options.headers = { ...commonOptions.headers, ...options.headers };
  }
  return await fetch("http://localhost:8580/api" + URL, options)
    .then(async (res) => {
      console.log(res.status);
      if (res.status === 401) {
        clearToken();
        window.location.href = "/login";
      }
      return await res.json();
    })
    .then(async (res) => {
      return await res;
    })
    .catch((e) => {
      return e;
    });
}

export default httpClient;
