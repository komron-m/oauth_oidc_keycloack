import axios from "axios";
import {GetAccessToken} from "./auth";

axios.interceptors.request.use((conf) => {
    conf.headers!["Authorization"] = `Bearer ${GetAccessToken()}`
    return conf
})

export default axios;
