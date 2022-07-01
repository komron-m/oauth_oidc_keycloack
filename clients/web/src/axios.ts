import axios from "axios";
import {GetIDToken} from "./auth";

axios.interceptors.request.use((conf) => {
    conf.headers!["Authorization"] = `Bearer ${GetIDToken()}`
    return conf
})

export default axios;
