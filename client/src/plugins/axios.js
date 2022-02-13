import axios from "axios";

export const http = axios.create({
});

export default {
  install: (app, options) => {
    app.config.globalProperties.$http = http;
  }
}