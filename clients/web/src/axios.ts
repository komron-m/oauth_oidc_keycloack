import axios from "axios";

axios.interceptors.request.use((conf) => {
    conf.headers!["Authorization"] = "none"
    return conf
})

export default axios;
